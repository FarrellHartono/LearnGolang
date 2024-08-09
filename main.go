package main

import "fmt"

func main() {

	// strings
	var nameOne string = "Farrell"
	var nameTwo = "luigi"
	var nameThree string
	nameThree = "mario"

	fmt.Println(nameOne, nameTwo, nameThree)
	nameOne = "Hartono"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "Ignatius"

	fmt.Println(nameFour)

	// ints
	var ageOne int = 21
	var ageTwo = 3
	ageThree := 4

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	var numOne int16 = 25
	var numTwo = -129
	var numThree uint8 = 255
	fmt.Println(numOne, numTwo, numThree)

	// Float
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 888729.21231
	// Bydefault ini float64
	scoreThree:= -1.5

	fmt.Println(scoreOne, scoreTwo, scoreThree)

}