package string

import (
	"fmt"
	"testing"
)

func TestTitleToNumber(t *testing.T) {
	columnTitle := "BA"
	fmt.Println(titleToNumber(columnTitle))
}

func TestFizzBuzz(t *testing.T) {
	fmt.Println(fizzBuzz(10))
}

func TestFirstUniqChar(t *testing.T) {
	fmt.Println(firstUniqChar("leetcode"))
}

func TestFrequencySort(t *testing.T) {
	fmt.Println(frequencySort("leetcode"))
}

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"bcar", "bd"}
	fmt.Println(longestCommonPrefix(strs))
}

func TestLengthOfLastWord(t *testing.T) {
	str := "luffy is still joyboy"
	fmt.Println(lengthOfLastWord(str))
}
