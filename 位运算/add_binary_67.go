package bit

// 二进制求和
// https://leetcode.cn/problems/add-binary/description/

func addBinary(a string, b string) string {
	result := ""
	if len(a) < len(b) {
		a, b = b, a
	}
	start, end := len(a)-1, len(a)-len(b)
	var carry = "0" //进位
	//var carry byte //进位
	for i := start; i >= 0; i-- {
		//carry = a[i] - '0' + carry
		//if i >= end {
		//	carry += b[i-end] - '0'
		//}
		//result = strconv.Itoa(int(carry%2)) + result
		//carry = carry / 2
		d := a[i] - '0' + carry[0] - '0'
		if i >= end {
			d += b[i-end] - '0'
		}
		if d > 1 {
			if d == 2 {
				result = "0" + result
			} else {
				result = "1" + result
			}
			carry = "1"
		} else {
			result = string(d+'0') + result
			carry = "0"
		}
	}
	if carry == "1" {
		result = carry + result
	}
	return result
}
