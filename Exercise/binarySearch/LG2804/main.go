// -m , 求前缀和，求顺序对数量
// 求那个，那个条件为不包含等
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 92084931

var a []int

func main() {
	defer ot.Flush()
	n, m := rnI(), rnI()

	a = make([]int, n+1)
	for i := 1; i <= n; i++ {
		t := rnI() - m
		a[i] = t + a[i-1]
	}

	fmt.Println(mergeSort(0, n))
}

func mergeSort(l, r int) int {
	if l >= r {
		return 0
	}

	mid := (l + r) >> 1
	res := mergeSort(l, mid) + mergeSort(mid+1, r)

	i, j := l, mid+1
	tmp := make([]int, 0)
	for i <= mid && j <= r {
		if a[i] < a[j] {
			// 顺序对, 那么j后面的所有数都对a[i]大
			tmp = append(tmp, a[i])
			res += r - j + 1
			res %= MOD
			i++
		} else {
			tmp = append(tmp, a[j])
			j++
		}
	}

	for ; i <= mid; i++ {
		tmp = append(tmp, a[i])
	}
	for ; j <= r; j++ {
		tmp = append(tmp, a[j])
	}

	for i, j := l, 0; i <= r; i, j = i+1, j+1 {
		a[i] = tmp[j]
	}

	return res
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
