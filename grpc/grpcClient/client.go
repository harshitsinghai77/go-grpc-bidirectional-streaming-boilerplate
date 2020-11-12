package main

import (
	"io"
	"log"

	pb "github.com/harshitsinghai/grpc/datafiles"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const address = "localhost:5001"

// ReceiveStream listens to the stream contents and use them
func ReceiveStream(client pb.MoneyTransactionClient, request *pb.TransactionRequest) {
	log.Println("Started listening to the server stream!")
	stream, err := client.MakeTransaction(context.Background(), request)
	if err != nil {
		log.Fatalf("%v.MakeTransaction(_) = _, %v", client, err)
	}
	//Listen to the stream of messages
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			// If there are no more messages, get out of loop
			break
		}
		if err != nil {
			log.Fatalf("%v.MakeTransaction(_) = _, %v", client, err)
		}
		log.Printf("Status: %v Operation: %v", response.Status, response.Description)
	}
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewMoneyTransactionClient(conn)

	// Prepare data. Get this from clients like Frontend or App
	from := "1234"
	to := "5678"
	amount := float32(1250)

	// r, err := c.MakeTransaction(context.Background(), &pb.TransactionRequest{From: from, To: to, Amount: amount})
	// if err != nil {
	// 	log.Fatalf("Could not transact: %v", err)
	// }
	// log.Printf("Transaction confirmed: %t", r.Confirmation)

	ReceiveStream(c, &pb.TransactionRequest{From: from, To: to, Amount: amount})
}
