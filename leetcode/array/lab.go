package main

import "fmt"

func main() {
	println(fmt.Sprintf("%T", [10]int{}))
	println(fmt.Sprintf("%T", make([]int, 10)))
	println(fmt.Sprintf("%T", map[int]int{}))
}
