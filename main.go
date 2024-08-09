package main

import "fmt"

func main() {

	name := "Farrell"
	age := 21

	// Print
	fmt.Print("hello, ")
	fmt.Print("world, \n")
	fmt.Print("new line \n")

	// Println
	fmt.Println("Hello World!")
	fmt.Println("my name is", name, "and my age is", age)

	// Printf (formatted strings) %_ = merupakan format specifier
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("you scored %f points! \n", 225.55)
	fmt.Printf("you scored %.2f points! \n", 225.55)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)

	fmt.Println("The save string is:", str)

}