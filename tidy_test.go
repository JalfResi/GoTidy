package tidy

import (
	"testing"
	"strings"
)

func TestTidy(t *testing.T) {
	tdy := New()
	defer tdy.Free()

	var output string
	html := "<title id='bob' class='frank'>Hello, 世界</title><p>Foo!"
	match := "<html><head><meta name=\"generator\" content=\"HTML Tidy for Mac OS X (vers 25 March 2009), see www.w3.org\"><title id='bob' class='frank'>Hello,&#228;&#184;&#241;&#231;&#239;&#229;</title></head><body><p>Foo!</p></body></html>"

	output, _ = tdy.Tidy(html)
	
	if output != match {
		t.Log(output)
		t.FailNow()
	}
}

func TestAddXmlDecl(t *testing.T) {
	tdy := New()
	defer tdy.Free()

	var output string
	html := "<title id='bob' class='frank'>Hello, 世界</title><p>Foo!"
	match := "<?xml"

	tdy.AddXmlDecl(true)
	output, _ = tdy.Tidy(html)
	
	if !strings.Contains(output, match) {
		t.Log(output)
		t.FailNow()
	}
}

func TestTidyMark(t *testing.T) {
	tdy := New()
	defer tdy.Free()

	var output string
	html := "<title id='bob' class='frank'>Hello, 世界</title><p>Foo!"
	match := "<meta name=\"generator\" content=\"HTML Tidy for Mac OS X"

	tdy.TidyMark(false)
	output, _ = tdy.Tidy(html)
	
	if strings.Contains(output, match) {
		t.Log(output)
		t.FailNow()
	}

	tdy.TidyMark(true)
	output, _ = tdy.Tidy(html)
	
	if strings.Contains(output, match) {
		t.Log(output)
		t.FailNow()
	}
}