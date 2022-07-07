package main

import "fmt"

func main() {
	strs := "LLALL"
	fmt.Println(checkRecord(strs))
}

/**
 *	设置A、L计数器acount、lcount
 * 	遍历字符串：
 *		如果是A(65)则acount++
 *		如果是L(76)并且lcount少于3次,则lcount++,否则计数器置0
 *		如果acount大于1或者lcount大于2 直接返回false
 *  返回true
**/
func checkRecord(s string) bool {
	acount, lcount := 0, 0
	length := len(s)
	for i := 0; i < length; i++ {
		if s[i] == 65 {
			acount++
		}
		if s[i] == 76 && lcount < 3 {
			lcount++
		} else {
			lcount = 0
		}
		if acount > 1 || lcount > 2 {
			return false
		}
	}
	return true
}
