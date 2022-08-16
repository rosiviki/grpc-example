package main

import (
	"example.com/rpcTest/greeting"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("failed to close connection: %s", err)
		}
	}(conn)

	c := greeting.NewGreeterClient(conn)

	i := 0
	for {
		if i > 10 {
			return
		}
		now := time.Now()
		response, err := c.SayHello(context.Background(), &greeting.HelloRequest{Name: "Hello From Client!"})
		if err != nil {
			log.Fatalf("Error when calling SayHello: %s", err)
		}
		log.Printf("%v\n", time.Since(now))
		log.Printf("Response from server: %s", response.Message)
		i++
	}
}
