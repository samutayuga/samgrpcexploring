package main

import (
	"log"
	"net"

	"github.com/samutayuga/samgrpcexploring/greet/greetpb"
	"github.com/samutayuga/samgrpcexploring/greet/serverimpl"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	certFile := "ssl\\server.crt"
	keyFile := "ssl\\server.pem"
	creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
	if sslErr != nil {
		log.Fatalf("Failed loading %v", sslErr)
		return
	}
	opt := grpc.Creds(creds)
	s := grpc.NewServer(opt)
	svr := serverimpl.Server{}
	reflection.Register(s)
	greetpb.RegisterGreetServiceServer(s, &svr)
	log.Printf("service is registered %v", s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
