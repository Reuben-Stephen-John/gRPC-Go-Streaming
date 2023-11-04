package main

import (
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	//"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Reuben-Stephen-John/grpc-streaming/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	tls := true
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Error while loading CA trust certificate: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	//conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)
	//doGreet(c)
	//doGreetManyTimes(c)
	//doLongGreet(c)
	doGreetWithDeadline(c,5*time.Second)
	//doGreetWithDeadline(c, 1*time.Second)

}
