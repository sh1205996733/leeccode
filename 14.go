package main

import "fmt"

func main() {
	strs := []string{"car", "bd"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	count := len(strs)
	min := 0
	for i := 0; i < count; i++ {
		if len([]byte(strs[min])) > len([]byte(strs[i])) {
			min = i
		}
	}
	index := 0
	prefix := strs[min]
	for index < len([]byte(prefix)) { //strs元素不为空
		c := string(prefix[index])
		for i := 0; i < count; i++ {
			if c != string(strs[i][index]) {
				return prefix[0:index]
			}
		}
		index++
	}
	return prefix
}
