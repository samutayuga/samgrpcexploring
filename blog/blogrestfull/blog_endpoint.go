package main

import (
	"context"
	"flag"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	gw "github.com/samutayuga/samgrpcexploring/blog/blogpb"
	"google.golang.org/grpc"
)

var (
	//command line option
	//gRPC server endpoint
	grpcServerEndPoint = flag.String("blog-server-endpoint", "localhost:50051", "gRPC Server Endpoint")
)

func main() {
	flag.Parse()
	defer glog.Flush()
	if err := run(); err != nil {
		glog.Fatalf("error %v", err)
	}
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	//Register the gateway end point to grpc Server
	//make sure the gRPC server is running
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterBlogServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndPoint, opts)
	if err != nil {
		return err
	}
	//start http server and proxy the call to gRPC endpoint
	return http.ListenAndServe(":8001", mux)
}
