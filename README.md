## 使用方法
### 构建docker镜像
```
docker build -t audio2mp3 .
```
### 运行容器
```
docker run --rm -it -v /path/to/audio/dir:/app/src_audio/ -v /path/to/savemp3/dir:/app/mp3/ audio2mp3
```