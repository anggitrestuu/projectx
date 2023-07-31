package core

import (
	"testing"
	"time"

	"github.com/anggitrestuu/projectx/crypto"
	"github.com/anggitrestuu/projectx/types"
	"github.com/stretchr/testify/assert"
)

func randomBlock(height uint32) *Block {
	header := &Header{
		Version:       1,
		PrevBlockHash: types.RandomHash(),
		Height:        height,
		Timestamp:     uint64(time.Now().UnixNano()),
	}

	tx := Transaction{
		Data: []byte("hello world"),
	}

	return NewBlock(header, []Transaction{tx})

}

func TestSignBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	block := randomBlock(0)

	assert.Nil(t, block.Sign(privKey))
	assert.NotNil(t, block.Signature)
}

func TestVerifyBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	block := randomBlock(0)

	assert.Nil(t, block.Sign(privKey))
	assert.Nil(t, block.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	block.Validator = otherPrivKey.PublicKey()
	// should fail because the public key is not the same
	assert.NotNil(t, block.Verify())

	block.Height = 100
	assert.NotNil(t, block.Verify())
}
