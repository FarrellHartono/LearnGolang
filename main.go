package main

import "fmt"

func main() {
	// ARRAY

	// var ages [3]int = [3]int{20, 25, 30}
	// sama aja kek yg diatas
	var ages = [3]int{20, 25, 30}

	names := [4]string{"farrell", "hartono", "luigi", "mario"}
	names[1] = "Hartono"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// SLICES (use arrays under the hood)

	var scores = []int{100, 50, 60}
	scores[2] = 25
	// dia nambahin dirow paling belakang si append
	scores = append(scores, 85)
	fmt.Println(scores, len(scores))

	// slice ranges : dia cuman ngambil dari posisi 1-3 (inklusif)
	// inklusif jadi 1-3 array posisi ketiga tidak diambil
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)

}