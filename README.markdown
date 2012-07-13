GoTidy
======
Simple package that wraps [libtidy](http://tidy.sourceforge.net/).

## Install Package

	go get github.com/JalfResi/GoTidy

And import into your program like so:

	import "github.com/JalfResi/GoTidy"

## Install Example Binary

	go get github.com/JalfResi/GoTidy/gotidy

Example usage of example binary:

	$ echo "<html><body><p>Rad" | gotidy

## Example

	package main

	import (
		"fmt"
		"log"
		"github.com/JalfResi/GoTidy"
	)

	func main() {

		// Create an instance of Tidy
		t := tidy.New()
		defer t.Free()

		// Set our options
		t.OutputXml(true)
		t.AddXmlDecl(true)
		t.QuoteAmpersand(true)
		t.TidyMark(false)

		// Tidy our source HTML!
		output, err := t.Tidy("<title id='bob' class='frank'>Welcome</title><p>Hello, 世界</p><p>Foo!")
		if err != nil {
			log.Fatal(err, output)
		}
		fmt.Println(output)
	}		

The package if very straight forward: simply create an instance of Tidy using New(), set the options you want and
then call the Tidy() instance method, passing it the string of HTML to tidy. The Tidy() method returns the output
(if any) and maybe an Error object.
	
Thanks
------
Thanks obviously to Dave Raggett for the awesome HTML Tidy library.
Thanks to Evan Shaw, Gustavo Niemeyer, Matt Kane's Brain and TJ Yang for help on [golang-nuts](http://groups.google.com/group/golang-nuts)
Thanks to shark and remy-o on irc.freenode.net#go-nuts with some of the trickier C casting stuff

To-do
-----
* Modify Tidy() method to accept Reader interface
* MOAR Documentation!
* Unit Tests ya fool!