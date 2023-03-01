## 学习使用go-zero

安装goctl

```go
go install github.com/zeromicro/go-zero/tools/goctl@latest
```

生成服务

```go
goctl api new helloworld 
```

微服务版的hello-world

新建一个user服务，新建一个order服务

```go
goctl api new user
goctl api new order

// 将两个目录加到工作区
go work init
go work use order
go work use user
```
创建一个rpc目录，写一个hello.proto文件
到rpc目录下生成文件
```go
goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
```

在order目录下生成
```go
goctl api go -api order.api -dir ./gen
```