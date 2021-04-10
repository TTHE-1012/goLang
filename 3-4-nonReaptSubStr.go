package main

import (
	"fmt"
	"unicode/utf8"
)

func lenOfNonReaptingSubStrByByte(s string) int {
	lastOccurred := make(map[byte]int) //记录每个字母最后出现的位置
	// fmt.Println(lastOccurred)
	start := 0
	maxLen := 0
	// for i, ch := range s{ //这里的ch类型为int32，学会rune之后使用
	for i, ch := range []byte(s) { //这里的ch为byte，是8位的
		// fmt.Println(i, ch)

		if lastI, ok := lastOccurred[ch]; ok {
			// if lastI, ok := lastOccurred[ch]; ok && lastI >= start {

			start = lastI + 1
		}
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
		lastOccurred[ch] = i
		// fmt.Println(lastOccurred)
	}
	return maxLen
}

func lenOfNonReaptingSubStrByRune(s string) int {
	lastOccurred := make(map[rune]int) //记录每个字母最后出现的位置
	start := 0
	maxLen := 0
	// for i, ch := range s {
	for i, ch := range []rune(s) {
		// if lastI, ok := lastOccurred[ch]; ok && start <= lastI {
		if lastI, ok := lastOccurred[ch]; ok {

			start = lastI + 1
		}
		if maxLen < i-start+1 {
			maxLen = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLen
}
func transerByte2Rune(s string) {
	for i, ch := range s {
		fmt.Printf("(%d %X)", i, ch)
		//(0 6280)(3 80FD)(6 FF1A)(9 4A)(10 61)(11 76)(12 61)(13 548C)(16 47)(17 6F)(18 4C)(19 61)(20 6E)(21 67)
	}
	fmt.Println()
	for i, ch := range []byte(s) {
		fmt.Printf("(%d %X)", i, ch)
		//(0 E6)(1 8A)(2 80)(3 E8)(4 83)(5 BD)(6 EF)(7 BC)(8 9A)(9 4A)(10 61)(11 76)(12 61)(13 E5)(14 92)(15 8C)(16 47)(17 6F)(18 4C)(19 61)(20 6E)(21 67)
	}
	fmt.Println()
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %X)", i, ch)
		//(0 6280)(1 80FD)(2 FF1A)(3 4A)(4 61)(5 76)(6 61)(7 548C)(8 47)(9 6F)(10 4C)(11 61)(12 6E)(13 67)
	}
	fmt.Println()
	fmt.Printf("RuneCountInString %s is %d", s, utf8.RuneCountInString(s))
	//Rune Count In String 技能：Java和GoLang is 14
	fmt.Println()
	bytes := []byte(s)
	// r, i := utf8.DecodeRuneInString(s)
	// fmt.Println("utf8.DecodeRuneInString(s) is ", r, i)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c", ch) //最后结果：技能：Java和GoLang
		// fmt.Print(ch)//最后结果：25216330216530674971189721644711117697110103
	}
}
func main() {
	fmt.Println(lenOfNonReaptingSubStrByByte("abcdefe"))
	fmt.Println(lenOfNonReaptingSubStrByByte("aaaa"))
	fmt.Println(lenOfNonReaptingSubStrByRune("abcdefe"))
	fmt.Println(lenOfNonReaptingSubStrByRune("aaaa"))
	fmt.Println(lenOfNonReaptingSubStrByRune("我喜欢吃苹果也喜欢吃梨"))
	transerByte2Rune("技能：Java和GoLang")
}
