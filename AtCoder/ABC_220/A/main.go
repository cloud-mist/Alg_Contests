package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	a, b, c int
	flag    bool
)

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &a, &b, &c)
	tmp := c
	for c <= 1000 {
		if c >= a && c <= b {
			flag = true
			fmt.Println(c)
			return
		}
		c += tmp
	}
	if !flag {
		fmt.Println(-1)
	}
}
