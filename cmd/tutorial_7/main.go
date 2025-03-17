package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	var i int32
	fmt.Printf("p location: %v; p value: %v", p, *p)
	fmt.Printf("\ni location: %v; i value: %v", &i, i)
	*p = 10
	fmt.Printf("\np location: %v; p value: %v", p, *p)

	p = &i
	fmt.Printf("\np location: %v; p value: %v", p, *p)
	fmt.Printf("\ni location: %v; i value: %v", &i, i)
	*p = 11
	fmt.Printf("\np location: %v; p value: %v", p, *p)
	fmt.Printf("\ni location: %v; i value: %v\n", &i, i)

	//Slices point to the same memory location
	var slice1 = []int32{1, 2, 3}
	var slice2 = slice1
	fmt.Printf("\nslice1: %v; slice2: %v", slice1, slice2)
	slice2[2] = 4 // Change 3 to 4 only in slice2
	fmt.Printf("\nslice1: %v; slice2: %v\n", slice1, slice2)

	var numberArr = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nMemory location 1: %p", &numberArr)
	var squaredArr [5]float64 = square(&numberArr)
	fmt.Printf("\nSquared array: %v", squaredArr)
	fmt.Printf("\nOriginal array: %v\n", numberArr)
}

func square(numberArr *[5]float64) [5]float64 {
	fmt.Printf("\nMemory location 2: %p\n", numberArr)
	for i := range numberArr {
		numberArr[i] = numberArr[i] * numberArr[i]
	}
	return *numberArr
}
