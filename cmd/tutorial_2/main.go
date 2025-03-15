package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const a uint64 = 32767
	fmt.Println(a)

	const b float32 = 12345678.9
	fmt.Println(b)

	const result float32 = b * float32(a)
	fmt.Println(result)

	const gamma string = "γ"
	fmt.Printf("Gamma: %s, Length: %v\n", gamma, len(gamma))
	fmt.Printf("Gamma: %s, RuneCount: %v\n", gamma, utf8.RuneCountInString(gamma))

	const charInGo rune = 'a'
	fmt.Printf("Char: %c, Unicode: %U\n", charInGo, charInGo)

	var intWIthoutInitValue int
	fmt.Println("Initial value of int is: ", intWIthoutInitValue)

	shortcutVar := "Short init"
	fmt.Println("Shortcut var: ", shortcutVar)
}
