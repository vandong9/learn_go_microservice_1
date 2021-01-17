module github.com/vandong9/learn_go_microservice_1/consignment-cli

go 1.14

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/micro/go-micro/v2 v2.8.1-0.20200603084508-7b379bf1f16e
	github.com/micro/protobuf v0.0.0-20180321161605-ebd3be6d4fdb // indirect
	github.com/vandong9/learn_go_microservice_1/consignment-service v0.0.0-20210117062734-dd258c4cf4cb
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.24.0 // indirect
)
