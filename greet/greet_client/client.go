package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/samutayuga/samgrpcexploring/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

func main() {
	log.Println("This is client ")
	//connect to server
	certFile := "ssl\\ca.crt"
	creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
	if sslErr != nil {
		log.Fatalf("error file loading the ca Trust certificate %v", sslErr)
		return
	}
	opts := grpc.WithTransportCredentials(creds)
	conn, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("Error while dialing %v", err)
	}
	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)
	doUnary(c)
	doBiDiStreaming(c)
	doUnaryWithDeadLine(c, 5*time.Second)
	doUnaryWithDeadLine(c, 1*time.Second)

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
func doUnaryWithDeadLine(c greetpb.GreetServiceClient, second time.Duration) {
	fmt.Println("Starting to do a unary With deadline")
	req := &greetpb.GreetWithDeadLineRequest{Greeting: &greetpb.Greeting{
		FirstName: "Sam",
		LastName:  "Mutayuga",
	}}
	ctx, cancel := context.WithTimeout(context.Background(), second)
	defer cancel()
	//	fmt.Printf("Created the client %v\n", c)
	res, err := c.GreetWithDeadLine(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			if st.Code() == codes.DeadlineExceeded {
				fmt.Printf("Deadline exceeded after %v seconds", second)
			} else {
				fmt.Printf("Unexpected error %v", st)
			}
		} else {
			log.Fatalf("error while calling GreetWithDeadLineRequest rpc %v", err)
		}
		return

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
