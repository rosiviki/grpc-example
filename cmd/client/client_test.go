package main

import (
	"context"
	"example.com/rpcTest/greeting"
	"google.golang.org/grpc"
	"log"
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type ClientSuite struct {
	conn          *grpc.ClientConn
	greeterClient greeting.GreeterClient
}

var _ = Suite(&ClientSuite{})

// Set up client connection and greeter client before suite execution
func (s *ClientSuite) SetUpSuite(c *C) {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	s.conn = conn
	s.greeterClient = greeting.NewGreeterClient(conn)
}

// Clean up after suite execution
func (s *ClientSuite) TearDownSuite(c *C) {
	err := s.conn.Close()
	if err != nil {
		log.Fatalf("failed to close connection: %s", err)
	}
}

func (s *ClientSuite) TestSayHello(c *C) {
	response, err := s.greeterClient.SayHello(context.Background(), &greeting.HelloRequest{Name: "Hello From Client!"})
	c.Assert(err, IsNil)
	c.Assert(response.Message, Equals, "Hello")
}

func (s *ClientSuite) BenchmarkSayHello(c *C) {
	for i := 0; i < c.N; i++ {
		response, err := s.greeterClient.SayHello(context.Background(), &greeting.HelloRequest{Name: "Hello From Client!"})
		c.Assert(err, IsNil)
		c.Assert(response.Message, Equals, "Hello")
	}
}
