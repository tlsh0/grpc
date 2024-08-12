package main

import (
	"fmt"
	"net"
	"os"

	"github.com/tlsh0/protoapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var port = ":8080"

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Using default port:", port)
	} else {
		port = os.Args[1]
	}

	server := grpc.NewServer()

	var randomServer RandomServer
	protoapi.RegisterRandomServer(server, randomServer)

	reflection.Register(server)

	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Serving requests...")
	server.Serve(listen)
}
