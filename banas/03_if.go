package main

import "fmt"

func main() {

	yourAge := 18

	if yourAge >= 16 {
		fmt.Println("You can drive")
	} else {
		fmt.Println("You can't drive")
	}

	if yourAge >= 16 {
		fmt.Println("You can drive")
	} else if yourAge >= 18 {
		fmt.Println("You can vote")
	} else {
		fmt.Println("You can have fun")
	}

	switch yourAge {
		case 16: fmt.Println("Go drive")
		case 18: fmt.Println("Go vote")
		default: fmt.Println("Go have fun")
	}

}
