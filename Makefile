include $(GOROOT)/src/Make.inc

TARG=ourscienceistight/tidy
CGOFILES=\
   main.go\
   options.go\

#CGO_LDFLAGS=/usr/local/lib/libtidy.a
CGO_OFILES=/Users/bendavies/Downloads/tidy/build/gmake/obj/*.o

include $(GOROOT)/src/Make.pkg