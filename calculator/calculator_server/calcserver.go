package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"

	"github.com/samutayuga/grpcudemy/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type calcServer struct {
}

func (s *calcServer) Calculate(ctx context.Context, req *calculatorpb.OperationRequest) (*calculatorpb.OperationResponse, error) {
	log.Printf("Calculator is invoked with operation %v", req.Operand)
	return &calculatorpb.OperationResponse{
		Result: doSum(req.GetOperand()),
	}, nil
}
func (s *calcServer) Decompose(req *calculatorpb.DecomposeRequest, stream calculatorpb.CalculatorService_DecomposeServer) error {
	log.Printf("Decompose is invoked with operation %v", req)
	doDecompose(int(req.GetInput().GetNumberInput()), stream)
	return nil
}

func (*calcServer) Average(stream calculatorpb.CalculatorService_AverageServer) error {
	log.Println("Average is invoked with operation")
	return doAverage(stream)
}
func (*calcServer) Maximum(stream calculatorpb.CalculatorService_MaximumServer) error {
	log.Println("Maximum is invoked")
	currentMax := 0
	//allNumbers := make([]int, 0)
	for {
		req, errRcv := stream.Recv()
		if errRcv == io.EOF {
			return nil
		}
		if errRcv != nil {
			log.Fatalf("Error while receiving client request %v", errRcv)
			return errRcv
		}
		//deal with request
		in := req.GetMaxInput().GetMaxInput()
		//allNumbers = append(allNumbers, int(in))
		//apply the algo

		if int(in) > currentMax {
			currentMax = int(in)
			if sendErr := stream.Send(&calculatorpb.MaximumResponse{
				MaxOut: int32(currentMax)}); sendErr != nil {
				log.Fatalf("Error while sending the response to client %v", sendErr)
				return sendErr
			}
		}

	}
}

//Error handling
//This rpc will throw an exception if the sent number is negative
//The error being sent is of type INVALID_ARGUMENT
func (*calcServer) RootSquare(stream calculatorpb.CalculatorService_RootSquareServer) error {
	for {
		req, errRcv := stream.Recv()
		if errRcv == io.EOF {
			return nil
		}
		if errRcv != nil {
			log.Fatalf("Error while receiving request %v", errRcv)
		}
		//good request
		aNumber := req.GetNumber()
		if aNumber < 0 {
			//build a custom error
			return status.Errorf(codes.InvalidArgument,
				fmt.Sprintf("Negative number is not allowed: %d", aNumber))
		}
		//input is ok then build the response payload
		rootSq := math.Sqrt(float64(aNumber))
		stream.Send(&calculatorpb.RootSquareResponse{Input: aNumber,
			RootSquare: rootSq})
	}
}

func main() {
	log.Println("calculator server is started..")
	l, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("error while creating listener %v", err)
	}
	//register the service
	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &calcServer{})
	log.Printf("Service is registered %v", s)

	if err := s.Serve(l); err != nil {
		log.Fatalf("Error to server %v", err)
	}
}

func doSum(op *calculatorpb.Operand) int32 {
	res := op.GetFirstOperand() + op.SecondOperand
	return res
}
func doDecompose(aNumber int, stream calculatorpb.CalculatorService_DecomposeServer) {
	q := aNumber
	i := 2
	for {
		//when to stop
		//the i should reach the half of original number

		if q%i == 0 {
			q = q / i

			stream.Send(&calculatorpb.DecomposeResponse{PrimeNumber: int64(i)})
			if q == 1 {
				break
			}
		} else {
			//adjust divider to the next prime number
			//use the specific algo

			i = nextPrime(i)
			log.Printf("diviser increased to %d\n", i)
		}

	}
}
func doAverage(stream calculatorpb.CalculatorService_AverageServer) error {
	avg := 0.0
	allVal := 0.0
	counter := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			//take the average
			//build the response
			avg = allVal / float64(counter)

			return stream.SendAndClose(&calculatorpb.AverageResponse{Average: avg})
		}
		if err != nil {
			log.Fatalf("error while reading client stream %v", err)
		}
		allVal += req.GetAvgInput().GetAvgInput()
		counter++
	}
}

func isPrime(aNumber int) bool {
	if aNumber == 2 {
		return true
	}
	//if even
	if aNumber > 2 && aNumber%2 == 0 {
		return false
	}
	//if odd
	for i := 3; float64(i) <= math.Sqrt(float64(aNumber)); i += 2 {
		if aNumber%i == 0 {
			return false
		}
	}
	return true
}

func nextPrime(aNumber int) int {
	i := aNumber
	for {
		i++
		if isPrime(i) {
			log.Printf("got %d\n", i)
			return i
		}

	}
}
