package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

//获得商和余数
func div(a, b int) (int, int) {
	return a / b, a % b
}

//获得商和余数
func div2(a, b int) (q, r int) {
	return a / b, a % b
}

func div3(a, b int) (q, r int) {
	q, r = a/b, a%b
	return
}

func getResult(a, b int) (sum, minus, division int) {
	sum, minus, division = a+b, a-b, a/b
	return
}

func useDiv() int {
	q, _ := div(3, 4)
	return q
}

func normalSwitch(a, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operator: %s", operator)
	}
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer() //获取指针
	opName := runtime.FuncForPC(p).Name()
	fmt.Println("Calling function %s with args "+" (%d, %d)", opName, a, b)
	return op(a, b)
}

func getPowResult(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func multiSum(nums ...int) int {
	result := 0
	for i := range nums {
		result += nums[i]
	}
	return result
}

func main() {
	fmt.Println(div(13, 4))
	fmt.Println(div3(13, 4))
	fmt.Println(normalSwitch(13, 4, "*"))
	fmt.Println(normalSwitch(13, 4, "/"))
	a, b, c := getResult(13, 4)
	fmt.Println("rsult - sum : %s, "+"minus: %s, "+"division: %s", a, b, c)

	fmt.Println(apply(getPowResult, 3, 4))

	fmt.Println(apply(func(i1, i2 int) int { return int(math.Pow(float64(i1), float64(i2))) }, 3, 4))

	fmt.Println(multiSum(13, 4, 5))
}
