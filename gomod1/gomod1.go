package main

import (
	"text/template"
	"os"

	"fmt"
	"unicode"


	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"golang.org/x/text/width"
)


type Joke struct {
	Who string
	Punchline string
}


func ExampleRemove() {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	s, _, _ := transform.String(t, "résumé")
	fmt.Println(s)

	// Output:
	// resume
}

func ExampleMap() {
	replaceHyphens := runes.Map(func(r rune) rune {
		if unicode.Is(unicode.Hyphen, r) {
			return '|'
		}
		return r
	})
	s, _, _ := transform.String(replaceHyphens, "a-b‐c⸗d﹣e")
	fmt.Println(s)

	// Output:
	// a|b|c|d|e
}

func ExampleIn() {
	// Convert Latin characters to their canonical form, while keeping other
	// width distinctions.
	t := runes.If(runes.In(unicode.Latin), width.Fold, nil)
	s, _, _ := transform.String(t, "ｱﾙｱﾉﾘｳ tech / アルアノリウ ｔｅｃｈ")
	fmt.Println(s)

	// Output:
	// ｱﾙｱﾉﾘｳ tech / アルアノリウ tech
}

func ExampleIf() {
	// Widen everything but ASCII.
	isASCII := func(r rune) bool { return r <= unicode.MaxASCII }
	t := runes.If(runes.Predicate(isASCII), nil, width.Widen)
	s, _, _ := transform.String(t, "ｱﾙｱﾉﾘｳ tech / 中國 / 5₩")
	fmt.Println(s)

	// Output:
	// アルアノリウ tech / 中國 / 5￦
}

func main() {
	t := template.New("Knock Knock Joke")
	text := `Knock Knock\nWho's there?
             {{.Who}}
             {{.Who}} who?
             {{.Punchline}}
            `
	t.Parse(text)

	jokes := []Joke{
		{"Etch", "Bless you!"},
		{"Cow goes", "No, cow goes moo!"},
	}

	for _, joke := range jokes {
		t.Execute(os.Stdout, joke)
	}


	ExampleIn()
	ExampleIf()
}