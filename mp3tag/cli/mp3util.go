package cli

import (
	"github.com/bogem/id3v2"
	"log"
	"fmt"
	"strings"
	"path/filepath"
	"io/ioutil"
)

func init() {

	return
}


type CALLBACKUpdateTag func (fullpath string, album string, title string, artist string, genre string, tracknum int)



func updateMP3Tags(fullpath string, album string, title string, artist string, genre string, tracknum int) {

	tag, err := id3v2.Open(fullpath, id3v2.Options{Parse: true})
	if err != nil {
		log.Println("Error while opening mp3 file: ", err)
	}
	defer tag.Close()

	// Read frames.
	//fmt.Println(tag.Artist())
	//fmt.Println(tag.Title())

	// Set simple text frames.

/*	tag.SetAlbum(album)
	tag.SetArtist(artist)
	tag.SetTitle(title)
	tag.SetGenre(genre)

	tag.SetTrack(tracknum)
*/
	fmt.Println(tracknum, "album: ", album, "title: ", title, "artist: ", artist, "Genre:", genre)
	// Set comment frame.
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
			funcAction(filepath.Join(dirname, fname), filepath.Base(dirname), FilenameWithoutExt(fname), artist, genre, i)
		}

	}

	return
}
