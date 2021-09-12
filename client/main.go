package main

import (
	"context"
	"log"
	"time"

	"github.com/ptflp/grpc-example/types"

	"github.com/ptflp/gomapper"

	pb "github.com/ptflp/grpc-example/pb"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

// ur input structure
type YourBusinessType struct {
	Name string      `mapper:"name"`
	Test interface{} `mapper:"test"`
}

type HelloRequestTemp struct {
	Name string        `json:"name" mapper:"name"`
	Test types.AnyType `json:"test" mapper:"test"`
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Create ur structure
	urType := YourBusinessType{
		Name: defaultName,
	}

	// use any type here which in ur interface{}
	urType.Test = struct {
		Age        int
		SomeString string
	}{
		Age:        16,
		SomeString: "asfafasfasfasf",
	}

	var helloReqTemp HelloRequestTemp

	// Map ur type to temp request type we need cast interface to AnyType (custom type)
	err = gomapper.MapStructs(&helloReqTemp, &urType)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// Then map temp request type to protobuf request type
	var grpcReq pb.HelloRequest
	err = gomapper.MapStructs(&grpcReq, &helloReqTemp, "json")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &grpcReq)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s, %s", r.GetMessage(), r.GetTest())
}
