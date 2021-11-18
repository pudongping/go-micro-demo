module go-micro-demo

go 1.16

require (
	github.com/golang/protobuf v1.5.2
	github.com/micro/go-micro v1.18.0
)

require (
	github.com/json-iterator/go v1.1.8 // indirect
	github.com/miekg/dns v1.1.22 // indirect
	github.com/nats-io/nats.go v1.9.1 // indirect
	go.uber.org/zap v1.12.0 // indirect
	google.golang.org/genproto v0.0.0-20191108220845-16a3f7862a1a // indirect
	google.golang.org/grpc v1.25.1 // indirect
)

replace (
	github.com/golang/protobuf => github.com/golang/protobuf v1.3.2
	github.com/micro/go-micro => github.com/Lofanmi/go-micro v1.16.1-0.20210804063523-68bbf601cfa4
	google.golang.org/genproto => google.golang.org/genproto v0.0.0-20191216164720-4f79533eabd1
	google.golang.org/grpc v1.30.0 => google.golang.org/grpc v1.26.0
)
