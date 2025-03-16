package main

import (
	"fmt"
	"strings"
)

func main() {
	const myString = "résumé"
	var indexed byte = myString[1]
	fmt.Printf("%v, %T\n", indexed, indexed)
	fmt.Println("String iteration")
	for i, v := range myString {
		fmt.Printf("Index: %v, Value: %c\n", i, v)
	}

	var myRunes []rune = []rune(myString)
	var indexedRune rune = myRunes[1]
	fmt.Printf("\n%v, %T", indexedRune, indexedRune)
	fmt.Println("\nRunes iteration")
	for i, v := range myRunes {
		fmt.Printf("Index: %v, Value: %c\n", i, v)
	}

	var strSlice = []string{"h", "e", "l", "l", "o"}
	//Inefficient way
	// var catStr string
	// for i := range strSlice {
	// 	catStr += strSlice[i]
	// }
	// fmt.Println(catStr)

	//Efficient way
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Println(catStr)
}
