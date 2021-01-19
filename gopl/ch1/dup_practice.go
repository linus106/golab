package main

import (
	"bufio"
	"fmt"
	"os"
)

//修改dup2，出现重复的行时打印文件名称。
func main() {
	counts := make(map[string][]string)
	files := os.Args[1:]

	if len(files) == 0 {
		mapLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			mapLines(f, counts)
			f.Close()
		}
	}

	for line, filenames := range counts {
		if len(filenames) > 1 {
			fmt.Println(line, filenames)
		}
	}
}

func mapLines(f *os.File, counts map[string][]string)  {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()] = append(counts[input.Text()], f.Name())
	}
}
