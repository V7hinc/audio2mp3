package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func convertToMP3(inputPath, outputPath string) error {
	cmd := exec.Command("ffmpeg", "-i", inputPath, "-codec:a", "libmp3lame", "-qscale:a", "3", outputPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
func DirExistOrCreateDir(path string) (bool, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		return false, err
	}
	return true, nil
}
func splitPath(filePath string) []string {
	// 使用filepath.ToSlash将路径中的斜杠标准化为正斜杠
	normalizedPath := filepath.ToSlash(filePath)

	// 使用strings.Split按照正斜杠分割路径
	dirSlice := strings.Split(normalizedPath, "/")

	// 过滤掉空的部分
	dirSlice = removeEmptyElements(dirSlice)

	return dirSlice
}

func removeEmptyElements(slice []string) []string {
	var filtered []string
	for _, s := range slice {
		if s != "" {
			filtered = append(filtered, s)
		}
	}
	return filtered
}
func main() {
	inDirPath := "src_audio" // 替换为目标目录的路径
	outDirPath := "mp3"      // 替换为目标目录的路径
	inDir := inDirPath
	outDir := outDirPath
	var err error
	_, err = DirExistOrCreateDir(outDir)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = filepath.WalkDir(inDir, func(fp string, fi os.DirEntry, err error) error {
		if err != nil {
			log.Println(err) // 此处处理访问文件时的错误
			return nil
		}
		if fi.IsDir() {
			return nil // 忽略目录
		}
		fmt.Println(fp)
		inputFilePath := fp
		filenameWithExtension := filepath.Base(inputFilePath)
		if strings.HasPrefix(filenameWithExtension, ".") {
			return nil
		}
		inputFilePathWithoutExtension := inputFilePath[:len(inputFilePath)-len(filepath.Ext(inputFilePath))]
		normalizedPath := splitPath(inputFilePathWithoutExtension + ".mp3")
		outputFilePathSlice := append([]string{outDirPath}, normalizedPath[1:]...)

		outputFilePath := filepath.Join(outputFilePathSlice...)
		outputFileDir := filepath.Dir(outputFilePath)
		_, err = DirExistOrCreateDir(outputFileDir)
		if err != nil {
			log.Fatal(err)
			return err
		}
		err = convertToMP3(inputFilePath, outputFilePath)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Conversion complete.")
		}
		return nil
	})
	if err != nil {
		fmt.Println(err) // 处理遍历目录时的错误
	}
}
