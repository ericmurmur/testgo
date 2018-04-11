package main

import (
	"github.com/bogem/id3v2"
	"log"
	"fmt"
	"io/ioutil"
	//"os"
	"unsafe"
	"reflect"
	//"testing"
	"path/filepath"

	"./cli"
	//"os"
	//"github.com/spf13/cobra"
	//"github.com/spf13/viper"
	//"github.com/mitchellh/go-homedir"
)

type CALLBACKUpdateTag func (fullpath string, album string, title string, artist string, tracknum int)



func updateMP3Tags(fullpath string, album string, title string, artist string, tracknum int) {

	tag, err := id3v2.Open(fullpath, id3v2.Options{Parse: true})
	if err != nil {
		log.Println("Error while opening mp3 file: ", err)
	}
	defer tag.Close()

	// Read frames.
	//fmt.Println(tag.Artist())
	//fmt.Println(tag.Title())

	// Set simple text frames.
/*
	tag.SetAlbum(album)
	tag.SetArtist(artist)
	tag.SetTitle(title)

	tag.SetTrack(tracknum)
*/
	fmt.Println(tracknum, "album: ", album, "title: ", title, "artist: ", artist)
	// Set comment frame.
}

func EnumDir(dirname string, funcAction CALLBACKUpdateTag)  {

	// fmt.Println("dirname only is :: ", filepath.Base(dirname))

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Println(err)

		return //nil //make([]os.FileInfo, 0)
	}

	for i, file := range files {
		if file.IsDir() {
			log.Println(i, " DIR:: ", file.Name())
			EnumDir(dirname + "\\" + file.Name(), funcAction)

		} else {

			fname := file.Name()

			//fmt.Println(i, " ", filepath.Base(file.Name()), file.Name())
			funcAction(filepath.Join(dirname, fname), filepath.Base(dirname), fname, "artist", i)
		}

	}

	return
}

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
	//EnumDir("g:\\temp\\My Weird School 01 Miss Daisy Is Crazy")
	//EnumDir("g:\\temp\\My Weird School", updateMP3Tags)

	//var cmdline = cli.CLIStruct{}



	//cmdline.init()

	//cmdline.rootCmd.Execute()
	cli.Execute()
	//cli.rootCmd.
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

	tag.SetTrack(10)
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