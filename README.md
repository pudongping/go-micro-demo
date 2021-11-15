# 学习 go-micro 时的 demo

## 安装 Protobuf 相关工具

### 安装 protoc-gen-micro

该工具适用于在 Micro 框架中基于 Protobuf 文件生成服务代码

```shell
# 在项目根目录下运行
go get -u github.com/micro/protoc-gen-micro
```

### 安装 protoc

```shell

# 根据你自己的系统下载对应的源码包 （比如我这里使用的是 mac book，我就要下载 osx 压缩包）
cd ~/go-tools && wget https://github.com/protocolbuffers/protobuf/releases/download/v3.19.1/protoc-3.19.1-osx-x86_64.zip

# 解压缩
unzip protoc-3.19.1-osx-x86_64.zip

# 编辑配置文件
vim ~/.zshrc

# 写入以下配置信息
export PATH="/Users/pudongping/go-tools/protoc-3.19.1-osx-x86_64/bin:$PATH"

source ~/.zshrc

```

## 使用 Etcd 作为注册中心

### 安装 Etcd

```shell

# 根据你自己的系统下载对应的源码包（比如我这里使用的是 mac book，我就要下载 etcd-v3.5.1-darwin-amd64.zip 包）
cd ~/go-tools && wget https://github.com/etcd-io/etcd/releases/download/v3.5.1/etcd-v3.5.1-darwin-amd64.zip

unzip unzip etcd-v3.5.1-darwin-amd64.zip

```

### 测试

```shell

# 进入 etcd 解压缩后的文件夹中
cd etcd-v3.5.1-darwin-amd64

# 启动 etcd 服务器
./etcd

# 然后新开一个命令行窗口，通过 etcdctl 指令进行客户端测试
./etcdctl put alex "hello"  # output: OK

./etcdctl get alex  # output: alex hello
```

## 通过服务声明生成原型代码文件

在项目根目录下执行以下命令

```shell
protoc -I. --go_out=plugins=micro:. proto/meet.proto

# 亦或者可以在项目根目录下直接执行以下指令 （详情可见项目根目录下的 Makefile 文件）
make build
```

如果在执行上述 protoc 命令时出错，提示 `micro` 插件不存在时，可以执行以下命令

```shell
protoc --proto_path=. --micro_out=. --go_out=. proto/meet.proto
```

## 启动服务

在项目根目录下运行如下命令自动下载服务实现代码中的依赖

```shell
go mod tidy
```

启动服务

> 你也可以设置 `MICRO_REGISTRY=etcd` 环境变量，这样每次启动的时候就不需要带上 `--registry=etcd` 参数
> 了，直接使用 `go run main.go`

```shell
go run main.go --registry=etcd
```

## 关于报错

```shell
# 启动服务时：如果出现如下报错
panic: qtls.ConnectionState not compatible with tls.ConnectionState
```

那么解决方式是：要么你将 go 的版本降到 1.15 以下，要么使用 go-micro v1.16.1 的版本。
但是貌似如果是 M1 芯片的 Mac 的话，还不能将 go 降版本到 1.16 以下，具体的我没有操作，
如果你操作了，发现可以，希望你也告知我一下，谢谢。

这里感谢 [修复 go-micro 框架在高版本 Go 编译的运行时错误](https://imlht.com/archives/239/) 提供的解决方案。

还有一种方式就是直接更新包，操作方式如下：

直接将以下复制到 go.mod 文件中，然后执行 `go mod tidy` 

```shell

replace (
	github.com/micro/go-micro => github.com/Lofanmi/go-micro v1.16.1-0.20210804063523-68bbf601cfa4
	github.com/golang/protobuf => github.com/golang/protobuf v1.3.2
	google.golang.org/genproto => google.golang.org/genproto v0.0.0-20191216164720-4f79533eabd1
	google.golang.org/grpc v1.30.0 => google.golang.org/grpc v1.26.0
	google.golang.org/protobuf => google.golang.org/protobuf v1.25.0
)

```

重新启动服务，我们可以看到 etcd 控制台有类似于如下的输出内容

```shell

{"level":"info","ts":"2021-11-16T02:23:01.727+0800","caller":"traceutil/trace.go:171","msg":"trace[214976224] put","detail":"{key:/micro/registry/Meet/Meet-ee5b8119-4ed0-474b-b948-86ee166325a2; req_size:614; response_revision:3; }","duration":"102.469459ms","start":"2021-11-16T02:23:01.623+0800","end":"2021-11-16T02:23:01.725+0800","steps":["trace[214976224] 'process raft request'  (duration: 25.521542ms)","trace[214976224] 'marshal mvccpb.KeyValue'  (duration: 76.424667ms)"],"step_count":2}

```