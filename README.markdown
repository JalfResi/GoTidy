GoTidy
======
Simple package that wraps [libtidy](http://tidy.sourceforge.net/).

Provides one function:
	func Tidy(htmlSource string) (string, os.Error)
	
Thanks
------
Thanks obviously to Dave Raggett for the awesome HTML Tidy library.
Thanks to Evan Shaw, Gustavo Niemeyer, Matt Kane's Brain and TJ Yang for help on [golang-nuts](http://groups.google.com/group/golang-nuts)
Thanks to shark and remy-o on irc.freenode.net#go-nuts with some of the trickier C casting stuff