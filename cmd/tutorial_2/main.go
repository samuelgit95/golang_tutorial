package main

import (
	"errors"
	"fmt"
)

func main(){
	var printValue string = "Hello World!"
	printMe(printValue) // "Hello World!"

	var numerator int = 11
	var denominator int = 2
	var result, remainder, err  = intDivision(numerator, denominator)
	
	switch {
		case err != nil:
			fmt.Println(err.Error())
		case remainder == 0:
			fmt.Printf("The result of the integer division is %v \n", result)
		default:
			fmt.Printf("The result of the integer division in %v with remainder %v \n", result, remainder) // "The result of the integer division in 5 with remainder 1"
	}

	switch remainder {
	case 0:
		fmt.Printf("The division was exact \n")
	case 1,2:
		fmt.Printf("The division was close \n")
	default:
		fmt.Printf("The division was not close \n")
	}
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Printf("The result of the integer division is %v", result)
	// } else {
	// 	fmt.Printf("The result of the integer division in %v with remainder %v \n", result, remainder) // "The result of the integer division in 5 with remainder 1"
	// }
}

func printMe(printValue string){
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0{
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator/denominator
	var remainder int = numerator%denominator
	return result, remainder, err
}