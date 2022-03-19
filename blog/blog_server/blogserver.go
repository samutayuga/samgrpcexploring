package main

import (
	"fmt"
	"github.com/samutayuga/samgrpcexploring/pg"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/samutayuga/samgrpcexploring/blog/blogcommon"
	"github.com/samutayuga/samgrpcexploring/blog/blogpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	reflection.Register(s)
	blogpb.RegisterBlogServiceServer(s, &blogcommon.PgBlogServer{})
	log.Println("service is registered")
	go func() {
		fmt.Println("Server starting....")
		//connect to db
		pg.Init()
		pg.PingDb()
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to server %v", err)
		}

	}()
	//wait the control c key
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	//disconnect from db
	pg.CloseDb()
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Closing the listener")
	lis.Close()
}
