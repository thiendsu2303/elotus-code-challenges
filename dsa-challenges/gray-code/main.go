package main

import "fmt"

func grayCode(n int) []int {
	ans := make([]int, 0)
	ans = append(ans, 0)

	if n == 0 {
		return ans
	}

	ans = append(ans, 1)
	cur := 1
	for i := 2; i <= n; i++ {
		cur *= 2
		for j := len(ans) - 1; j >= 0; j-- {
			ans = append(ans, ans[j]+cur)
		}
	}
	return ans
}

func main() {
	fmt.Println(grayCode(2))
}
