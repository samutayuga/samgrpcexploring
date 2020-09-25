package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	cqlcrude "github.com/samutayuga/samgrpcexploring/cqlcrude/cqlcrudepb"
	"github.com/samutayuga/samgrpcexploring/cqlcrude/cqlimpl"
	"github.com/samutayuga/samgrpcexploring/sandra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	reflection.Register(s)
	//blogpb.RegisterBlogServiceServer(s, &blogcommon.Server{})
	cqlcrude.RegisterCrudeServiceServer(s, &cqlimpl.CQLService{})
	log.Println("service is registered")
	go func() {
		fmt.Println("Server starting....")
		//connect to db
		sandra.Csessinit(cqlimpl.DevKS)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to server %v", err)
		}

	}()
	//wait the control c key
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	//disconnect from db
	sandra.Csessclose()
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Closing the listener")
	lis.Close()
}
