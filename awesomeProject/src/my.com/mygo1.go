package main

import (
	"math/rand"
	"fmt"
	"math"

)


func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}


const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}


func MySqrt(x float64) float64 {
	var z float64

	z = 1.0

	for i:=0; i<100; i++ {
		z = z - (z*z - x)/2/z

	}

	return z
}

func Pic(dx, dy int) [][]uint8 {
	var mypic [][]uint8

	mypic = make([][]uint8, dy)

	for i := 0; i< dy; i++ {
		mypic[i] = make([]uint8, dx)

		for j:=0; j<dx; j++ {
			mypic[i][j] = (uint8)(j*4 % 256)
		}

	}



	return mypic
}


func main() {
	fmt.Printf("Hello, world!!!.\n")
	fmt.Println("this is another line", 100, 200)
	//"aasdaldk;sa")

	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Println(math.Pi, add(100, 200))

	fmt.Println(swap("myfirst string \n", "my second string line\n"))

	if true {
		fmt.Println(swap("2ND :: myfirst string \n", " 2ND:: my second string line\n"))

	}

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	var i int
	sum := 0
	for i = 0; i < 1000 && sum < 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	fmt.Println("sqrt root of x is ", MySqrt(2.0), math.Sqrt(2.0))

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

	i = 5
	switch (i) {

	case 1, 3, 4, 5:
		fmt.Println("case 1,3,4,5 ")
	case 2:
		fmt.Println("case 2")
	default:

	}

	s := []int{2, 3, 5, 7, 11, 13}

	//s = s[1:4]
	fmt.Println(s[1:2])

	s = s[1:3]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	fmt.Println("len is ", len(s), cap(s))

	var pow= []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow = make([]int, 10)
	for i = range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	Show(Pic)
}
