// 跳石头
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var l, n, m int
var a []int

func main() {
	defer ot.Flush()
	l, n, m = rnI(), rnI(), rnI()
	a = rsI(1, n+1)
	a = append(a, l)
	sort.Ints(a)

	l, r := 0, int(1e9)
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
	res := 0
	i, j := 1, 0
	for i <= n+1 {
		if a[i]-a[j] < mid {
			res++
		} else {
			j = i
		}
		i++
	}
	return res <= m
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
