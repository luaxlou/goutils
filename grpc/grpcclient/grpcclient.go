package grpcclient

import (
	"log"
	"time"

	"google.golang.org/grpc"
)

func Connect(address string, onConnected func(conn *grpc.ClientConn)) {

	log.Println("grpc connecting to", address)

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)

		time.Sleep(time.Millisecond * 100)
		Connect(address, onConnected)
	}

	onConnected(conn)

}
