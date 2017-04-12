package main

import "fmt"

func main() {

	listOfNums := []float64{1,2,3,4,5}

	// Create addThemUp()

	fmt.Println("Sum :", addThemUp(listOfNums))

	// Get 2 vals from a function

	// Create next2Values()

	num1, num2 := next2Values(5)

	fmt.Println(num1, num2)

	// Send an undefined number of values to a fcuntion (Variadic function)

	// Create subtractThem(args ...type)

	fmt.Println(subtractThem(1,2,3,4,5))

	num3 := 3

	// function with no local vars is a closure

	doubleNum := func() int {
		num3 *= 2
		return num3
	}

	fmt.Println(doubleNum());
	fmt.Println(doubleNum());

	// Calling a recursive function

	// Create factorial()

	fmt.Println(factorial(3))

	// Defer exectures a function after the inclosing function finishes

	defer printTwo()
	printOne()

	// Use recover() to catch a division by 0 error

	fmt.Println(safeDiv(3, 0))
	fmt.Println(safeDiv(3, 2))

	// We can catch our own errors and recover with panic & recover

	demPanic()

}

// func name(paramName paramType) returnType {}

func addThemUp(numbers []float64) float64 {

	sum :=  0.0

	for _, value := range numbers {
		sum += value
	}

	return sum

}

func next2Values(number int) (int, int) {
	return number + 1, number + 2
}

// undefined number of values with args ...int

func subtractThem(args ...int) int {

	remainder := 0

	for _, value := range args {
		remainder -= value 
	}

	return remainder

}

func factorial(number int) int {

	if number == 0 {
		return 1
	}

	return number * factorial(number - 1)

}

// Used to demonstrate defer

func printOne(){ fmt.Println(1) }
func printTwo(){ fmt.Println(2) }

// If an error occurs we can catch the error with recover and allow
// code to continue to execute

func safeDiv(num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()
	
	solution := num1 / num2
	return solution

}

// Demonstrate how to call panic and handle it with recover

func demPanic() {
	defer func() {
		// If I didn't print the message nothing would show
		fmt.Println(recover())
	}()
	panic("PANIC")
}
