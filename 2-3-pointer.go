package main

import (
	"fmt"	"strconv"
)


func swap1(a, b int) (int, int) {
	return b, a
}

func swap2(a, b *int) {
	a, b = b, a
}

func swap3(a, b *int) {
	*a, *b = *b, *a
}

func main() {

	var a string = "12"
	var pa *string = &a
	var b = a
	fmt.Println(a)
	*pa = "123"
	fmt.Println(a)
	b = "1234"
	fmt.Println(b)
	fmt.Println(a)

	var d int = 3
	var pd *int = &d
	*pd = 5
	fmt.Println(d)

	l1, l2 := 3, 4
	swap2(&l1, &l2)
	fmt.Println(l1, l2)
	swap3(&l1, &l2)
	fmt.Println(l1, l2)
	l1, l2 = swap1(l1, l2)
	fmt.Println(l1, l2)

	p1, p2 := 3, 4
	pp1, pp2 := &p1, &p2
	fmt.Println(pp1, pp2)
	pp2 = pp1
	fmt.Println(pp1, pp2)

}
