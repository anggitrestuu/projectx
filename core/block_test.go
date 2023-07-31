package core

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/anggitrestuu/projectx/types"
	"github.com/stretchr/testify/assert"
)

func TestHeader(t *testing.T) {
	h := &Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: uint64(time.Now().UnixNano()),
		Height:    10,
		Nonce:     989394,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, h.EncodeBinary(buf))

	hDecode := &Header{}
	assert.Nil(t, hDecode.DecodeBinary(buf))
	assert.Equal(t, h, hDecode)

}

func TestHeaderEncodeDecode(t *testing.T) {
	h := &Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: uint64(time.Now().UnixNano()),
		Height:    10,
		Nonce:     989394,
	}

	buff := &bytes.Buffer{}
	assert.Nil(t, h.EncodeBinary(buff))

	hDecode := &Header{}
	assert.Nil(t, hDecode.DecodeBinary(buff))
	assert.Equal(t, h, hDecode)
	fmt.Printf("%+v\n", hDecode)
}

func TestBlockEncodeDecode(t *testing.T) {
	h := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: uint64(time.Now().UnixNano()),
			Height:    10,
			Nonce:     989394,
		},
		Transactions: nil,
	}

	buff := &bytes.Buffer{}
	assert.Nil(t, h.EncodeBinary(buff))

	hDecode := &Block{}
	assert.Nil(t, hDecode.DecodeBinary(buff))
	assert.Equal(t, h, hDecode)
	fmt.Printf("%+v\n", hDecode)

}

func TestBlockHash(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: uint64(time.Now().UnixNano()),
			Height:    10,
			Nonce:     989394,
		},
		Transactions: nil,
	}

	hash := b.Hash()
	fmt.Printf("%x\n", hash)
	assert.Equal(t, hash, b.Hash())
}
