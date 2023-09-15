build_proto:
	rm -rf ./proto/thegamblr; protoc --proto_path=proto proto/*.proto --go_out=proto --go-grpc_out=proto