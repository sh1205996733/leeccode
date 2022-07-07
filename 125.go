/**
 * @Author:注释中编瞎话
 * @Date: 2022/3/30 0030
 */

package main

import (
	"fmt"
	"strings"
)

func main() {
	//str1 :="A man, a plan, a canal: Panama"
	//fmt.Printf("%d %d  %d %d\n",'a','z','A','Z')
	instants := []string{
		"  A man, a plan, a canal: Panama",
		"race a car",
	}
	for i := 0; i < len(instants); i++ {
		fmt.Println(instants[i], "结果", isPalindrome0(instants[i]))
	}
	//fmt.Println(isPalindrome(str1))
}

//48-57 ---0-9
//65-90
//97-122
//leetcode submit region begin(Prohibit modification and deletion)
//在原字符串上直接判断
func isPalindrome0(str string) bool {
	str = strings.ToLower(str)
	length := len(str)
	head, tail := 0, length-1
	for head < tail {
		if !isVaild(str[head]) {
			head++
		} else if !isVaild(str[tail]) {
			tail--
		} else if str[head] == str[tail] {
			head++
			tail--
		} else {
			return false
		}
	}
	return true
}
func isPalindrome(str string) bool {
	tmp := make([]byte, len(str))
	index := 0
	for i := 0; i < len(str); i++ {
		if isVaild(str[i]) {
			if str[i] >= 'A' && str[i] <= 'Z' {
				tmp[index] = str[i] + 32
			} else {
				tmp[index] = str[i]
			}
			index++
		}
	}
	head, tail := 0, index-1
	for head < tail {
		if tmp[head] == tmp[tail] {
			head++
			tail--
		} else {
			return false
		}
	}
	return true
}

func isVaild(c byte) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z')
}

//leetcode submit region end(Prohibit modification and deletion)
