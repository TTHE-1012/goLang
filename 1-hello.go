package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"runtime"
)

var aa, bb, cc, dd = 1, 2, true, "123"
var (
	aaa, bbb = 1, 2
	ccc      = true
	ddd      = "123"
)

func varWithZeroValue() {
	var a, b int
	var c bool
	var d string
	fmt.Println(a, b, c, d)
}

func varWithInitialValue() {
	var a, b int = 1, 2
	var c bool = true
	var d string = "123"
	fmt.Println(a, b, c, d)
}

func varWithValueShorter() {
	var a, b, c, d = 1, 2, true, "123"
	e, f, g, h := 1, 2, true, "123"
	fmt.Println(a, b, c, d)
	fmt.Println(e, f, g, h)

}

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
}

func forceTypeTransfer() {
	a, b := 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	const a, b = 3, 4
	const filename = "abc.txt"
	var c int
	const d int = 5
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c, filename, d)
}

func enums() {
	const (
		cpp        = 0
		java       = 1
		python     = 3
		javascript = 4
	)
	const (
		cppx = iota
		_
		pythonx
		javascriptx
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(java, cpp, javascript, python)
	fmt.Println(javascriptx, pythonx)
	fmt.Println(b, kb, mb, gb, tb, pb)

}

func main() {
	fmt.Println(runtime.GOARCH)
	fmt.Println(aa, bb, cc, dd, aaa, bbb, ccc, ddd)
	varWithZeroValue()
	varWithInitialValue()
	varWithValueShorter()
	euler()
	forceTypeTransfer()
	enums()
}
