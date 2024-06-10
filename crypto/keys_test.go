package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()

	assert.Equal(t, len(privKey.Bytes()), privKeyLen)
	assert.Equal(t, len(pubKey.Bytes()), pubKeyLen)
}

func TestNewPrivateKeyFromString(t *testing.T) {
	seed := "848db367bccc1b71bdfbc539f5d0114e262d0baf2e453f097f101541221d6c5c"
	privKey := NewPrivateKeyFromString(seed)
	assert.Equal(t, privKeyLen, len(privKey.Bytes()))
	address := privKey.Public().Address()
	assert.Equal(t, "e5223af1b4498c466a33d2ce8ef5244718a21031", address.String())
}

func TestPrivateKeySign(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()
	msg := []byte("foo bar baz")
	sig := privKey.Sign(msg)

	assert.True(t, sig.Verify(pubKey, msg))

	// test with invalid msg
	assert.False(t, sig.Verify(pubKey, []byte("foo")))

	// test with invalid pubKey
	privKey2 := GeneratePrivateKey()
	pubKey2 := privKey2.Public()
	assert.False(t, sig.Verify(pubKey2, msg))
}

func TestPublicKeyToAddress(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()
	address := pubKey.Address()

	assert.Equal(t, addressLen, len(address.Bytes()))
}
