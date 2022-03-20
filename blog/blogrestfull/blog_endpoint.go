package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/samutayuga/samgrpcexploring/blog/cfg"
	"github.com/samutayuga/samgrpcexploring/blog/restutil"

	gw "github.com/samutayuga/samgrpcexploring/blog/blogpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

var (
	blogConfig cfg.Config
	client     gw.BlogServiceClient
)

func init() {
	blogConfig = cfg.LoadConfig()
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	serverString := fmt.Sprintf("localhost:%d", blogConfig.ServerPort)
	grpcServerEndPoint := flag.String("blog-server-endpoint", serverString, "gRPC Server Endpoint")
	log.Printf("Using gRPC server at %s\n", *grpcServerEndPoint)
	//dial grpc server
	if conn, errDial := grpc.Dial(serverString, opts); errDial != nil {
		panic(errDial)
	} else {
		client = gw.NewBlogServiceClient(conn)
		restutil.BlogClient = client
	}
}
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
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

	routers := mux.NewRouter()
	routers.HandleFunc("/v1/samblog/{blogId}", restutil.ProcessASingleBlog).Methods("GET", "DELETE")
	routers.HandleFunc("/v1/samblog", restutil.UpdateBlog).Methods("PATCH")
	routers.HandleFunc("/v1/samblog", restutil.CreateBlog).Methods("POST")
	routers.HandleFunc("/v1/samblog", restutil.ListBlog).Methods("GET")
	routers.HandleFunc("/liveness", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) }).Methods("GET")
	routers.HandleFunc("/readiness", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) }).Methods("GET")

	httpServerString := fmt.Sprintf(":%d", blogConfig.EndpointPort)
	log.Printf("Running REST end point at %d\n", blogConfig.EndpointPort)
	return http.ListenAndServe(httpServerString, routers)
}
