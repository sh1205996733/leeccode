package dp

import (
	"fmt"
	"github.com/bmizerany/assert"
	"testing"
)

func TestNameClimbStairs(t *testing.T) {
	fmt.Println(climbStairs(10))
}

func TestMaxProfit(t *testing.T) {
	nums := []int{7, 18, 1, 5, 3, 6, 4}
	assert.Equal(t, maxProfit(nums), 5)
}
func TestMaxProfit2(t *testing.T) {
	nums := []int{7, 18, 1, 5, 6, 6, 4}
	assert.Equal(t, maxProfit2(nums), 18)
}
