package types

import "fmt"

type Address [20]uint8

func (a Address) ToSlice() []byte {
	b := make([]byte, 20)
	for i := 0; i < 20; i++ {
		b[i] = byte(a[i])
	}
	return b
}

func NewAddressFromBytes(b []byte) Address {
	if len(b) != 20 {
		msg := fmt.Sprintf("Invalid address length: %d", len(b))
		panic(msg)
	}

	var val [20]uint8
	for i := 0; i < 20; i++ {
		val[i] = uint8(b[i])
	}

	return Address(val)

}
