build_proto:
	rm -rf ./proto/thegamblr; mkdir proto/thegamblr; mkdir proto/thegamblr/java; mkdir proto/thegamblr/python; protoc --proto_path=proto proto/*.proto --go_out=proto --go-grpc_out=proto  --python_out=proto/thegamblr/python --java_out=proto/thegamblr/java


casino:
	go build /Users/nlerner/GolandProjects/TheGamblr/server/main/main.go; ./main

casino.kill:
	lsof -t -i:8080 | xargs kill -9
