package cli

import (
	"github.com/bogem/id3v2"
	"log"
	"fmt"
	"strings"
	"path/filepath"
	"io/ioutil"
	"strconv"
)

func init() {

	return
}


type CALLBACKUpdateTag func (fullpath string, album string, title string, artist string, genre string, tracknum int)


func isUTF8(byteStr string) bool {
	var bUnicode = false

	for i:=0; i<len(byteStr);i++ {
		if byteStr[i]>0x80 {
			bUnicode = true
			break
		}
	}

	return bUnicode
}

func setTrack(tag id3v2.Tag, num int) {

	//tag.SetTrack(tracknum)

	tag.AddFrame(tag.CommonID("Track number/Position in set"), id3v2.TextFrame{Encoding: id3v2.EncodingUTF8, Text: strconv.Itoa(num)})
}

func updateMP3Tags(fullpath string, album string, title string, artist string, genre string, tracknum int) {

	tag, err := id3v2.Open(fullpath, id3v2.Options{Parse: true})
	if err != nil {
		log.Println("Error while opening mp3 file: ", err)
	}
	defer tag.Close()

	//fmt.Println("The tag version is : ", tag.Version(), tag.AllFrames())
	//tag.SetVersion(3)
	// Read frames.
	//fmt.Println(tag.Artist())
	//fmt.Println(tag.Title())

	// Set simple text frames.


	tag.SetAlbum(album)
	tag.SetArtist(artist)
	//tag.SetTitle(title)
	tag.SetGenre(genre)


	if isUTF8(title) {
		//tag.SetVersion(4)
		tag.AddFrame(tag.CommonID("Title"), id3v2.TextFrame{Encoding: id3v2.EncodingUTF16, Text: title})
	} else {
		tag.SetTitle(title)
	}


	//tag.SetTrack(tracknum)

	//tag.AddFrame(tag.CommonID("Track number/Position in set"), TextFrame{Encoding: tag.defaultEncoding, Text: strconv.Itoa(num)})

	fmt.Println(tracknum, "album: ", album, "title: ", title, "artist: ", artist, "Genre:", genre)
	// Set comment frame.

	if err = tag.Save(); err != nil {
		log.Println("Error while saving a tag: ", err)
	}
}


func FilenameWithoutExtension(fn string) string {
	return strings.TrimSuffix(fn, filepath.Ext(fn))
}

func FilenameWithoutExt(fn string) string {
	return fn[:len(fn)-len(filepath.Ext(fn))]
}

//EnumDir("g:\\temp\\My Weird School 01 Miss Daisy Is Crazy", updateMP3Tags)

func EnumDir(dirname string, artist string, genre string) {
	EnumDir2(dirname, artist, genre, updateMP3Tags)
}
func EnumDir2(dirname string, artist string, genre string, funcAction CALLBACKUpdateTag)  {

	// fmt.Println("dirname only is :: ", filepath.Base(dirname))

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Println(err)

		return //nil //make([]os.FileInfo, 0)
	}

	for i, file := range files {
		if file.IsDir() {
			log.Println(i, " DIR:: ", file.Name())
			EnumDir2(dirname + "\\" + file.Name(), artist, genre, funcAction)

		} else {

			fname := file.Name()

			//fmt.Println(i, " ", filepath.Base(file.Name()), file.Name())
			funcAction(filepath.Join(dirname, fname), filepath.Base(dirname), FilenameWithoutExt(fname), artist, genre, i+1)
		//	funcAction(filepath.Join(dirname, fname), filepath.Base(dirname), "track" + strconv.Itoa(i+1), artist, genre, i+1)

		}

	}

	return
}
