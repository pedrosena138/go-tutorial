package main

import (
	"errors"
	"fmt"
)

func getAge(name string) (uint8, error) {
	var err error
	var myMap = map[string]uint8{"Jonh": 10, "Doe": 20}

	var age, ok = myMap[name]
	if !ok {
		err = errors.New("error: key " + name + " does not exist")
	}

	return age, err
}

func main() {

	var intArr [3]int32
	intArr[1] = 123
	fmt.Println("First element:", intArr[0])
	fmt.Println("Slice second and third element", intArr[1:3])

	// Memory location -> &
	fmt.Printf("Memory location of first element: %p\n", &intArr[0])
	fmt.Printf("Memory location of second element: %p\n", &intArr[1])
	fmt.Printf("Memory location of third element: %p\n", &intArr[2])

	// Init
	var floatArr [4]float32 = [4]float32{1.1, 2.2, 3.3, 4.4}
	fmt.Println("Float Arr:", floatArr)

	strArr := [...]string{"first", "second", "third"}
	fmt.Println("String Arr:", strArr)

	// Slices
	var intSlice []int32 = []int32{1, 2, 3, 4, 5}
	fmt.Printf("Slice: %v; Lenght: %v; Capacity: %v\n", intSlice, len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 6)
	fmt.Printf("Slice: %v; Lenght: %v; Capacity: %v\n", intSlice, len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{2, 3}
	intSlice = append(intSlice, intSlice2...) //Spread operator -> ...
	fmt.Printf("Slice: %v; Lenght: %v; Capacity: %v\n", intSlice, len(intSlice), cap(intSlice))

	// make function
	var intSlice3 []int32 = make([]int32, 5, 10)
	fmt.Printf("Slice3: %v; Lenght: %v; Capacity: %v\n", intSlice3, len(intSlice3), cap(intSlice3))

	var myMap map[string]uint8 = make(map[string]uint8)
	myMap["Bob"] = 34
	myMap["Alice"] = 23
	// fmt.Printf("This key Adam does not exist. The value is: %v\n", myMap["Adam"]) // value = 0
	fmt.Println("Map:", myMap)
	delete(myMap, "Bob")
	fmt.Println("Map:", myMap)

	var name string = "Jonh"
	var age, err = getAge(name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Age of %v is: %v\n", name, age)

	//Loops
	for key, value := range myMap {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
	}

	for i, value := range intArr {
		fmt.Printf("Index: %v, Value: %v\n", i, value)
	}
	// 'While'
	// var i uint8
	// for i < 10 {
	// 	fmt.Println("While: ", i)
	// 	i++
	// }

	// for {
	// 	if i >= 10 {
	// 		break
	// 	}
	// 	fmt.Println("While: ", i)
	// 	i++
	// }
	for i := 0; i < 5; i++ {
		fmt.Println("Index: ", i)
	}
}
