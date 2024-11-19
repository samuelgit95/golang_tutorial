package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "résumé"
	var indexed = myString[0]
	fmt.Printf("%v, %T \n",indexed, indexed) // 114, uint8
	for i, v := range myString {
		fmt.Println(i,v)
		//Index 
		//_______
		// 0 114
		// 1 233
		// 3 115
		// 4 117
		// 5 109
		// 6 233
	}
	fmt.Printf("\nThe length of myString is %v \n", len(myString)) // The length of myString is 8
	var myRune = 'a'
	fmt.Printf("%v, %T \n",myRune, myRune) // 97, int32

	// String Building
	var strBuilder strings.Builder
	var strSlice = []string{"s", "a", "m"}
	for i := range strSlice{
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v\n", catStr) // sam
}