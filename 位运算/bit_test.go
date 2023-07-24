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
