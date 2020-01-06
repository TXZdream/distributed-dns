package dns

import "github.com/Workiva/go-datastructures/bitarray"

func toString(id bitarray.BitArray) (string, error) {
	var ret string
	for i := uint64(0); i < id.Capacity(); i++ {
		if data, err := id.GetBit(i); err != nil {
			return "", err
		} else if data {
			ret += "1"
		} else {
			ret += "0"
		}
	}
	return ret, nil
}

func toBitArr(id string) bitarray.BitArray {
	ret := bitarray.NewBitArray(uint64(len(id)))
	for i, v := range id {
		if v == '0' {
			ret.ClearBit(uint64(i))
		} else {
			ret.SetBit(uint64(i))
		}
	}
	return ret
}
