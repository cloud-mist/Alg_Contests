// 枚举时间，先所有减去自然干的，然后累加用机器的，如果累加之和<=mid,说明可行
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const N = 500010

var n, x, y int

func main() {
	defer ot.Flush()
	n, x, y = rnI(), rnI(), rnI()
	var a [N]int

	for i := 0; i < n; i++ {
		a[i] = rnI()
	}

	l, r := 1, 500010
	for l < r {
		mid := (l + r) >> 1
		if valid(a, mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	fmt.Println(l)
}

func valid(a [N]int, mid int) bool {
	k := 0
	for i := 0; i < n; i++ {
		a[i] -= x * mid
		if a[i] > 0 {
			k += int(math.Ceil(float64(a[i]) / float64(y)))
		}
	}

	return k <= mid
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
