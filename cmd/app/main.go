package main

import (
	"api-server-template/cmd/server"
	"fmt"
)

func main() {
	fmt.Println("Starting Server")
	server.Run()
}
