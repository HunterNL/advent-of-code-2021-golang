package day16

import "io"

type BitsReaderStruct struct {
	currentBits   int
	currentBuffer uint16
	scanner       io.ByteScanner
	totalScanSize int
}

type BitScannerInterface interface {
	readBits(n int) byte
	totalBitCount() int
}

func newBs(s io.ByteScanner) *BitsReaderStruct {
	a := new(BitsReaderStruct)
	a.scanner = s
	return a
}

func (b *BitsReaderStruct) readBits(n int) byte {
	if n > b.currentBits {
		b.currentBuffer = b.currentBuffer << 8
		b.currentBits = b.currentBits + 8
		newbyte, err := b.scanner.ReadByte()
		if err != nil {
			panic(err)
		}
		b.currentBuffer = b.currentBuffer | uint16(newbyte)
	}
	b.totalScanSize += n
	ret := b.currentBuffer >> (b.currentBits - n)
	b.currentBits = b.currentBits - n
	shiftBy := 16 - uint16(b.currentBits)
	b.currentBuffer = (b.currentBuffer << shiftBy) >> shiftBy //11 there

	return byte(ret)
}
func (b *BitsReaderStruct) totalBitCount() int {
	return b.totalScanSize
}
