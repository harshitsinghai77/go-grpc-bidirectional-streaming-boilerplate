package main

import (
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/harshitsinghai/grpc/datafiles"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":5001"
const noOfSteps = 5

// server is used to create MoneyTransactionServer.
type server struct{}

// MakeTransaction implements MoneyTransactionServer.MakeTransaction
func (s *server) MakeTransaction(in *pb.TransactionRequest, stream pb.MoneyTransaction_MakeTransactionServer) error {
	log.Printf("Got request for money transfer..")
	log.Printf("Amount: %f, From A/c:%s, To A/c:%s", in.Amount, in.From, in.To)

	// Do database logic here...

	for i := 0; i < noOfSteps; i++ {
		// Simulating I/O or Computation process using sleep........
		// Usually this will be saving money transfer details in DB or
		// talk to the third party API

		time.Sleep(time.Second * 5)
		// Once task is done, send the successful message back to the client
		if err := stream.Send(&pb.TransactionResponse{Status: "good", Step: int32(i), Description: fmt.Sprintf("Description of step %d", int32(i))}); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, "status", err)
		}
	}

	log.Printf("Successfully transfered amount $%v from %v to %v", in.Amount, in.From, in.To)
	return nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMoneyTransactionServer(s, &server{})
	reflection.Register(s)

	fmt.Println("Listening to server")
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
