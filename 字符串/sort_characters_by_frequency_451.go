package string

import (
	"bytes"
	"sort"
)

// 根据字符出现频率排序
// https://leetcode.cn/problems/sort-characters-by-frequency/description/
// 哈希 + 排序
func frequencySort0(s string) string {
	cnt := map[byte]int{}
	for i := range s {
		cnt[s[i]]++
	}

	type pair struct {
		ch  byte
		cnt int
	}
	pairs := make([]pair, 0, len(cnt))
	for k, v := range cnt {
		pairs = append(pairs, pair{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].cnt > pairs[j].cnt })

	ans := make([]byte, 0, len(s))
	for _, p := range pairs {
		ans = append(ans, bytes.Repeat([]byte{p.ch}, p.cnt)...)
	}
	return string(ans)
}

// 桶排序
func frequencySort(s string) string {
	cnt := map[byte]int{}
	maxFreq := 0
	for i := range s {
		cnt[s[i]]++
		maxFreq = max(maxFreq, cnt[s[i]])
	}

	buckets := make([][]byte, maxFreq+1)
	for ch, c := range cnt {
		buckets[c] = append(buckets[c], ch)
	}

	ans := make([]byte, 0, len(s))
	for i := maxFreq; i > 0; i-- {
		for _, ch := range buckets[i] {
			ans = append(ans, bytes.Repeat([]byte{ch}, i)...)
		}
	}
	return string(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
