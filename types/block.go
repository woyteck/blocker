package types

import (
	"crypto/sha256"

	pb "google.golang.org/protobuf/proto"
	"woyteck.pl/blocker/crypto"
	"woyteck.pl/blocker/proto"
)

func SignBlock(privKey *crypto.PrivateKey, b *proto.Block) *crypto.Signature {
	return privKey.Sign(HashBlock(b))
}

func HashBlock(block *proto.Block) []byte {
	b, err := pb.Marshal(block)
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256(b)

	return hash[:]
}
