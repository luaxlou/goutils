package grpcserver

import (
	"fmt"
	"log"
	"net"
	"os"
	"runtime"

	"github.com/kazegusuri/grpc-panic-handler"
	"google.golang.org/grpc"
)

func Start(port string, onListen func(s *grpc.Server), onError func(msg string)) {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Start rpc server at", lis.Addr())


	panichandler.InstallPanicHandler(func(p interface{}) {

		msg := ""
		switch p.(type) {
		case runtime.Error:
			msg = fmt.Sprintf("%v", p)
		default:
			msg = fmt.Sprintf("%v", p)
		}

		onError(msg)

		os.Exit(0)

		return
	})

	s := grpc.NewServer(
		grpc.UnaryInterceptor(panichandler.UnaryPanicHandler),

	)

	onListen(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
