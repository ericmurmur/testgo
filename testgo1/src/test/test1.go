package main



import (
	"flag"
	"fmt"
	"errors"
	"../util.net/myimg"

	//"github.com/ericmurmur/testgo/awesomeProject/src/my.com/picutil"

	//_ "github.com/pkg/errors"


)
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
	name := flag.String("name", "", "name")
	flag.Parse()

	if err := greet(*name); err != nil {
		fmt.Printf("Failed asdas to greet you: %v", err)
	}


	myimg.Show(Pic)
}

func greet(name string) error {
	if name == "" {
		//return errors.New("HAHAH no name provided")
		return errors.New("HAHAHAHHAHAHA HAHHAHAHAHH HAHHAHHAH")
	}
	fmt.Printf("Hello %s! I'm not in the $GOPATH!\n", name)
	return nil
}