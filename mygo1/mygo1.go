package main

import (
	"fmt"
	"./subdir"
	"reflect"
)

// return a slice of arrary float64
func identityMat4Slice() []float64 {
	return []float64{
		1, 0, 0, 0,
		0, 1, 4, 0,
		0, 0, 3, 0,
		0, 0, 0, 10 }
}

func identityMat4() [16]float64 {
	return [...]float64{
		1, 0, 0, 0,
		0, 1, 4, 0,
		0, 0, 3, 0,
		0, 0, 0, 10 }
}

func main() {
	var size= 15
	var s1= identityMat4Slice()
	var s0= identityMat4()
	var arr0= make([]int, size)
	var arr1= new([10]int)

	var s2= s0
	var s3= s1

	s2[3] = 1033
	s1[1] = 99
	s3[0] = -1

	fmt.Println(s0)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	//fmt.Println("identityMat4 type is ", reflect.TypeOf(s0).Kind())

	fmt.Println("identityMat4Slice type is ", reflect.TypeOf(s1).Kind())
	fmt.Println("identityMat4 type is ", reflect.TypeOf(s0).Kind())
	fmt.Println("s2 type is ", reflect.TypeOf(s2).Kind())

	fmt.Println("make ([]int, 10) type is ", reflect.TypeOf(arr0).Kind())
	fmt.Println("new ([]int, 10) type is ", reflect.TypeOf(arr1).Kind())

	fmt.Println("The value is", subdir.Plus(2, 3))
	fmt.Println("The value is", subdir.Plus(2, 3))

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		fmt.Print("num: ", num, ", ")
		sum += num
	}
	fmt.Println("sum:", sum)

	str1 := "中文Here is a string...."
	byte1 := []byte(str1)

	fmt.Println("str1:", str1[2:6], "len is ", len(str1), " byte1: ", byte1)

	a := []rune(str1)
	for i, r := range str1 {
		fmt.Printf("i%d r %d\n", i, r)
	}


	for i, v := range a {
		fmt.Printf("i%d r %x\n", i, v)
	}

	str2 := "\xF1\x82\x82\xF1\x65\x87"
	str2 = "\u82f1\u6587"
	str3 := "\xE6\x96\x87"

	for i, r := range str2 {
		fmt.Printf("str2 :: i%d r %x  %s\n", i, r, str2)
	}

	fmt.Printf("str2 len is %d %x %x %x \n", len(str2), str2[0], str2[1], str2[2])
	fmt.Printf("str3 len is %d %x %x %x \n", len(str3), str3[0], str3[1], str3[2])

	for i, r := range str3 {
		fmt.Printf("str3 :: i%d r %x  %s\n", i, r, str3)
	}


	for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
		fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	}
}

