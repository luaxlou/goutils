package grpcclient

import (
	"log"
	"time"

	"google.golang.org/grpc"
)

type Client struct {
	conn *grpc.ClientConn

	address string
}

func New(address string) *Client {

	return &Client{
		address: address,
	}
}

func (c *Client) GetConn() *grpc.ClientConn {

	return c.conn
}

func (c *Client) Connect(address string) error {

	log.Println("grpc connecting to", address)

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)

		time.Sleep(time.Millisecond * 100)
		return c.Connect(address)
	}

	c.conn = conn

	return nil
}

func (c *Client) Disconnect() {

	if c != nil && c.conn != nil {
		c.conn.Close()

	}

}
