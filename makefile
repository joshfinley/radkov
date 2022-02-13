.PHONY: proto
proto:
	protoc -I pkg/rkpb/ pkg/rkpb/rkpb.proto --go_out=paths=source_relative:pkg/rkpb/ --go-grpc_out=pkg/rkpb/ --go-grpc_opt=paths=source_relative 
