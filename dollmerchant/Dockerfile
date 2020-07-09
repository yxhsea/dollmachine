FROM golang:latest

MAINTAINER yxhSea "yxhsea@foxmail.com"

WORKDIR $GOPATH/src/dollmachine/dollmerchant

ADD . $GOPATH/src/dollmachine/dollmerchant

RUN go build .

EXPOSE 9555

ENTRYPOINT ["./DollMerchant", "-c", "dev_conf.toml"]