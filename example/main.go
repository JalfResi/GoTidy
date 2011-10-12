package main

import (
	"fmt"
	"log"
	"tidy"
)

func main() {
	t := tidy.New()
	defer t.Free()

	t.OutputXml(true)
	t.AddXmlDecl(true)
	t.QuoteAmpersand(true)
	t.TidyMark(false)
	t.CharEncoding(tidy.Utf8)
	t.AsciiChars(true)
	t.NumericEntities(true)
	t.FixUri(true)
	t.DropEmptyParas(true)
	t.DropFontTags(true)
	t.DropProprietaryAttributes(true)
	t.FixBackslash(true)
	t.JoinClasses(true)
	t.JoinStyles(true)
	t.ShowBodyOnly(tidy.True)

	output, err := t.Tidy("<title id='bob' class='frank'>Welcome</title><p>Hello, 世界</p><p>Foo!")
	if err != nil {
		log.Fatal(err, output)
	}
	fmt.Println(output)
}