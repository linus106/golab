package main

/**
https://leetcode-cn.com/problems/kids-with-the-greatest-number-of-candies/
 */

func main()  {
	res := kidsWithCandies([]int{2, 3, 5, 1, 3}, 3)
	for i := range res {
		print(res[i])
	}
}

func kidsWithCandies(candies []int, extraCandies int) []bool {

	var max = 0
	for i := range candies {
		max = maxInt(max, candies[i])
	}

	res := make([]bool, len(candies))
	for i, candy := range candies {
		res[i] = (candy + extraCandies) >= max
	}
	return res
}

func maxInt(a, b int) int {
	if (a < b) {
		return b;
	}
	return a;
}