package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100010

var (
	n, x int
	a    [N]int
)

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	fmt.Fscan(in, &x)

	sum_a := 0
	for _, v := range a {
		sum_a += v
	}

	res := x / sum_a
	tmp := res * sum_a
	res *= n

	for _, v := range a {
		if tmp <= x {
			tmp += v
			res++
		}
	}

	fmt.Println(res)
}
