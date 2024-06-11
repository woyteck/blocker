package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"woyteck.pl/blocker/node"
	"woyteck.pl/blocker/proto"
)

func main() {
	makeNode(":3000", []string{})
	time.Sleep(time.Second)
	makeNode(":4000", []string{":3000"})
	time.Sleep(4 * time.Second)
	makeNode(":5000", []string{":4000"})

	select {}
}

func makeNode(listenAddr string, bootstrapNodes []string) *node.Node {
	n := node.New()
	go n.Start(listenAddr, bootstrapNodes)

	return n
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
		Version:    "blocker-0.1",
		Height:     1,
		ListenAddr: ":4000",
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
