package main

import (
	"fmt"
)

func main() {
	// Slices
	fmt.Println("======================")
	fmt.Println("SLICES")
	fmt.Println("======================")
	var arrInt []int = []int{1, 2, 3, 4, 5}
	var sliceInt = arrInt[3:]

	fmt.Printf("Init\n")
	fmt.Printf("ArrInt: %+v, SliceInt: %+v\n\n", arrInt, sliceInt)

	sliceInt[0] = 10
	fmt.Printf("After\n")
	fmt.Printf("ArrInt: %+v, SliceInt: %+v\n", arrInt, sliceInt)

	// MAP
	fmt.Println("======================")
	fmt.Println("MAP")
	fmt.Println("======================")
	var emptyMap = make(map[string]interface{})
	fmt.Println("Init")
	fmt.Printf("emptyMap : %+v\n", emptyMap)

	MapFunc(emptyMap)
	fmt.Println("After")
	fmt.Printf("emptyMap : %+v\n", emptyMap)
}

func MapFunc(val map[string]interface{}) {
	val["this is a new value"] = 100
}
