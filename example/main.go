package main

import (
	"flag"
	"fmt"
	"log"
	"io/ioutil"
	"os"
	"ourscienceistight/tidy"
)

var (
	debug  *bool = flag.Bool("debug", false, "Output debuggin messages")  
)

func main() {
	flag.Parse()

	t := tidy.New()
	defer t.Free()

	t.OutputXml(true)
	t.AddXmlDecl(false)
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

	in, _ := ioutil.ReadAll(os.Stdin)
	output, err := t.Tidy(string(in))
	if *debug == true && err != nil {
		log.Fatal(err, output)
	}
	fmt.Println(output)
}