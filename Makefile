.PHONY:build-proto
build-proto:
	protoc -I grpcserver --go_out=plugins=grpc:grpcserver/ grpcserver/grpcserver.proto
build-docker:
	docker build . -t grpc-http