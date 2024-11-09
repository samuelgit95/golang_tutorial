package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// fmt.Println("Hello World!")
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum32 float32 = 12345678.9
	var intnum32 int32 = 2
	var result float32 = floatNum32 + float32(intnum32)
	fmt.Println(result) // 1.2345681e+07

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1/intNum2) // 1
	fmt.Println(intNum1%intNum2) // 1

	var myString string = "Hello" + " " + "World"
	fmt.Println(myString) // Hello World

	fmt.Println("len of 'test':", len("test")) // return the number of bytes
	// So if a char isn't in 256 ASCII characters it will be store with more that 1 byte
	fmt.Println("len of '😎':", len("😎")) // len of '😎': 4

	// if you are expecting fancy strings in your program you should use rune
	fmt.Println("len of '😎':", utf8.RuneCountInString("😎")) // len of '😎': 1

	var myBoolean bool = false
	fmt.Println("boolean:", myBoolean) // boolean: false

	var intNum3 rune
	fmt.Println("rune:", intNum3) // rune: 0

	myvar := "text"
	fmt.Println(myvar)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	const myConst string = "constant value"
	fmt.Println(myConst)
}