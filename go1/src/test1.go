package main

import "fmt"

//
// This is a test
// add  a new line
//

func main() {
	var i int

	i = 1
	i := 1
	fmt.Printf("hello, world\n")
	fmt.Println("this is a new line")

	var tt string = "this is a string"
	if (true) {

		//tt = "hhh "
		fmt.Println(tt)
		var qq int
		qq = 100
		qq = qq + 1

	}

	for (i <= 3) {
		fmt.Println(i)
		i = i + 1
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
