package day16

import (
	"aoc2021/assert"
	"bytes"
	"encoding/hex"
	"testing"
)

func TestParsePackage1(t *testing.T) {
	b, err := hex.DecodeString("8A004A801A8002F478")
	if err != nil {
		t.Error(err)
	}
	bs := newBs(bytes.NewReader(b))

	bPack, err := parsePacket(bs)

	if err != nil {
		t.Error(err)
	}

	if bPack.version != 4 {
		t.Errorf("Expected 4, got %v\n", bPack.version)
	}

	if len(bPack.children) != 1 {
		t.Errorf("Expected 1 child, got %v\n", len(bPack.children))
	}
}

func parseHex(t *testing.T, hx string) *bitsPackage {
	t.Helper()
	b, err := hex.DecodeString(hx)
	if err != nil {
		t.Error(err)
	}
	bs := newBs(bytes.NewReader(b))
	bPack, err := parsePacket(bs)
	if err != nil {
		t.Error(err)
	}

	return bPack
}

func TestParsePackage2(t *testing.T) {
	b := parseHex(t, "EE00D40C823060")

	if b.version != 7 {
		t.Errorf("Expected 7, got %v\n", b.version)
	}
	if b.id != 3 {
		t.Errorf("Expected 3, got %v\n", b.id)
	}

	childCount := len(b.children)
	if childCount != 3 {
		t.Errorf("Expected child count to be 3 but got %v\n", childCount)
	}

	for i, c := range b.children {
		if c.literal != uint(i)+1 {
			t.Errorf("Expected child at index %v to have literal %v but got %v", i, i+1, c.literal)
		}
	}

}
func TestParsePackage3(t *testing.T) {
	b := parseHex(t, "38006F45291200")

	assert.Equal(t, b.version, 1)
	assert.Equal(t, b.id, 6)

	assert.Equal(t, len(b.children), 2)

	assert.Equal(t, b.children[0].literal, 10)
	assert.Equal(t, b.children[1].literal, 20)
}
func TestParsePackage4(t *testing.T) {
	b := parseHex(t, "D2FE28")

	assert.Equal(t, b.version, 6)
	assert.Equal(t, b.id, 4)

	assert.Equal(t, b.literal, 2021)
}

func TestSums(t *testing.T) {
	tests := map[string]int{
		"8A004A801A8002F478":             16,
		"620080001611562C8802118E34":     12,
		"C0015000016115A2E0802F182340":   23,
		"A0016C880162017C3686B18A3D4780": 31,
	}

	for hex, expected := range tests {
		packet := parseHex(t, hex)
		sum := 0

		sumVersions(packet, &sum)
		if sum != expected {
			t.Errorf("Expected a sum of %v but got %v", expected, sum)
		}
	}

	// b := parseHex(t, "A0016C880162017C3686B18A3D4780")

	// if sum != 31 {
	// 	t.Errorf("Expected 31, got %v", sum)
	// }

}

// func TestParseLiteral(t *testing.T) {
// 	strconv.
// 	bs := newBs(bytes.NewReader([]byte))

// }
