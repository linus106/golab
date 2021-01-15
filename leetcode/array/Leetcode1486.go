package main

/**
https://leetcode-cn.com/problems/xor-operation-in-an-array/
*/

func xorOperation(n int, start int) int {
	res := start
	for i := 1; i < n; i++ {
		res ^= start + 2 * i
	}
	return res
}

func main() {
	res := xorOperation(5, 0)
	println(res)
}
