package main

import "fmt"

func printArr(arr [5]int) { //带参数的[5]int 才是数组类型，而不带数量的[]int时切片，后面提到
	for i := range arr {
		fmt.Println(arr[i])
	}
	arr[0] = 100
}

func printArrUsingPointer(arr *[5]int) { //带参数的[5]int 才是数组类型，而不带数量的[]int时切片，后面提到
	for i := range arr {
		fmt.Println(arr[i])
	}
	arr[0] = 100 // 这里不像其他变量，需要加*然后重新赋值
}
func main() {
	var arr1 [5]int
	arr2 := []int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10} //编译器自动计算个数
	fmt.Println(arr1, arr2, arr3)

	var grid [4][5]int //4行5列
	fmt.Println(grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for i := range arr3 {
		fmt.Println(i)
	}

	for i, v := range arr3 {
		fmt.Println(i, v)
	}
	for _, v := range arr3 {
		fmt.Println(v)
	}

	printArr(arr3)
	fmt.Println(arr3)
	printArrUsingPointer(&arr3)

	sliceArray := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sliceArray[2:6]
	fmt.Println(s)
}
