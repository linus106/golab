package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	p3()
}

//修改echo程序，使其能够打印os.Args[0]，即被执行命令本身的名字。
func p1() {
	fmt.Println(strings.Join(os.Args, " "))
}

//修改echo程序，使其打印每个参数的索引和值，每个一行。
func p2() {
	for i, s := range os.Args[1:] {
		fmt.Println(i, s)
	}
}

//： 做实验测量潜在低效的版本和使用了strings.Join的版本的运行时间差异。
func p3() {
	// slow
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(time.Since(start).Nanoseconds())

	//faster
	start2 := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(time.Since(start2).Nanoseconds())
}
