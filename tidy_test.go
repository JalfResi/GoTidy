package tidy

import (
	. "launchpad.net/gocheck"
	"testing"
)

// Hook up gocheck into the gotest runner.
func Test(t *testing.T) { TestingT(t) }

type S struct{}
var _ = Suite(&S{})


func (s *S) TestTidy(c *C) {
	tdy := New()
	defer tdy.Free()

	var output string
	var html string = "<title id='bob' class='frank'>Hello, 世界</title><p>Foo!"
	var match string = "<html>\n<head>\n<meta name=\"generator\" content=\n\"HTML Tidy for Mac OS X (vers 25 March 2009), see www.w3.org\">\n<title id='bob' class='frank'>Hello,\n&#228;&#184;&#241;&#231;&#239;&#229;</title>\n</head>\n<body>\n<p>Foo!</p>\n</body>\n</html>\n"

	output, _ = tdy.Tidy(html)
	c.Check(output, Matches, match)
}

func (s *S) TestAddXmlDecl(c *C) {
	tdy := New()
	defer tdy.Free()

	var output string
	html := "<title id='bob' class='frank'>Hello, 世界</title><p>Foo!"
	match := "^<?xml.*"

	tdy.AddXmlDecl(true)
	output, _ = tdy.Tidy(html)
	c.Check(output, Matches, match)

	tdy.AddXmlDecl(false)
	output, _ = tdy.Tidy(html)
	c.Check(output, Not(Matches), match)
}

func (s *S) TestTidyMark(c *C) {
	tdy := New()
	defer tdy.Free()

	var output string
	html := "<title id='bob' class='frank'>Hello, 世界</title><p>Foo!"
	match := "<meta name=\"generator\" content=\"HTML Tidy for Mac OS X"

	tdy.TidyMark(true)
	output, _ = tdy.Tidy(html)
	c.Check(output, Matches, match)

	tdy.TidyMark(false)
	output, _ = tdy.Tidy(html)
	c.Check(output, Not(Matches), match)
}