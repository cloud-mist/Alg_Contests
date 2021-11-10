package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	k    int
	a, b string
)

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &k)
	fmt.Fscan(in, &a, &b)
	fmt.Println(change(a) * change(b))

}

func change(x string) int {
	// method 1:
	//	res, _ := strconv.ParseInt(x, k, 64)

	// method 2:
	res := 0
	for _, v := range x {
		res = res*k + int(v-'0')
	}

	return res
}
