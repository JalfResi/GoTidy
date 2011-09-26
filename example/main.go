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
	t.Doctype("omit")
	t.MergeDivs(tidy.False)

	output, err := t.Tidy("<title>Foo</title><p>Foo!")
	if err != nil {
		fmt.Println(output)
		log.Fatal(err)
	}
	fmt.Println(output)
}