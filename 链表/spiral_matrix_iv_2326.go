package linkedlist

// https://leetcode.cn/problems/spiral-matrix-iv/
// 螺旋矩阵 IV

func spiralMatrix0(m int, n int, head *ListNode) [][]int {
	top, left, button, right := 0, m-1, 0, n-1
	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		t := make([]int, n)
		for j := 0; j < n; j++ {
			t[j] = -1
		}
		ans[i] = t
	}
	for top <= button && left <= right {
		for i := left; i <= right && head != nil; i++ {
			ans[top][i] = head.Val
			head = head.Next
		}
		top++
		for i := top; i <= button && head != nil; i++ {
			ans[i][right] = head.Val
			head = head.Next
		}
		right--

		for j := right; j >= left && top <= button && head != nil; j-- {
			ans[button][j] = head.Val
			head = head.Next
		}
		button--

		for j := button; j >= top && left <= right && head != nil; j-- {
			ans[j][left] = head.Val
			head = head.Next
		}
		left++

	}
	return ans
}

var dirs = []struct{ x, y int }{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 右下左上

func spiralMatrix(n int, m int, head *ListNode) [][]int {
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
		for j := range ans[i] {
			ans[i][j] = -1
		}
	}
	for x, y, di := 0, 0, 0; head != nil; head = head.Next {
		ans[x][y] = head.Val
		d := dirs[di&3]
		if xx, yy := x+d.x, y+d.y; xx < 0 || xx >= n || yy < 0 || yy >= m || ans[xx][yy] != -1 {
			di++
			d = dirs[di&3]
		}
		x += d.x
		y += d.y
	}
	return ans
}
