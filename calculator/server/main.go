//go:build !test
// +build !test

package main

import (
	"log"
	"net"

	// "time"

	pb "github.com/Reuben-Stephen-John/grpc-streaming/calculator/proto"

	"google.golang.org/grpc"
)

type Server struct {
	pb.CalculatorServiceServer
}

// var greetWithDeadlineTime time.Duration = 1 * time.Second

var addr string = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	// defer lis.Close()
	log.Printf("Listening at %s\n", addr)

	// opts := []grpc.ServerOption{}

	// tls := false // change that to true if needed
	// if tls {
	//     certFile := "ssl/server.crt"
	//     keyFile := "ssl/server.pem"
	//     creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

	//     if err != nil {
	//         log.Fatalf("Failed loading certificates: %v\n", err)
	//     }
	//     opts = append(opts, grpc.Creds(creds))
	// }

	// opts = append(opts, grpc.ChainUnaryInterceptor(LogInterceptor(), CheckHeaderInterceptor()))

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})
	
	// defer s.Stop()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
