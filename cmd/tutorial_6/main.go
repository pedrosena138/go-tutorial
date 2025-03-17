package main

import "fmt"

// type owner struct {
// 	name string
// }

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

type eletricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e eletricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) bool {
	return miles <= e.milesLeft()
}

func main() {
	engine1 := gasEngine{mpg: 10, gallons: 100}
	// var engine gasEngine = gasEngine{10, 100}
	fmt.Println(engine1.mpg, engine1.gallons)
	engine1.mpg = 15
	fmt.Println(engine1.mpg, engine1.gallons)

	// var engine2 = struct {
	// 	mpg       uint8
	// 	gallons   uint8
	// 	ownerInfo owner
	// }{25, 15, owner{"Adam"}}
	// fmt.Println(engine2.mpg, engine2.mpg, engine2.ownerInfo.name)

	// var engine3 = struct {
	// 	mpg     uint8
	// 	gallons uint8
	// 	owner
	// }{25, 15, owner{"Adam"}}
	// fmt.Println(engine3.mpg, engine3.mpg, engine3.name)

	engine2 := eletricEngine{10, 100}
	var miles uint8 = 200

	if canMakeIt(engine1, miles) {
		fmt.Println("You can make it!!")
	} else {
		fmt.Println("Need fuel!!")
	}

	if canMakeIt(engine2, miles) {
		fmt.Println("You can make it!!")
	} else {
		fmt.Println("Need fuel!!")
	}
}
