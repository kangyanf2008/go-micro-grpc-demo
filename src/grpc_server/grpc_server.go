package grpc_server

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/service"
	"github.com/micro/go-micro/v2/service/grpc"
	"github.com/micro/go-plugins/registry/zookeeper/v2"
	"go-micro-grpc-demo/protobuffer_def"
	"time"
)

func Start(address string)  {

	r := zookeeper.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"127.0.0.1:2181"}
		op.Context = context.Background()
		op.Timeout = time.Second * 5
	})


	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create GRPC service
	service := grpc.NewService(
		service.Address(address),
		service.Name("test.server"),
		service.Registry(r),
		service.RegisterTTL(time.Second*30),
		service.RegisterInterval(time.Second*20),
		service.Context(ctx),
	)

	service.Init()

	// register test handler
	protobuffer_def.RegisterTestServiceHandler(service.Server(), &TestGrpcServiceImpl{})

	 if err :=  service.Run(); err != nil {
		 fmt.Println(err)
		 panic(err)
	 }
}
