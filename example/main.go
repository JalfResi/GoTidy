package main

import (
	"fmt"
	"log"
	"tidy"
)
/*
func main() {
	output, err := tidy.Tidy("<title>Foo</title><p>Foo!")
	if err != nil {
		fmt.Println(output)
		log.Fatal(err)
	}
	fmt.Println(output)
}
*/

func main() {
	t := tidy.New()
	defer t.Free()

	t.TidyMark(false)
	t.Doctype("strict")
	t.MergeDivs(tidy.False)
	t.NewBlocklevelTags("DIV")
	t.Indent(tidy.True)
	t.IndentSpaces(8)
	t.Wrap(10)
	t.RepeatedAttributes(tidy.KeepFirst)
	t.AccessibilityCheck(tidy.Priority3Checks)
	t.SortAttributes(tidy.Alpha)
	t.Newline(tidy.LF)
	t.CharEncoding(tidy.Raw)

	output, err := t.Tidy("<title id='bob' class='frank'>Hello, 世界</title><p>Foo!")
	if err != nil {
		fmt.Println(output)
		log.Fatal(err)
	}
	fmt.Println(output)
}