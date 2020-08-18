package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/samutayuga/grpcudemy/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	log.Println("This is client ")
	//connect to server

	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while dialing %v", err)
	}
	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)
	//doUnary(c)
	doBiDiStreaming(c)

}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a unary RPC")
	//	fmt.Printf("Created the client %v\n", c)
	res, err := c.Greet(context.Background(), &greetpb.GreetRequest{Greeting: &greetpb.Greeting{FirstName: "Sam", LastName: "Mutayuga"}})
	if err != nil {
		log.Fatalf("error while calling Greet rpc %v", err)
	}
	log.Printf("response from greet: %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC")

	resStream, err := c.GreetManyTimes(context.Background(), &greetpb.GreetManyTimesRequest{Greeting: &greetpb.Greeting{FirstName: "Jack", LastName: "Conor"}})

	if err != nil {
		log.Fatalf("error while calling GreetManyTimes rpc %v", err)
	}
	//log.Printf("response from GreetManyTimes: %v", res.r)

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		} else {
			log.Printf("received %s", msg.GetResult())
		}
	}

}

func doClientStreaming(c greetpb.GreetServiceClient) {
	st, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet %v", err)
	}

	//prepare the payload

	p := make([]*greetpb.LongGreetRequest, 0)
	//sam
	s := &greetpb.LongGreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Sam"}}
	//jack
	j := &greetpb.LongGreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Jacks"}}
	//Ella
	e := &greetpb.LongGreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Ella"}}
	p = append(p, s, j, e)

	for _, req := range p {
		if err := st.Send(req); err != nil {
			log.Fatalf("Error while sending the request %v", err)
		}
		time.Sleep(100 * time.Millisecond)
	}
	r, err := st.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while closing the sending %v", err)
	}

	log.Printf("receive greeting %s", r.GetResult())

}

func doBiDiStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Bi Directional Streaming RPC")
	p := make([]*greetpb.GreetEveryoneRequest, 0)
	//sam
	s := &greetpb.GreetEveryoneRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Sam"}}
	//jack
	j := &greetpb.GreetEveryoneRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Jacks"}}
	//Ella
	e := &greetpb.GreetEveryoneRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Ella"}}
	p = append(p, s, j, e)
	if stream, err := c.GreetEveryone(context.Background()); err == nil {
		//send a bunch of greet
		waitc := make(chan struct{})
		go func() {
			for _, r := range p {
				log.Printf("Client dispatches %s\n", r.Greeting.GetFirstName())
				sendErr := stream.Send(r)
				if sendErr != nil {
					log.Fatalf("Error while sending the message %v", sendErr)
				}
				time.Sleep(1000 * time.Microsecond)

			}
			stream.CloseSend()
		}()
		//receive a bunch of messages
		go func() {

			for {
				rsp, errRcv := stream.Recv()
				if errRcv == io.EOF {
					break
				}
				if errRcv != nil {
					log.Fatalf("Error while reading client stream %v", errRcv)
					break
				}
				log.Printf("Client receives greet %s", rsp.GetResult())

			}
			close(waitc)
		}()
		//block until everything is done
		<-waitc
	} else {
		return
	}

}
