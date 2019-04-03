## gRPC_Golang 学习

模仿`grpc/examples/helloworld`写一个Login的gRPC服务端和客户端

### protoc命令

- `--proto_path`: 指定了在哪个目录中搜索 import 中导入的和要编译为.go 的 proto 文件，可以定义多个
- `--go_out`: 指定了生成的 go 文件的目录

```shell
protoc --go_out=$GOPATH/src ./model/model.proto
protoc --proto_path=. --proto_path=./model/ --go_out=plugins=grpc:$GOPATH/src ./service/login.proto
```

### 运行

```sh
go run login_server/main.go
go run login_client/main.go
```

### 遗留问题

* 发现protoc命令不支持`go mod`项目, 我想在当前文件夹下生成文件,且使用指定的package名称, 结果就会在当前文件下生成package多级目录, 需要手工copy下, 这个暂时我还没找到好方法.

 