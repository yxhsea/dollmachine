### 娃娃机用户端Api

```
制作镜像: docker build -t doll-user:v1 .

//以特权方式启动容器 指定--privileged参数
创建容器:  
docker run -d \
--privileged \
--name dollUser \
-p 9554:9554 \
-v /var/gopath/src/dollmachine/dolluser:/go/src doll-user:v1
```