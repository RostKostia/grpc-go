package main

import (
	"context"
	"github.com/RostKostia/grpc-go/echo"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	echo.UnimplementedEchoServiceServer
}

func (s *server) Shout(ctx context.Context, in *echo.EchoRequest) (*echo.EchoResponse, error) {
	return &echo.EchoResponse{Message: in.String() + " Accepted"}, nil
}

func main() {
	for {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen on port 50051: %v", err)
		}

		s := grpc.NewServer()
		echo.RegisterEchoServiceServer(s, &server{})
		log.Printf("gRPC server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}

}
