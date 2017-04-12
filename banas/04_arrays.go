package main

import "fmt"

func main() {

	var favNums2[5] float64

	favNums2[0] = 163
	favNums2[1] = 78557
	favNums2[2] = 691
	favNums2[3] = 3.141
	favNums2[4] = 1.618

	fmt.Println(favNums2[3])

	favNums3 := [5]float64 { 1, 2, 3, 4, 5 }

	for i, value := range favNums3 {
		fmt.Println(value, i)
	}

	numSlice := []int {5,4,3,2,1}

	numSlice2 := numSlice[3:5]

	fmt.Println("numSlice2[0] =", numSlice2[0]) // 2

	fmt.Println("numSlice[:2] = ", numSlice[:2]) // min:2

	fmt.Println("numSlice[2:] =", numSlice[2:]) // 2:max

	// You can also create an empty slice and define the data type,
	// length (receive value of 0), capacity (max size)
	// make(type, first x indexes to receive default val of 0, maximum size (?))
	numSlice3 := make([]int, 5, 10)

	fmt.Println("numSlice3 = ", numSlice3)

	// You can copy a slice to another

	copy(numSlice3, numSlice)

	fmt.Println(numSlice3[0])

	// Append values to the end of a slice

	numSlice3 = append(numSlice3, 0, -1)

	fmt.Println(numSlice3[6])


	for i := 0; i < 20; i++ {
		numSlice3 = append(numSlice3, i, i)
		fmt.Println(numSlice3[i])
	}

	fmt.Println(numSlice3[:])

	numSlice4 := make([]int, 20, 20)

	fmt.Println(numSlice4[:])

	for i, value := range numSlice4 {
		fmt.Println("Before: ", i, value)
		numSlice4[i] = i
		fmt.Println("After: ", i, numSlice4[i])
	}

	fmt.Println(numSlice4[:])

}
