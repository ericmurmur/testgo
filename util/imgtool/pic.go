package imgtool

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	//	"image/draw"
	"os"
)

func Show(f func(int, int) [][]uint8) {
	const (
		dx = 256
		dy = 256
	)
	data := f(dx, dy)
	m := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			m.Pix[i] = v
			m.Pix[i+1] = v
			m.Pix[i+2] = 255
			m.Pix[i+3] = 255
		}
	}
	ShowImage(m)
}

func ShowImage(m image.Image) {
	var buf bytes.Buffer
	err := png.Encode(&buf, m)
	if err != nil {
		panic(err)
	}

	//image.show

	enc := base64.StdEncoding.EncodeToString(buf.Bytes())
	fmt.Println("IMAGE:" + enc)

	oimg, _ := os.Create("mynew122.png")

	defer oimg.Close()
	png.Encode(oimg, m)

}

