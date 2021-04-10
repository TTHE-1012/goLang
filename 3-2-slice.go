package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 1000
}

func modifyArr(arr []int) {
	arr[0] = 100
}

func main() {
	sliceArray := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sliceArray[2:6]
	fmt.Println(s) // [2 3 4 5]
	s = sliceArray[:6]
	fmt.Println(s) //[0 1 2 3 4 5]
	s = sliceArray[2:]
	fmt.Println(s) //[2 3 4 5 6 7 8 9]
	s = sliceArray[:]
	fmt.Println(s) //[0 1 2 3 4 5 6 7 8 9]

	s1 := sliceArray[2:6]
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(sliceArray)

	reslice := sliceArray[:]
	fmt.Println(reslice) //[0 1 1000 3 4 5 6 7 8 9]
	reslice = reslice[2:8]
	fmt.Println(reslice) //[1000 3 4 5 6 7]
	reslice = reslice[:3]
	fmt.Println(reslice) //[1000 3 4]

	modifyArr(sliceArray[:])

	sliceArray = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	extSlice1 := sliceArray[2:6] //[2 3 4 5]
	extSlice2 := extSlice1[3:5]  //[5 6] 没有报错，这里就是slice的扩展，底层其实知道后面还有两个数
	// fmt.Println(extSlice1[4])    //这里会报错，下标越界
	fmt.Println(extSlice2)
	fmt.Println(extSlice1)
	fmt.Printf("extSlice1 = %v, len(extSlice1) = %d, cap(extSlice1) = %d \n", extSlice1, len(extSlice1), cap(extSlice1))

	extSlice2 = sliceArray[5:9]
	extSlice3 := append(extSlice2, 10)
	fmt.Println(extSlice2) // [5 6 7 8]
	fmt.Println(extSlice3) // [5 6 7 8 10]  10成功添加
	extSlice4 := append(extSlice2, 11)
	extSlice5 := append(extSlice3, 12)
	fmt.Println(extSlice4) // [5 6 7 8 11]  11 成功添加，但是覆盖了10，而且最终的长度应该超过了cap
	fmt.Println(extSlice5) // [5 6 7 8 11 12] 12 成功添加，而且最终的长度应该超过了cap
	//这里的extSlice4和extSlice5的长度已经超出了slice的cap，其实是使用了新的array，而不再view sliceArray
	// extSlice4 and extSlice5 no longer view sliceArray
	fmt.Println(sliceArray) // [0 1 2 3 4 5 6 7 8 11]  最终11被成功添加， 但是覆盖了原来的9，原数组长度没变

	var newSlice []int //Zero value for slice is nil
	fmt.Printf("newSlice = %v, len(newSlice) = %d, cap(newSlice) = %d \n", newSlice, len(newSlice), cap(newSlice))
	//newSlice = [], len(newSlice) = 0, cap(newSlice) = 0
	newSlice = append(newSlice, 1)
	fmt.Printf("newSlice = %v, len(newSlice) = %d, cap(newSlice) = %d \n", newSlice, len(newSlice), cap(newSlice))
	//newSlice = [1], len(newSlice) = 1, cap(newSlice) = 1
	newSlice = append(newSlice, 2)
	fmt.Printf("newSlice = %v, len(newSlice) = %d, cap(newSlice) = %d \n", newSlice, len(newSlice), cap(newSlice))
	//newSlice = [1 2], len(newSlice) = 2, cap(newSlice) = 2

	newSlice2 := []int{2, 4, 6, 8} //数组的创建是 [...]int{2, 4, 6, 8}
	fmt.Println(newSlice2)

	newSlice3 := make([]int, 8)     //8是len
	newSlice4 := make([]int, 8, 32) //8是len, 32是cap
	fmt.Println(newSlice3, newSlice4)

	copy(newSlice4, newSlice2)
	fmt.Println(newSlice2, newSlice4)

	newSlice4 = append(newSlice4[:3], newSlice4[4:]...) //这里的newSlice4[4:]...是转换成可变参数elems ...Type
	fmt.Println(newSlice4)                              //[2 4 6 0 0 0 0] 将下表为3的元素8删除了

	pop := newSlice4[:1]
	newSlice4 = newSlice4[1:]
	fmt.Println(pop, newSlice4, cap(newSlice4)) // [2] [4 6 0 0 0 0] 31

	newSlice4 = newSlice4[5:]
	fmt.Println(newSlice4, cap(newSlice4)) // [2] [4 6 0 0 0 0] 31
}
