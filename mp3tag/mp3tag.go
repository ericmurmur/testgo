package main

import (
	"github.com/bogem/id3v2"
	"log"
	"fmt"

	//"os"
	"unsafe"
	"reflect"
	//"testing"

	"./cli"

	//"os"
	//"github.com/spf13/cobra"
	//"github.com/spf13/viper"
	//"github.com/mitchellh/go-homedir"
	//"sync"
)


func testSlice1() {

	str0 := "This is a full stringn!!"
	str1 := str0[2:8]
	typeOfStr1 := reflect.TypeOf(str1)


	arr0 := [5]int {0,1,2,3,4}
	arr1 := arr0[1:4]
	typeofArr1 := reflect.TypeOf(arr1)

	fmt.Println("str1 is ", str1, typeOfStr1, typeOfStr1.Kind(), arr1, reflect.TypeOf(arr1), typeofArr1.Kind())

}

type T1 struct {
	name string
}

type T2 struct {
	name string
}


type I1 interface {
	M1()
	//M2()
}

type I2 interface {
	M1()
	M2()
	//I3
}

type I3 interface {
	M2()
}

type T struct{
	v int
}

func (T) M1() {}
func (T) M2() {}


func main() {

	//EnumDir("g:\\temp\\My Weird School", updateMP3Tags)

	//var cmdline = cli.CLIStruct{}


	s1 := "this is a test string"
	sub1 := s1[1:7]
	run1 := []rune(s1)
	subr1 := run1[2:9]


	fmt.Println("sub1 is ", sub1, reflect.TypeOf(sub1))
	fmt.Println("run1 is ", run1, subr1, " ==> ", reflect.TypeOf(run1), reflect.TypeOf(subr1))
	//cmdline.init()

	//cmdline.rootCmd.Execute()
	cli.Execute()
	//cli.rootCmd.

/*
	var mm = map[string]map[interface{}]interface{} {"key1" : {1: "222"}, "key2" : {2: "2222222"}}

	var hits struct {
		sync.Mutex
		n int
	}

	hits.Lock()
	fmt.Println(mm)
*/
}


type Dog struct {

}


func (Dog) IsADog() bool {
	return true
}

type Cat struct {

}

func (Cat) IsACat() bool {
	return true
}

type Animal struct {
	*Dog
	*Cat
}

func main2() {

	var mamal *Animal = new(Animal)
	if mamal != nil {
		fmt.Println("")
		fmt.Println("Is mamal a Dog?", mamal.Dog.IsADog())
	}


	//EnumDir("g:\\temp\\My Weird School 01 Miss Daisy Is Crazy")
	//EnumDir("g:\\temp\\My Weird School")
	// Open file and parse tag in it.
	tag, err := id3v2.Open("g:\\temp\\01 I Hate School.mp3", id3v2.Options{Parse: true})
	if err != nil {
		log.Fatal("Error while opening mp3 file: ", err)
	}
	defer tag.Close()

	// Read frames.
	fmt.Println(tag.Artist())
	fmt.Println(tag.Title())

	// Set simple text frames.
	tag.SetArtist("New artist")
	tag.SetTitle("New title")

	//tag.SetTrack(10)
	// Set comment frame.
	/*comment := id3v2.CommentFrame{
		Encoding:    id3v2.EncodingUTF8,
		Language:    "eng",
		Description: "My opinion",
		Text:        "Very good song",
	}
	tag.AddCommentFrame(comment)
*/
	// Write it to file.
	if err = tag.Save(); err != nil {
		log.Fatal("Error while saving a tag: ", err)

	}

	arr1 := [6]int {0, 1, 2, 3, 4, 5 }
	p :=  unsafe.Pointer(&arr1[0])
	*(*int)(p) = 33

	v1 := uintptr(p)
	fmt.Printf("v1 is %x\n", v1)

	p = unsafe.Pointer(uintptr(p) + 8)
	*(*int)(p) = 999
	v1 = uintptr(p)

	fmt.Printf("v1 is %x\n", v1)

	*((*int)(unsafe.Pointer( uintptr(p)+8))) = 100

	var v0 int
	fmt.Println("arr1 is ", arr1, "pointer type size", reflect.TypeOf(v0).Size())

	//fmt.Println("this is ", p.(type))


	vs := []interface{}{T2(T1{"foo"}), T1{"foo"}, string(322), []byte("abł")}
	for _, v := range vs {
		fmt.Printf("interfaces %v %T\n", v, v)
	}


	arr0 := make([]int, 0)
	p0 := []int {};
	var z0 []int

	if arr0 == nil {
		fmt.Println("this array is nil")
	} else {
		fmt.Println("this array is NOT nil", len(arr0), cap(arr0), p0)
	}

	if p0 == nil {
		fmt.Println("this array is nil")
	} else {
		fmt.Println("this array is NOT nil", len(p0), cap(p0), p0)
	}
	if z0 == nil {
		fmt.Println("z0 this array is nil")
		fmt.Println("z0 this array is NOT nil", len(z0), cap(z0), z0)
	} else {
		fmt.Println("z0 this array is NOT nil", len(z0), cap(z0), z0)
	}



	var vv1 I1 = T{}
	var vv2 I2 = T{}
	_ = vv2
	_ = vv1

	log.Println("The end!!")
}