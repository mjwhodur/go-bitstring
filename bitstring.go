// Package bitstring provides functions for parsing and manipulation of
//
package bitstring

import (
	"bytes"
	"math"
)

// Base BitString structure consisting of bytes representing
type BitString struct {
	BitArray []uint8
}

// BUG(mjwhodur): Not yet working properly
func (bs *BitString) AsString() string {
	var s = ""
	for _, bit := range bs.BitArray {
		s = s + string(bit)
	}
	return s
}

//BUG(mjwhodur): Not yet working properly
func (bs *BitString) AsBinaryString() string {
	return ""
}

//BUG(mjwhodur): Not yet working properly
func (bs *BitString) AsHexString() string {
	return ""
}

//BUG(mjwhodur): Not yet working properly
func (bs *BitString) AsOctString() string {
	return ""
}

// Returns subsstring of one BitString as new BitString
func (bs *BitString) Substring(start int, end int) BitString {
	return BitString{BitArray: bs.BitArray[start:end]}
}

//BUG(mjwhodur): Not yet working properly
func (bs *BitString) BitLength() int {
	return len(bs.BitArray)
}

//BUG(mjwhodur): Not yet working properly
func (bs *BitString) BitStringToUint(start int, end int) int {
	return 0
}

//BUG(mjwhodur): Not yet working properly
func (bs *BitString) BitStringToICID(start int, end int) (error, string) {
	return nil, ""
}

//BUG(mjwhodur): Not yet working properly
func (bs *BitString) BitStringToIBCD(start int, end int) (error, string) {
	return nil, ""
}

// Telephony Binary Coded Decimal (TBCD) helper function provide assistance in
//converting TBCD strings from BitString slice form as in standard ITU-T Q.825.
//BUG(mjwhodur): Not yet working properly
func (bs *BitString) BitStringToTBCD(start int, end int) (error, string) {
	return nil, ""
}

// Appends byte to BitArray as bits
func (bs *BitString) appendByte(rune byte) error {
	for _, bitVal := range ByteToBitArray(rune).BitArray {
		bs.BitArray = append(bs.BitArray, bitVal)
	}

	return nil
}

// Creates BitString struct from byte slice.
func FromBytes(bytes []byte) (error, BitString) {
	bs := BitString{}
	for _, rune := range bytes {
		err := bs.appendByte(rune)
		if err != nil {
			return err, BitString{}
		}
	}
	return nil, bs
}

// Creates BitString struct from buffered data.
func FromStream(buffer bytes.Buffer) (error, BitString) {
	return nil, BitString{}
}

// Creates BitString struct from single byte.
func ByteToBitArray(b byte) BitString {
	var revbitSlice []uint8
	var bitSlice []uint8
	for x := 0; x < 8; x++ {
		c := byte(math.Pow(2, float64(x)))
		if c < b {
			b = b - c
			revbitSlice = append(revbitSlice, 1)

		} else {
			revbitSlice = append(revbitSlice, 0)
		}
	}

	for i := len(revbitSlice) - 1; i == 0; i-- {
		bitSlice = append(bitSlice, revbitSlice[i])
	}
	return BitString{BitArray: bitSlice}
}
