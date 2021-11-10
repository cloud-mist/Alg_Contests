package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

const BUFSIZE = int(1e6)
const N = 200010

func main() {
	defer ot.Flush()

	// Input
	n := rnI()
	l := rsI(n)
	m := rnI()
	vc := rsI(m)
	cc := rsI(m)

	sign := make(map[int]int)
	for i := range l {
		sign[l[i]]++ // count
	}

	// -----------------------------------------------

	// 找出 懂语音最多的
	maxVc := 0
	for i := 0; i < m; i++ {
		maxVc = max(maxVc, sign[vc[i]])
	}

	// 如果和 语音最多相等，就看字幕人数
	maxCc, res := -1, 0
	for i := 0; i < m; i++ {
		x, y := sign[vc[i]], sign[cc[i]]
		if x == maxVc {
			if y > maxCc {
				maxCc = y
				res = i + 1
			}
		}
	}
	fmt.Print(res)

}

/* ======================================================================== */
//
//
//							_____   _   _   ____
//							| ____| | \ | | |  _ \
//							|  _|   |  \| | | | | |
//							| |___  | |\  | | |_| |
//							|_____| |_| \_| |____/
//
//
//
/* ============================PART1: I/O ================================== */

/* ==== G0 ===== */
var (
	ot = bufio.NewWriterSize(os.Stdout, BUFSIZE)
	in = newScanner(os.Stdin)
)

type scanner struct{ sc *bufio.Scanner }

func newScanner(input io.Reader) *scanner {
	sc := bufio.NewScanner(input)
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 1024), int(1e9))
	return &scanner{sc}
}

/* ==== G1 ===== */
func rnS() string  { in.sc.Scan(); return in.sc.Text() }
func rnI() int     { i, _ := strconv.Atoi(rnS()); return i }
func rnF() float64 { f, _ := strconv.ParseFloat(rnS(), 64); return f }

/* ==== G2 ===== */
func rsI(n int) []int {
	t := make([]int, n)
	for i := 0; i < n; i++ {
		t[i] = rnI()
	}
	return t
}
func rsF(n int) []float64 {
	t := make([]float64, n)
	for i := 0; i < n; i++ {
		t[i] = rnF()
	}
	return t
}
func rsS(n int) []string {
	t := make([]string, n)
	for i := 0; i < n; i++ {
		t[i] = rnS()
	}
	return t
}

/*				   +++++++++++++ PART1 END ++++++++++++					*/

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
