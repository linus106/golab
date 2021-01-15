package main

/**
https://leetcode-cn.com/problems/number-of-good-pairs/
*/

func numIdenticalPairs(nums []int) int {
	res := 0
	countMap := map[int]int{}
	for i := range nums {
		count, ok := countMap[nums[i]]
		if ok {
			res += count
			countMap[nums[i]] = count + 1
		} else {
			countMap[nums[i]] = 1
		}
	}
	return res
}

func main() {
	res := numIdenticalPairs([]int{1, 2, 3, 1, 1, 3})
	print(res)
}
