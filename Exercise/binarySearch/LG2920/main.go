// 排序 贪心 二分
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var n int
var a []p

type p struct {
	t, s int
}

func main() {
	defer ot.Flush()
	n = rnI()
	for i := 0; i < n; i++ {
		a = append(a, p{rnI(), rnI()})
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i].s < a[j].s
	})

	l, r := -1, a[0].s-a[0].t
	for l < r {
		mid := (l + r + 1) >> 1
		if valid(mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}

	fmt.Println(l)
}

func valid(mid int) bool {
	k := mid + a[0].t
	for i := 1; i < n; i++ {
		k += a[i].t
		if k > a[i].s {
			return false
		}
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
