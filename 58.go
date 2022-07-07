/**
 * @Author:shangyc
 * @Date: 2022/5/9 0009
 */

package main

import "fmt"

func main() {
	str := "luffy is still joyboy"
	fmt.Println(lengthOfLastWord(str))
}
func lengthOfLastWord(s string) (ans int) {
	index := len(s) - 1
	for s[index] == ' ' {
		index--
	}
	for index >= 0 && s[index] != ' ' {
		ans++
		index--
	}
	return
}
