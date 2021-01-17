module github.com/vandong9/learn_go_microservice_1/consignment-service

go 1.14

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

// replace github.com/lucas-clemente/quic-go => github.com/lucas-clemente/quic-go v0.14.1
// replace github.com/coreos/etcd => github.com/coreos/etcd v3.4.14+incompatible

require (
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/cmd/protoc-gen-micro/v2 v2.0.0-20210105173217-bf4ab679e18b // indirect
	github.com/micro/go-micro/v2 v2.8.1-0.20200603084508-7b379bf1f16e
	github.com/micro/micro/v2 v2.8.1-0.20200603100651-e57d42a20d26
	github.com/pkg/errors v0.9.1
	go.mongodb.org/mongo-driver v1.2.1
	golang.org/x/net v0.0.0-20200520182314-0ba52f642ac2
	google.golang.org/grpc v1.29.1
)
