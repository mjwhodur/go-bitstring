package bitstring

import "testing"

func TestBitString_AsBinaryString(t *testing.T) {
	bitString := BitString{bitArray: []uint8{0, 0, 0, 0}}

	s := bitString.AsBinaryString()

	if s != "0b0000" {
		t.Fail()
	}

	bitString = BitString{bitArray: []uint8{1, 1, 1, 1}}

	s = bitString.AsBinaryString()

	if s != "0b1111" {
		t.Fail()
	}
}

func TestBitString_AsHexString(t *testing.T) {
	bitString := BitString{bitArray: []uint8{0, 0, 0, 0}}

	s := bitString.AsHexString()

	if s != "0x0" {
		t.Fail()
	}

	bitString = BitString{bitArray: []uint8{1, 1, 1, 1}}

	s = bitString.AsHexString()

	if s != "0xf" {
		t.Fail()
	}

	bitString = BitString{bitArray: []uint8{1, 0, 0, 1}}

	s = bitString.AsHexString()

	if s != "0x9" {
		t.Fail()
	}
}

func TestBitString_AsOctString(t *testing.T) {

}

func TestBitString_AsString(t *testing.T) {

}

func TestBitString_BitLength(t *testing.T) {

}

func TestBitString_BitStringToIBCD(t *testing.T) {

}

func TestBitString_BitStringToICID(t *testing.T) {

}

func TestBitString_BitStringToTBCD(t *testing.T) {

}

func TestByteToBitArray(t *testing.T) {

}

func TestBitString_BitStringToUint(t *testing.T) {

}

func TestBitString_Substring(t *testing.T) {

}

func TestFromBytes(t *testing.T) {

}
func TestFromStream(t *testing.T) {

}

func BenchmarkByteToBitArray(b *testing.B) {

}

func BenchmarkFromStream(b *testing.B) {

}
