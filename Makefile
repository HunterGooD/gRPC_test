compile: 
	protoc -I api/proto --go_out=plugins=grpc:pkg/api api/proto/adder.proto 
gocompile: 
	go build -v ./cmd/server/