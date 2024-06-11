package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"woyteck.pl/blocker/node"
	"woyteck.pl/blocker/proto"
)

func main() {
	node := node.New()

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}
	proto.RegisterNodeServer(grpcServer, node)
	fmt.Println("node running on:", ":3000")

	defer grpcServer.Stop()

	go func() {
		for {
			time.Sleep(2 * time.Second)
			// makeTransaction()
			makeHandshake()
		}
	}()
	grpcServer.Serve(ln)
}

func makeHandshake() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	client, err := grpc.NewClient(":3000", opts...)
	if err != nil {
		log.Fatal(err)
	}

	version := &proto.Version{
		Version: "blocker-0.1",
		Height:  1,
	}

	c := proto.NewNodeClient(client)
	_, err = c.Handshake(context.TODO(), version)
	if err != nil {
		log.Fatal(err)
	}
}

func makeTransaction() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	client, err := grpc.NewClient(":3000", opts...)
	if err != nil {
		log.Fatal(err)
	}

	tx := &proto.Transaction{
		Version: 1,
	}

	c := proto.NewNodeClient(client)
	_, err = c.HandleTransaction(context.TODO(), tx)
	if err != nil {
		log.Fatal(err)
	}
}
