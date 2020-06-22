package main

import (
	"fmt"
	grpc_server "go-micro-grpc-demo/grpc_server"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	grpc_server.Start("192.168.34.114:9090")

	for s := range c {
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			fmt.Println("退出", s)
			//服务退出

			time.Sleep(time.Second * 3) //等待三秒
			os.Exit(0)
		default:
			fmt.Println("other", s)
		}
	}
}
