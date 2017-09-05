package main

import (
	"log"
	"os"

	"golang.org/x/net/context"

	pb "github.com/satyamz/grpc-practice/add"

	"strconv"

	"google.golang.org/grpc"
)

const (
	address = "localhost:58888"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAdderClient(conn)
	var number1, number2 int64
	if len(os.Args) > 1 {
		number1, err = strconv.ParseInt(os.Args[1], 10, 32)
		if err != nil {
			log.Fatalf("Unable to parse integer: %v", err)
		}
		number2, err = strconv.ParseInt(os.Args[2], 10, 32)
		if err != nil {
			log.Fatalf("Unable to parse integer: %v", err)
		}
	}
	NumbersInstance := &pb.Numbers{
		Number1: number1,
		Number2: number2,
	}

	r, err := c.Add(context.Background(), NumbersInstance)
	if err != nil {
		log.Fatalf("Could not add due to: %v", err)
	}
	log.Printf("Addition is: %v", r.Addition)
}
