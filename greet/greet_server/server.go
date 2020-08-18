package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/samutayuga/grpcudemy/greet/greetpb"
	"google.golang.org/grpc"
)

//Server ...
type Server struct {
}

//Greet ...
func (s *Server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Printf("Greet function is invoked with %v\n", req)
	fName := req.GetGreeting().GetFirstName()
	result := fmt.Sprintf("hello %s", fName)
	res := greetpb.GreetResponse{
		Result: result,
	}
	return &res, nil
}

//GreetManyTimes ...
func (s *Server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function is invoked with %v\n", req)
	fName := req.Greeting.GetFirstName()
	for i := 0; i < 100; i++ {
		gw := fmt.Sprintf("Hello %s , greet number %d", fName, i)
		res := greetpb.GreetManyTimesResponse{Result: gw}
		stream.Send(&res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil

}

//LongGreet ...
func (*Server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	//format the response
	log.Println("LongGreet function is invoked with")
	rs := "Hello"
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			//send response to client
			return stream.SendAndClose(&greetpb.LongGreetResponse{Result: rs})
		}
		if err != nil {
			log.Fatalf("error while reading client stream %v", err)
		}
		//this is the place to exract the client stream

		fName := req.GetGreeting().GetFirstName()
		rs = fmt.Sprintf("%s %s!", rs, fName)
	}
	//return nil

}

//GreetEveryone ...
func (*Server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone function is invoked")

	for {
		req, errRcv := stream.Recv()
		//in case this is the last message
		if errRcv == io.EOF {
			return nil
		}
		//in case any error
		if errRcv != nil {
			log.Fatalf("Error while reading client stream %v", errRcv)
			return errRcv
		}
		//no error, no last message

		fName := req.GetGreeting().GetFirstName()
		log.Printf("Server greets %s\n", fName)
		result := fmt.Sprintf("Hello %s !", fName)

		//let send the response

		if sendErr := stream.Send(&greetpb.GreetEveryoneResponse{
			Result: result}); sendErr != nil {
			log.Fatalf("Error while sending the response to client %v", sendErr)
			return sendErr
		}

	}

}
func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &Server{})
	log.Printf("service is registered %v", s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
