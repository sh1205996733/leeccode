package bit

import (
	"fmt"
	"testing"
)

func TestReverseBits(t *testing.T) {
	//str := fmt.Sprintf("%b", 00000010100101000001111010011100)
	//fmt.Println(str)

	fmt.Println(reverseBits(uint32(43261596)))
}

func TestAddBinary(t *testing.T) {
	cases := []struct {
		A, B   string
		Expect string
	}{
		{"1010", "1011", "10101"},
		{"11", "11", "110"},
		{"11", "1", "100"},
		{"0", "1", "1"},
		{"1", "1", "10"},
		{"11", "10111", "11010"},
		{"11", "10011", "10110"},
	}
	for _, c := range cases {
		if !(addBinary(c.A, c.B) == c.Expect) {
			fmt.Println(c.A, c.B, "测试不能通过")
			return
		}
	}
	fmt.Println("测试通过")
}
