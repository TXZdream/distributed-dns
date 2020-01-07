package dns

import (
	mygrpc "distributed-dns/grpc"
	"github.com/willf/bitset"
	"golang.org/x/crypto/ripemd160"
	"google.golang.org/grpc"
)

func toString(id *bitset.BitSet) string {
	var ret string
	for i := uint(0); i < id.Len(); i++ {
		if data := id.Test(i); data {
			ret += "1"
		} else {
			ret += "0"
		}
	}
	return ret
}

func ToBitArr(id string) *bitset.BitSet {
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

func CalculateHash(raw string) (*bitset.BitSet, error) {
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

func dialGrpc(addr string) (mygrpc.KademilaClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return mygrpc.NewKademilaClient(conn), nil
}

// compare 计算a、b分别和target之间的距离远近
// 如果a更近，返回-1；b更近返回1；一样近返回0
func compare(target, a, b string) int {
	if !(len(a) == len(b) && len(a) == len(target)) {
		return -2
	}
	a = xor(a, target)
	b = xor(b, target)
	for i := range a {
		if a[i] > b[i] {
			return 1
		} else if a[i] < b[i] {
			return -1
		}
	}
	return 0
}

func xor(a, b string) string {
	ret := ""
	if len(a) != len(b) {
		return ""
	}
	for i := range a {
		if a[i] == b[i] {
			ret += "0"
		} else {
			ret += "1"
		}
	}
	return ret
}
