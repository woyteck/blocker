package types

import (
	"crypto/sha256"

	pb "google.golang.org/protobuf/proto"
	"woyteck.pl/blocker/crypto"
	"woyteck.pl/blocker/proto"
)

func VerifyBlock(b *proto.Block) bool {
	if len(b.PublicKey) != crypto.PubKeyLen {
		return false
	}

	if len(b.Signature) != crypto.SignatureLen {
		return false
	}

	sig := crypto.SignatureFromBytes(b.Signature)
	pubKey := crypto.PublicKeyFromBytes(b.PublicKey)
	hash := HashBlock(b)

	return sig.Verify(pubKey, hash)
}

func SignBlock(privKey *crypto.PrivateKey, b *proto.Block) *crypto.Signature {
	hash := HashBlock(b)
	sig := privKey.Sign(hash)
	b.PublicKey = privKey.Public().Bytes()
	b.Signature = sig.Bytes()

	return sig
}

func HashBlock(block *proto.Block) []byte {
	return HashHeader(block.Header)
}

func HashHeader(header *proto.Header) []byte {
	b, err := pb.Marshal(header)
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256(b)

	return hash[:]
}
