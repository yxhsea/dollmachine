FROM golang:latest

MAINTAINER yxhSea "yxhsea@foxmail.com"

WORKDIR $GOPATH/src/dollmachine/dollplatform

ADD . $GOPATH/src/dollmachine/dollplatform

RUN go build .

EXPOSE 9556

ENTRYPOINT ["./DollPlatform", "-c", "dev_conf.toml"]