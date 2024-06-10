package types

import (
	"crypto/sha256"

	pb "google.golang.org/protobuf/proto"
	"woyteck.pl/blocker/crypto"
	"woyteck.pl/blocker/proto"
)

func SignTransaction(privKey *crypto.PrivateKey, tx *proto.Transaction) crypto.Signature {
	return *privKey.Sign(HashTransaction(tx))
}

func HashTransaction(tx *proto.Transaction) []byte {
	b, err := pb.Marshal(tx)
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256(b)

	return hash[:]
}

func VerifyTransaction(tx *proto.Transaction) bool {
	for _, input := range tx.Inputs {
		var (
			sig    = crypto.SignatureFromBytes(input.Signature)
			pubKey = crypto.PublicKeyFromBytes(input.PublicKey)
		)
		// TODO: make sure we don't run into problems after verification
		// cause we've set the signature to nil
		input.Signature = nil
		if !sig.Verify(pubKey, HashTransaction(tx)) {
			return false
		}
	}

	return true
}
