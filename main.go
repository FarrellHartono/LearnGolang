package main

import "fmt"

func main() {
	// di go semuanya make for jadi while jadi for
	// WHILE LOOP
	// x := 0
	// for x < 5 {
	// 	fmt.Println("value of x is:", x)
	// 	x++
	// }

	// FOR LOOP
	// for i := 0; i < 5; i++{
	// 	fmt.Println("value of i is:", i)
	// }

		names := []string{"Farrell", "Hartono", "luigi", "mario"}

		// for i := 0; i < len(names); i++{
		// 	fmt.Println(names[i])
		// }

		// for index, value := range names{
		// 	fmt.Printf("the value at index %v is %v \n", index, value)
		// }

		// for _, value := range names{
		// 	fmt.Printf("the value is %v \n", value)
		// }

		// jika valuenya diganti jadi new string
		// tidak diganti tetep make names.
		for _, value := range names{
			fmt.Printf("the value is %v \n", value)
			value = "new string"
		}

		fmt.Println(names)
	
}