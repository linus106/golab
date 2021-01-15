package main

/**
https://leetcode-cn.com/problems/shuffle-the-array/
*/

func shuffle(nums []int, n int) []int {
	res := make([]int, 2*n)
	for i := 0; i < n; i++ {
		res[2*i] = nums[i]
		res[2*i+1] = nums[i+n]
	}
	return res
}

func main() {
	res := shuffle([]int{2, 5, 1, 3, 4, 7}, 3)
	for i := range res {
		println(res[i])
	}
}
