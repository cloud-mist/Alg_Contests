package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	n, k int
	a    []float64
)

func main() {
	defer ot.Flush()
	n, k = rnI(), rnI()
	a = rsF(0, n)

	l, r := 0.0, 100000.0
	for l+1e-6 < r {
		mid := (l + r) / 2
		if calc(mid) {
			l = mid
		} else {
			r = mid
		}
	}
	fmt.Printf("%.4f\n", l)
}

// 枚举绳子长度
func calc(x float64) bool {
	res := 0
	for i := 0; i < n; i++ {
		res += int(a[i] / x)
	}

	return res >= k
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
func rsF(l, r int) []float64 {
	t := make([]float64, r)
	for i := l; i < r; i++ {
		t[i] = rnF()
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
