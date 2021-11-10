// max median

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	a    []int
	n, k int
)

func main() {
	defer ot.Flush()
	n, k = rnI(), rnI()
	a = rsI(0, n)

	// 排序
	sort.Ints(a)

	// 二分
	l, r := int64(0), int64(2e9)
	for l < r {
		mid := int64((l + r + 1)) >> 1
		if calc(mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}

	fmt.Print(l)
}

// 将中位数后面的数都加到和中位数至少持平,如果没超过k次操作,说明可以，否则不行
func calc(x int64) bool {
	var tmp int64
	for i := n / 2; i < n; i++ {
		if int64(a[i]) < x {
			tmp += x - int64(a[i])
		}
	}

	return tmp <= int64(k)
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
	in = bufio.NewScanner(os.Stdin)
)

func init()        { in.Split(bufio.ScanWords); in.Buffer(make([]byte, 4096), int(1e9)) }
func rnS() string  { in.Scan(); return in.Text() }
func rnI() int     { i, _ := strconv.Atoi(rnS()); return i }
func rnF() float64 { f, _ := strconv.ParseFloat(rnS(), 64); return f }

func rsI(l, r int) []int {
	t := make([]int, r)
	for i := l; i < r; i++ {
		t[i] = rnI()
	}
	return t
}

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
