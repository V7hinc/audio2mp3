# 使用官方的Go镜像作为基础镜像
FROM golang:1.20

# 设置工作目录
WORKDIR /app
RUN set -x;\
    apt update;\
    apt install ffmpeg -y;\
    ffmpeg -version;

# 复制项目文件到容器中
COPY . .

# 编译Go应用
RUN go build -o audio2mp3 ./audio2mp3.go

# 运行应用
CMD ["./audio2mp3"]
