// Package bitstring provides functions for parsing and manipulation of
// bits in byte streams and bytearrays. Note, that this package is not so
// memory-efficient.
package bitstring

import (
	"bytes"
	"math"
	"strconv"
)

// Base BitString structure consisting of bytes representing
type BitString struct {
	bitArray []uint8
}

// Returns BitString.bitArray as string without further parsing.
func (bs *BitString) AsString() string {
	var s = ""
	for _, bit := range bs.bitArray {
		s = s + strconv.Itoa(int(bit))
	}
	return s
}

//Returns BitString value formatted as "0b...".
func (bs *BitString) AsBinaryString() string {
	return "0b" + bs.AsString()
}

//Returns BitString as Hex String with padding.
func (bs *BitString) AsHexString() string {
	padding := "0x"
	var sl []uint8
	for i := bs.BitLength() % 4; i > 0; i-- {
		sl = append(sl, 0)
	}
	for _, elem := range bs.bitArray {
		sl = append(sl, elem)
	}
	for i := 0; i < len(sl); i++ {
		k, l, m, n := sl[i], sl[i+1], sl[i+2], sl[i+3]
		l = l * 2
		m = l * 4
		n = n * 8
		switch k + l + m + n {
		case 10:
			padding = padding + "a"
		case 11:
			padding = padding + "b"
		case 12:
			padding = padding + "c"
		case 13:
			padding = padding + "d"
		case 14:
			padding = padding + "e"
		case 15:
			padding = padding + "f"
		default:
			padding = padding + strconv.Itoa(int(k+l+m+n))
		}

	}
	return padding
}

//BUG(mjwhodur): Not yet working properly
func (bs *BitString) AsOctString() string {
	return ""
}

// Returns subsstring of one BitString as new BitString
func (bs *BitString) Substring(start int, end int) BitString {
	return BitString{bitArray: bs.bitArray[start:end]}
}

//BUG(mjwhodur): Not yet working properly
func (bs *BitString) BitLength() int {
	return len(bs.bitArray)
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

// Appends byte to bitArray as bits
func (bs *BitString) appendByte(rune byte) error {
	for _, bitVal := range ByteToBitArray(rune).bitArray {
		bs.bitArray = append(bs.bitArray, bitVal)
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
	for x := 7; x > 0; x-- {
		c := uint8(math.Pow(2, float64(x-1)))
		if c <= b {

			b = b - c
			revbitSlice = append(revbitSlice, 1)

		} else {
			revbitSlice = append(revbitSlice, 0)
		}
	}

	return BitString{bitArray: revbitSlice}
}
