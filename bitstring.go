package bitstring

import "bytes"

type BitString struct {
	BitArray []uint8
}

//
func (bs *BitString) AsString() string {
	return string(bs.BitArray)
}

func (bs *BitString) Length() int {
     return len(bs.BitArray)
}

func (bs *BitString) Uint(start int, end int) int {
	return 0
}

func (bs *BitString) ICID(start int, end int) (error, string) {
	return nil, ""
}

func (bs *BitString) IBCD(start int, end int) (error, string) {
	return nil, ""
}

func FromBytes(bytes []rune) (error, BitString) {
	return nil, BitString{}
}

func FromStream(buffer bytes.Buffer) (error, BitString) {
	return nil, BitString{}
}

