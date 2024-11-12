package main

import "fmt"

func main() {
	intArr := [...]int32{1,2,3}
	fmt.Println(intArr) // [1 2 3]

	var intSlice []int32 = []int32{4,5,6}
	// print the length and the capacity with Printf
	fmt.Printf("Length: %v, Capacity: %v\n", len(intSlice), cap(intSlice)) // Length: 3, Capacity: 3
	fmt.Println(intSlice) // [4 5 6]
	intSlice = append(intSlice, 7)
	fmt.Printf("Length: %v, Capacity: %v\n", len(intSlice), cap(intSlice)) // Length: 4, Capacity: 6
	fmt.Println(intSlice) // [4 5 6 7]

	var intSlice2 []int32 = []int32{8,9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice) // [4 5 6 7 8 9]

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Printf("Length: %v, Capacity: %v\n", len(intSlice3), cap(intSlice3)) // Length: 3, Capacity: 8

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap) // map[]

	var myMap2 = map[string]uint8{"Adam":23, "Eve":24}
	fmt.Println(myMap2["Adam"]) // 23
	
	var age, ok = myMap2["Jason"]
	if ok {
		fmt.Printf("The age is %v \n", age)
	} else {
		fmt.Println("No age found")
	}

	for name:=range myMap2 {
		fmt.Printf("Name: %v \n", name)
	}

	for name, age:=range myMap2 {
		fmt.Printf("Name: %v, Age: %v \n", name, age)
	}

	for i, v := range intArr{
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

	for i:=0; i<10; i++ {
		fmt.Println(i)
	}
}