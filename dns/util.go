package dns

import (
	"github.com/willf/bitset"
	"golang.org/x/crypto/ripemd160"
)

func toString(id *bitset.BitSet) (string, error) {
	var ret string
	for i := uint(0); i < id.Len(); i++ {
		if data := id.Test(i); data {
			ret += "1"
		} else {
			ret += "0"
		}
	}
	return ret, nil
}

func toBitArr(id string) *bitset.BitSet {
	ret := bitset.New(uint(len(id)))
	for i, v := range id {
		if v == '0' {
			ret.Clear(uint(i))
		} else {
			ret.Set(uint(i))
		}
	}
	return ret
}

func calculateHash(raw string) (*bitset.BitSet, error) {
	rim := ripemd160.New()
	if _, err := rim.Write([]byte(raw)); err != nil {
		return nil, err
	}
	hash := rim.Sum(nil)
	ret := bitset.New(ripemd160.Size * 8)
	for i, v := range hash {
		for j := 0; j < 8; j++ {
			if bit := (v & (1 << (8 - j - 1))) == 1; bit {
				ret.Set(uint(i*8 + j))
			} else {
				ret.Clear(uint(i*8 + j))
			}
		}
	}
	return ret, nil
}
