# Go_gRPC_Peer
gRPC peer example by Golang


### Commands

```
apt install protobuf-compiler
```

```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative mao.proto

或

protoc   --go_out=Mgrpc/service_config/service_config.proto=/internal/proto/grpc_service_config:.   --go-grpc_out=Mgrpc/service_config/service_config.proto=/internal/proto/grpc_service_config:.   --go_opt=paths=source_relative   --go-grpc_opt=paths=source_relative mao.proto
```

### Prepare for protoc, protoc-gen-go, protoc-gen-go-grpc

```
apt install protobuf-compiler
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
添加路径到环境变量PATH：export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin
```

### 设置 GO111MODULE
可以用环境变量 GO111MODULE 开启或关闭模块支持，它有三个可选值：off、on、auto，默认值是 auto。

* GO111MODULE=off 无模块支持，go 会从 GOPATH 和 vendor 文件夹寻找包。
* GO111MODULE=on 模块支持，go 会忽略 GOPATH 和 vendor 文件夹，只根据 go.mod 下载依赖。
* GO111MODULE=auto 在 $GOPATH/src 外面且根目录有 go.mod 文件时，开启模块支持。

在使用模块的时候，GOPATH 是无意义的，不过它还是会把下载的依赖储存在 GOPATH/pkg/mod中，也会把goinstall的结果放在GOPATH/bin 中。
