package main

import "fmt"

type treeNode struct {
	//go语言没有构造函数的说法
	//如果想控制其构造，可以通过工厂函数，下面有函数例子
	value       int
	left, right *treeNode
}

func createNode(value int) *treeNode {
	//注意：这里返回的是局部变量的地址 -- 与java的对象类似
	//为什么不能直接return treeNode{value: value}： 因为go是值传递
	return &treeNode{value: value}

}
func main() {
	var root treeNode
	fmt.Println(root) //{0 <nil> <nil>}

	root2 := treeNode{value: 5}
	root2.left = &treeNode{3, nil, nil}
	root2.right = &treeNode{}
	//不像C++， go不管地址/指针还是结构本身/对象，一律使用' . '来访问成员
	root2.right.left = new(treeNode)
	root2.right.right = createNode(2)
	fmt.Println(root2) //&{2 <nil> <nil>}  这里是一个地址

	root2.right.right = createNode(2)
	fmt.Println(root2.right.right) //{5 0xc0000040a8 0xc0000040c0}

	nodes := []treeNode{
		{},
		{value: 5},
		{0, nil, &root},
	}
	fmt.Println(nodes) //[{0 <nil> <nil>} {5 <nil> <nil>} {0 <nil> 0xc000004078}]

}
