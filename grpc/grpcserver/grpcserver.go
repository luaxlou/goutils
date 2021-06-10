package grpcserver

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func Start(port string, onListen func(s *grpc.Server)) {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Start rpc server at", lis.Addr())
	s := grpc.NewServer()

	onListen(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
