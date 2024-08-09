package main

import (
	"fmt"
	"sort"
)

func main() {
	// greeting := "hello There Farrell!"

	// fmt.Println(strings.Contains(greeting, "Farrell"))
	// fmt.Println(strings.ReplaceAll(greeting, "hello", "Greetings"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, "rr"))
	// fmt.Println(strings.Split(greeting, " "))

	// function replace tidak menggantikan string nya jadi valuenya masih seperti hello There Farrell!
	// fmt.Println("Original string value =", greeting)

	ages := []int{45, 20, 35, 70, 85, 32, 53}
	
	sort.Ints(ages)
	fmt.Println(ages)

	// klo gk ketemu dia bakal nambah jadi indexnya di urutan tersebut
	index := sort.SearchInts(ages, 35)
	fmt.Println(index)

	names := []string{"farrell", "hartono", "auigi", "bario"}
	//klo ada huruf kapital dia sort yang besar terlebih dahulu
	names2 := []string{"farrell", "Hartono", "Auigi", "bario"}

	sort.Strings(names)
	sort.Strings(names2)
	
	fmt.Println(names, names2)

	fmt.Println(sort.SearchStrings(names, "auigi"))
}