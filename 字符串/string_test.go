package string

import (
	"fmt"
	"testing"
)

func TestTitleToNumber(t *testing.T) {
	columnTitle := "BA"
	fmt.Println(titleToNumber(columnTitle))
}
