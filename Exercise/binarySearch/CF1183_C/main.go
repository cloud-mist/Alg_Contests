package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var k, n, a, b int

func main() {
	defer ot.Flush()

	q := rnI()
	for ; q != 0; q-- {
		k, n, a, b = rnI(), rnI(), rnI(), rnI()

		if (k/b == n && k%b == 0) || k/b < n {
			fmt.Fprintln(ot, -1)
		} else {
			var l, r int64
			l, r = 0, int64(k/b)
			for l < r {
				var mid int64
				mid = (l + r + 1) >> 1
				if calc(mid) {
					l = mid
				} else {
					r = mid - 1
				}
			}
			fmt.Fprintln(ot, min(int64(n), l))
		}
	}
}

func calc(x int64) bool {
	tmp := int64(a)*x + int64(b)*(int64(n)-x)
	if tmp >= int64(k) {
		return false
	}

	return true
}

/* ======================================================================== */
//
//
//
//							 _____   _   _   ____
//							| ____| | \ | | |  _ \
//							|  _|   |  \| | | | | |
//							| |___  | |\  | | |_| |
//							|_____| |_| \_| |____/
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
func min(x, y int64) int64 {
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
