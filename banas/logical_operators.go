package main

import "fmt"

func main() {

	// Logical Operators

	fmt.Println("true && false =", true && false) // false
	fmt.Println("true || false =", true || false) // true
	fmt.Println("!true =", !true) // false

	// Fot loops

	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for j := 0; j < 5; j++ {
		fmt.Println(j);
	}

}
