package main

import "fmt"

func main() {

	fmt.Println("Hello World!")

	var age int64 = 40

	var favNum float64 = 1.61803398875

	randNum := 1;
	fmt.Println(randNum);

	fmt.Println(age, " ", favNum)

	var numOne = 1.000
	var num99 = .999

	fmt.Println(numOne - num99)

	// Artihmetic Operators : +, 0, *, /, %

	fmt.Println("6 + 4 =", 6 + 4)
	fmt.Println("6 - 4 =", 6 - 4)
	fmt.Println("6 * 4 =", 6 * 4)
	fmt.Println("6 / 4 =", 6 / 4)
	fmt.Println("6 % 4 =", 6 % 4)

	// A constant is a variable with a value that can't be changed

	const pi float64 = 3.14159265359

	// You can declare multiple variables like this

	var (
		varA = 2
		varB = 3
	)

	fmt.Println(varA, varB)

	// Strings are a series of characters between " or `

	var myName string = "George"

	// Get string length

	fmt.Println(len(myName))

	// You can combine strings with +

	fmt.Println(myName + " is a robot")

	// Strings between " can contain escape symbols like \n for newline

	fmt.Println("I like \n \n")

	fmt.Println("Newlines")

	// Booleans can be either true or false

	var isOver40 bool = true

	fmt.Println(isOver40)

	// Printf is used for format printing (%f is for floats)

	fmt.Printf("%f \n", pi)

	// You can also define the decimal precision of a float

	fmt.Printf("%.3f \n", pi)

	// %T prints the data type

	fmt.Printf("%T \n", pi)

	// %t prints booleans

	fmt.Printf("%t \n", isOver40)

	// %d is used for integers

	fmt.Printf("%d \n", 100)

	// %b prints in binary

	fmt.Printf("%b \n", 100)

	// %c prints the character associated with a keycode

	fmt.Printf("%c \n", 44)

	// %x prints in hexcode

	fmt.Printf("%x \n", 17)

	// %e prints  in scientific notation

	fmt.Printf("%e \n", pi)

}
