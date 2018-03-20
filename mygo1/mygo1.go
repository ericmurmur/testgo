package main

import (
	"fmt"
	"./subdir"
)

func identityMat4() [16]float64 {
	return [...]float64{
		1, 0, 0, 0,
		0, 1, 4, 0,
		0, 0, 3, 0,
		0, 0, 0, 10 }
}

func main() {
	fmt.Println(identityMat4())

	fmt.Println("The value is", subdir.Plus(2,3))
}

