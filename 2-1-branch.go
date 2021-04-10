package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func normalIfGrammar() {

	const filename = "abc.txt"
	const nonfilename = "abcd.txt"
	contents, err := ioutil.ReadFile(filename)
	contents2, err2 := ioutil.ReadFile(nonfilename)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Printf("%s\n", contents2)
	}
}
func shorterIfGrammar() {

	const filename = "abc.txt"
	const nonfilename = "abcd.txt"

	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	if contents2, err2 := ioutil.ReadFile(nonfilename); err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Printf("%s\n", contents2)
	}
}

func normalSwitch(a, b int, operator string) int {
	var result int
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	default:
		panic("unsupported operator" + operator)
	}
	return result
}

func shorterSwitch(a, b int, operator string) int {
	var result int
	switch {
	case operator == "+":
		result = a + b
	case operator == "-":
		result = a - b
	default:
		panic("unsupported operator" + operator)
	}
	return result
}

func covertToBinForLoop(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFileForLoop(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func foreverLoop() {
	for {
		fmt.Println("abc")
	}
}
func main() {
	normalIfGrammar()
	shorterIfGrammar()
	a := normalSwitch(5, 2, "-")
	fmt.Println(a)

	c := shorterSwitch(10, 2, "+")
	fmt.Println(c)

	fmt.Println(covertToBinForLoop(5), covertToBinForLoop(55555))

	printFileForLoop("abc.txt")

}
