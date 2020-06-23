package grpc_client

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/service"
	"github.com/micro/go-micro/v2/service/grpc"
	"github.com/micro/go-plugins/registry/zookeeper/v2"

	"go-micro-grpc-demo/protobuffer_def"
	"time"
)

func Start()  {

	r := zookeeper.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"127.0.0.1:2181"}
		op.Context = context.Background()
		op.Timeout = time.Second * 5
	})

	// create GRPC service
	service := grpc.NewService(
		service.Name("test.client"),
		service.Registry(r),
	)
	service.Client().Init(client.Retries(3),client.PoolSize(200), client.PoolTTL(time.Second*20), client.RequestTimeout(time.Second*5))

	test :=  protobuffer_def.NewTestService("test.server", service.Client())
	for i:=0; i < 20; i++ {
		go func() {
			i := 0
			for {
				_, err := test.BaseInterface(context.Background(), &protobuffer_def.BaseRequest{RequestId:"1111",C:protobuffer_def.CMD_TEST_1})
				if err != nil {
					fmt.Println(err)
				} else {
					i ++
				}
				if i % 10000 == 0 {
					fmt.Println(i, time.Now().Unix())
				}
			}
		}()
	}



}