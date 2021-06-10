package grpcserver

import (
	"fmt"
	"log"
	"net"
	"runtime"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
)

func Start(port string, onListen func(s *grpc.Server), onError func(msg string)) {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Start rpc server at", lis.Addr())

	recoverHandler := func(p interface{}) error {

		msg := ""
		switch p.(type) {
		case runtime.Error:
			msg = fmt.Sprintf("runtime error: %v", p)
		default:
			msg = fmt.Sprintf("error: %v", p)
		}

		onError(msg)

		return nil
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(recoverHandler)),
		)))

	onListen(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
