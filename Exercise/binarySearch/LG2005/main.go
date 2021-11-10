// 高精二分
package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

var (
	a, b big.Int
)

func main() {
	defer ot.Flush()
	fmt.Fscan(in, &a, &b)

	l, r := 0, 6227020810
	for l < r {
		mid := (l + r + 1) >> 1
		if calc(mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}

	fmt.Print(l)
}

func calc(mid int) bool {
	t := big.NewInt(int64(mid))
	t.Mul(t, &b)
	if t.Cmp(&a) > 0 {
		return false
	}
	return true
}

/* ======================================================================== */
//
//
//
//				 _____   _   _   ____
//				| ____| | \ | | |  _ \
//				|  _|   |  \| | | | | |
//				| |___  | |\  | | |_| |
//				|_____| |_| \_| |____/
//
//
//
/* ============================PART 1: I/O ================================== */

var (
	ot = bufio.NewWriterSize(os.Stdout, int(1e6))
	in = bufio.NewReader(os.Stdin)
)

/* ===========================PART 2: Math Func ============================  */
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func memset(a []int, v int) {
	if len(a) == 0 {
		return
	}
	a[0] = v
	for bp := 1; bp < len(a); bp *= 2 {
		copy(a[bp:], a[:bp])
	}
}
