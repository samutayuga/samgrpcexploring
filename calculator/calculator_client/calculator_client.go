package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"

	"github.com/samutayuga/samgrpcexploring/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	cc, err := grpc.Dial("127.0.0.1:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error while creating connection %v", err)
	}
	defer cc.Close()
	c := calculatorpb.NewCalculatorServiceClient(cc)
	//doBiDiStreaming(c)
	//req := calculatorpb.OperationRequest{Operand: &calculatorpb.Operand{FirstOperand: 10, SecondOperand: 45}}
	//res, err := c.Calculate(context.Background(), &req)
	//log.Printf("got result %v", res.GetResult())
	inputs := make([]int, 0)
	inputs = append(inputs, 5, 10, -10, 240, 225, -900, 900, 256, -256)
	findSquareRoot(c, inputs)
	inputs = make([]int, 0)
	inputs = append(inputs, 5, 10, 240, 225, 256)
	findSquareRoot(c, inputs)

	doServerStreaming(c)
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	s, err := c.Decompose(context.Background(), &calculatorpb.DecomposeRequest{Input: &calculatorpb.NumnerInput{NumberInput: 3435353555678}})

	if err != nil {
		log.Fatalf("error while calling Decompose %v", err)
	}
	pn := ""
	for {
		r, err := s.Recv()
		if err == io.EOF {
			break
		}
		pn = fmt.Sprintf("%s %s", pn, strconv.Itoa(int(r.GetPrimeNumber())))
	}
	log.Printf("received %s", pn)
}
func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Average service %v", err)
	}
	//declate numbers
	fs := make([]float64, 0)
	fs = append(fs, 70.5, 67.89, 334.67, 776.65)
	//send many times
	for _, f := range fs {
		if err := stream.Send(&calculatorpb.AverageRequest{
			AvgInput: &calculatorpb.AvgInput{
				AvgInput: f}}); err != nil {

			log.Fatalf("Error while sending the request %v", err)

		}
		time.Sleep(100 * time.Millisecond)
	}

	r, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while closing the sending %v", err)
	}

	log.Printf("receive result %f", r.GetAverage())
}

func findSquareRoot(c calculatorpb.CalculatorServiceClient, inputs []int) {
	//assuming there are few numbers
	if stream, err := c.RootSquare(context.Background()); err == nil {

		waitc := make(chan struct{})
		//submit
		go func() {
			for _, anum := range inputs {
				//build payload for the request
				pReq := &calculatorpb.RootSquareRequest{Number: int32(anum)}
				if errSend := stream.Send(pReq); errSend != nil {
					log.Fatalf("Error while requesting %v", errSend)
				}
			}
			stream.CloseSend()

		}()
		//handle response
		go func() {
			for {
				resp, errResp := stream.Recv()
				if errResp == io.EOF {
					break
				}
				if errResp != nil {
					if e, ok := status.FromError(errResp); ok {
						if e.Code() == codes.InvalidArgument {
							log.Printf("negative number sent %s\n", e.Message())
							break

						} else {
							log.Fatalf("Server error code %v, description=%s\n", e.Code(), e.Message())

						}

					} else {
						log.Fatalf("Server error code %v\n", errResp)
					}
				}
				//print response
				log.Printf("Root for %d is %f", resp.GetInput(), resp.GetRootSquare())

			}
			close(waitc)
		}()
		//close when completed
		<-waitc
	}

}

func doBiDiStreaming(c calculatorpb.CalculatorServiceClient) {
	if stream, err := c.Maximum(context.Background()); err == nil {
		ins := make([]int, 0)
		ins = append(ins, 1, 5, 3, 6, 2, 20)
		//for {
		waitc := make(chan struct{})
		//bunch or numbers
		go func() {
			for _, in := range ins {
				reqPayLoad := &calculatorpb.MaximumRequest{
					MaxInput: &calculatorpb.Maximum{
						MaxInput: int32(in)}}
				sendErr := stream.Send(reqPayLoad)
				if sendErr != nil {
					log.Fatalf("Error while sending the number %v", sendErr)
				}
				time.Sleep(1000 * time.Millisecond)
			}

			stream.CloseSend()

		}()

		//bunch of response
		go func() {
			//the last response will give the signal to close the channel
			for {
				rsp, errRsp := stream.Recv()
				if errRsp == io.EOF {
					break
				}
				if errRsp != nil {
					log.Fatalf("Error while reading response  %v", errRsp)
				}

				max := rsp.GetMaxOut()
				log.Printf("client receives %d\n", max)
			}
			close(waitc)
		}()

		//stop until all are done
		<-waitc
		//	}

	}

}
