package node

import (
	"context"
	"fmt"

	"google.golang.org/grpc/peer"
	"woyteck.pl/blocker/proto"
)

type Node struct {
	version string
	// peers   map[net.Addr]*grpc.ClientConn
	proto.UnimplementedNodeServer
}

func New() *Node {
	return &Node{
		version: "blocker-0.1",
	}
}

func (n *Node) Handshake(ctx context.Context, v *proto.Version) (*proto.Version, error) {
	ourVersion := &proto.Version{
		Version: n.version,
		Height:  100,
	}

	p, _ := peer.FromContext(ctx)
	fmt.Printf("received version from %s: %+v\n", v, p.Addr)

	return ourVersion, nil
}

func (n *Node) HandleTransaction(ctx context.Context, tx *proto.Transaction) (*proto.Ack, error) {
	peer, _ := peer.FromContext(ctx)
	fmt.Println("received tx from:", peer)
	return &proto.Ack{}, nil
}
