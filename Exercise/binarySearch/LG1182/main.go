// 分段，求最大值最小
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var n, m int
var a []int

func main() {
	defer ot.Flush()
	n, m = rnI(), rnI()
	a = rsI(1, n+1)

	l, r := 0, int(1e9)
	for l < r {
		mid := (l + r) >> 1
		if valid(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	fmt.Println(l)
}

// 分成m组，每组之和<=mid ,是否可行?
// 如果可行，说明可以继续缩小，所以r=mid
// 注意：如果单个a[i]>mid.，说明此方案不行，要扩大
func valid(mid int) bool {
	g, rest := 1, mid

	for i := 1; i <= n; i++ {
		if a[i] > mid {
			return false
		}

		if rest >= a[i] {
			rest -= a[i]
		} else {
			g++
			rest = mid - a[i]
		}
	}

	return g <= m
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
