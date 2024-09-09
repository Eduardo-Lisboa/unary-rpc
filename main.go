package main

import (
	"unary-rpc/client"
	"unary-rpc/server"
)

func main() {
	go server.Run()
	client.Run()
}
