#源镜像
FROM golang:latest

#作者
MAINTAINER yxhSea "yxhsea@foxmail.com"

#设置工作目录
WORKDIR $GOPATH/src/dollmachine/dollunique

#将服务器的Go工程代码加入到dokcer容器中
ADD . $GOPATH/src/dollmachine/dollunique

#Go构建可执行文件
RUN go build .

#最终运行docker的命令
ENTRYPOINT ["./main"]