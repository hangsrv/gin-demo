# 镜像基准
FROM golang:1.17
# 作者
MAINTAINER hang
# 环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    GIN_MODE=release \
    PORT=8080
# 工作目录
WORKDIR /app
# 拷贝项目
COPY . .
# 编译项目
RUN go build .
# 暴露端口
EXPOSE 8080
# 启动命令
ENTRYPOINT ["./gin_demo"]
