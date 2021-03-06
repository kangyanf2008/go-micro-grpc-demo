//go module proxy 配置  https://goproxy.cn
module go-micro-grpc-demo

go 1.13.4

replace github.com/imdario/mergo => github.com/imdario/mergo v0.3.8

replace golang.org/x/text => github.com/golang/text v0.3.3

replace golang.org/x/tools => github.com/golang/tools v0.0.0-20200619210111-0f592d2728bb

replace golang.org/x/net => github.com/golang/net v0.0.0-20200602114024-627f9648deb9

replace golang.org/x/crypto => github.com/golang/crypto v0.0.0-20200604202706-70a84ac30bf9

replace golang.org/x/exp => github.com/golang/exp v0.0.0-20200513190911-00229845015e

replace cloud.google.com/go => github.com/googleapis/google-cloud-go v0.58.0

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace google.golang.org/genproto => github.com/googleapis/go-genproto v0.0.0-20200620020550-bd6e04640131

replace google.golang.org/api => github.com/googleapis/google-api-go-client v0.28.0

replace github.com/coreos/bbolt v1.3.4 => go.etcd.io/bbolt v1.3.4

replace go.etcd.io/bbolt v1.3.4 => github.com/coreos/bbolt v1.3.4

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0

//go get -u -v github.com/golang/protobuf/protoc-gen-go@v1.2.0
require (
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/google/btree v1.0.0 // indirect
	github.com/micro/go-micro/v2 v2.9.0
	github.com/micro/go-plugins/registry/zookeeper/v2 v2.8.0
	github.com/miekg/dns v1.1.29 // indirect
	github.com/nats-io/jwt v1.0.1 // indirect
	github.com/nats-io/nats.go v1.10.0 // indirect
	github.com/nats-io/nkeys v0.2.0 // indirect
	go.uber.org/zap v1.15.0 // indirect
	golang.org/x/crypto v0.0.0-20200604202706-70a84ac30bf9 // indirect
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9
	golang.org/x/sys v0.0.0-20200620081246-981b61492c35 // indirect
	golang.org/x/text v0.3.3 // indirect
	golang.org/x/tools v0.0.0-20200619210111-0f592d2728bb // indirect
	google.golang.org/grpc v1.29.1
	gopkg.in/yaml.v2 v2.3.0 // indirect
)
