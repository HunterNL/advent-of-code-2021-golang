package day16

import (
	"aoc2021/file"
	"bytes"
	"encoding/hex"
	"log"
	"math"
)

const PackageSum uint8 = 0
const PackageProduct uint8 = 1
const PackageMin uint8 = 2
const PackageMax uint8 = 3
const PackageLiteral uint8 = 4
const PackageGreater uint8 = 5
const PackageLesser uint8 = 6
const PackageEqualTo uint8 = 7

type bitsPackage struct {
	version  uint8
	id       uint8
	literal  uint
	children []bitsPackage
}

func parseLiteral(scanner BitScannerInterface) uint {
	var out uint = 0
	for {
		out = out << 4
		byte := (scanner).readBits(5)
		isLast := byte&0b0010000 == 0
		byte = byte & 0b0001111
		out = out | uint(byte)
		if isLast {
			break
		}
	}

	return out
}

func parsePacket(scanner BitScannerInterface) (b *bitsPackage, e error) {
	bs := scanner
	version := uint8(bs.readBits(3))
	id := uint8(bs.readBits(3))
	children := make([]bitsPackage, 0)
	var literal uint

	if id == PackageLiteral {
		literal = parseLiteral(scanner)
		// log.Printf("Literal: %v\n", literal)
	} else {
		//Operator
		lengthType := bs.readBits(1)

		if lengthType == 0 {
			bitLen := uint16(bs.readBits(7)) << 8
			bitLen = bitLen | uint16(bs.readBits(8))

			log.Printf("BitLen: %v\n", bitLen)

			curScan := bs.totalBitCount()

			for bs.totalBitCount() < curScan+int(bitLen) {
				child, err := parsePacket(scanner)
				if err != nil {
					panic(err)
				}
				children = append(children, *child)
			}
		} else {
			childCount := uint16(bs.readBits(3)) << 8
			childCount = childCount | uint16(bs.readBits(8))

			log.Printf("ChildCount: %v\n", childCount)

			for i := 0; i < int(childCount); i++ {
				child, err := parsePacket(scanner)
				if err != nil {
					panic(err)
				}
				children = append(children, *child)
			}
		}

	}

	return &bitsPackage{
		version:  version,
		id:       id,
		literal:  literal,
		children: children,
	}, nil
}

func sumVersions(p *bitsPackage, sum *int) {
	(*sum) += int(p.version)

	for _, c := range p.children {
		sumVersions(&c, sum)
	}
}

func resolvePacket(p *bitsPackage) int {
	if p.id == PackageSum {
		sum := 0
		for _, bp := range p.children {
			sum += resolvePacket(&bp)
		}
		return sum
	}
	if p.id == PackageProduct {
		product := 1
		for _, bp := range p.children {
			product = product * resolvePacket(&bp)
		}
		return product
	}
	if p.id == PackageMin {
		min := math.MaxInt
		for _, bp := range p.children {
			val := resolvePacket(&bp)
			if val < min {
				min = val
			}
		}
		return min
	}
	if p.id == PackageMax {
		max := math.MinInt
		for _, bp := range p.children {
			val := resolvePacket(&bp)
			if val > max {
				max = val
			}
		}
		return max
	}
	if p.id == PackageGreater {
		if resolvePacket(&p.children[0]) > resolvePacket(&p.children[1]) {
			return 1
		} else {
			return 0
		}
	}
	if p.id == PackageLesser {
		if resolvePacket(&p.children[0]) < resolvePacket(&p.children[1]) {
			return 1
		} else {
			return 0
		}
	}
	if p.id == PackageEqualTo {
		if resolvePacket(&p.children[0]) == resolvePacket(&p.children[1]) {
			return 1
		} else {
			return 0
		}
	}

	if p.id == PackageLiteral {
		return int(p.literal)
	}
	panic("Unknown package")
}

func Solve() (int, int) {
	line := file.ReadFile("./day16/input.txt")[0]

	b, err := hex.DecodeString(line)
	if err != nil {
		panic(err)
	}

	bs := newBs(bytes.NewReader(b))
	packet, err := parsePacket(bs)
	if err != nil {
		panic(err)
	}

	sum := 0

	sumVersions(packet, &sum)

	res := resolvePacket(packet)

	log.Printf("P: %v\n versionSum: %v\n Evaluated: %v\n", packet, sum, res)

	return sum, res

}
