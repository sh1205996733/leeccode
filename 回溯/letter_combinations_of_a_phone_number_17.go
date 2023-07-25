package backtrack

// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
// 电话号码的字母组合

// 方法一：回溯法
// 23 [a b c、 d e f] ["ad","ae","af","bd","be","bf","cd","ce","cf"]
var m = map[uint8][]string{
	50: {"a", "b", "c"},      // 2
	51: {"d", "e", "f"},      // 3
	52: {"g", "h", "i"},      // 4
	53: {"j", "k", "l"},      // 5
	54: {"m", "n", "o"},      // 6
	55: {"p", "q", "r", "s"}, // 7
	56: {"t", "u", "v"},      // 8
	57: {"w", "x", "y", "z"}, // 9
}

func letterCombinations0(digits string) []string {
	if digits == "" {
		return nil
	}
	var ans []string
	ans = dfsLetter(ans, "", digits, 0)
	return ans
}
func dfsLetter(ans []string, combination, digits string, i int) []string {
	if i == len(digits) {
		ans = append(ans, combination)
		return ans
	}
	chars := m[digits[i]]
	for _, c := range chars {
		ans = dfsLetter(ans, combination+c, digits, i+1)
	}
	return ans
}

// 方法一：队列 先将输入的 digits 中第一个数字对应的每一个字母入队，然后将出队的元素与第二个数字对应的每一个字母组合后入队...直到遍历到 digits 的结尾。最后队列中的元素就是所求结果
// 23 [a b c、 d e f] ["ad","ae","af","bd","be","bf","cd","ce","cf"]
func letterCombinations1(digits string) []string {
	if digits == "" {
		return nil
	}
	var queue []string
	chars := m[digits[0]]
	// 先将输入的 digits 中第一个数字对应的每一个字母入队
	for _, c := range chars {
		queue = append(queue, c)
	}
	for i := 1; i < len(digits); i++ {
		chars := m[digits[i]]
		l := len(queue)
		for _, q := range queue {
			for _, c := range chars {
				queue = append(queue, q+c)
			}
		}
		// 原队列出队
		queue = queue[l:]
	}
	return queue
}

func letterCombinations(digits string) []string {
	var queue []string
	for i := 0; i < len(digits); i++ {
		chars := m[digits[i]]
		if len(queue) == 0 {
			queue = append(queue, chars...)
			continue
		}
		l := len(queue)
		for _, q := range queue {
			for _, c := range chars {
				queue = append(queue, q+c)
			}
		}
		queue = queue[l:]
	}
	return queue
}

// 相似题目
// 括号生成 https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
