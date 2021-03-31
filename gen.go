package main

import _ "github.com/go-redis/redis"

//go:generate protoc --proto_path=src/pb --go_out=src/pb --go_opt=paths=source_relative src/pb/proto.proto


func main() {

}