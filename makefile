generate protobuf: 
	protoc --go_out=internal/adapters/framework/left/gRPC --proto_path=internal/adapters/framework/left/gRPC/proto internal/adapters/framework/left/gRPC/proto/number_msg.proto

	protoc --go-grpc_out=require_unimplemented_servers=false:internal/adapters/framework/left/gRPC --proto_path=internal/adapters/framework/left/gRPC/proto internal/adapters/framework/left/gRPC/proto/arithmetic_svc.proto