package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"woyteck.pl/blocker/crypto"
	"woyteck.pl/blocker/node"
	"woyteck.pl/blocker/proto"
	"woyteck.pl/blocker/util"
)

func main() {
	makeNode(":3000", []string{}, true)
	time.Sleep(time.Second)
	makeNode(":4000", []string{":3000"}, false)
	time.Sleep(4 * time.Second)
	makeNode(":5000", []string{":4000"}, false)

	for {
		time.Sleep(time.Millisecond * 100)
		makeTransaction()
	}
}

func makeNode(listenAddr string, bootstrapNodes []string, isValidator bool) *node.Node {
	cfg := node.ServerConfig{
		Version:    "blocker-0.1",
		ListenAddr: listenAddr,
	}
	if isValidator {
		cfg.PrivateKey = crypto.GeneratePrivateKey()
	}
	n := node.NewNode(cfg)
	go n.Start(listenAddr, bootstrapNodes)

	return n
}

// func makeHandshake() {
// 	opts := []grpc.DialOption{
// 		grpc.WithTransportCredentials(insecure.NewCredentials()),
// 	}
// 	client, err := grpc.NewClient(":3000", opts...)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	version := &proto.Version{
// 		Version:    "blocker-0.1",
// 		Height:     1,
// 		ListenAddr: ":4000",
// 	}

// 	c := proto.NewNodeClient(client)
// 	_, err = c.Handshake(context.TODO(), version)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func makeTransaction() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	client, err := grpc.NewClient(":3000", opts...)
	if err != nil {
		log.Fatal(err)
	}

	privKey := crypto.GeneratePrivateKey()

	tx := &proto.Transaction{
		Version: 1,
		Inputs: []*proto.TxInput{
			{
				PrevTxHash:   util.RandomHash(),
				PrevOutIndex: 0,
				PublicKey:    privKey.Public().Bytes(),
			},
		},
		Outputs: []*proto.TxOutput{
			{
				Amount:  99,
				Address: privKey.Public().Address().Bytes(),
			},
		},
	}

	c := proto.NewNodeClient(client)
	_, err = c.HandleTransaction(context.TODO(), tx)
	if err != nil {
		log.Fatal(err)
	}
}
