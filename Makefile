include $(GOROOT)/src/Make.inc

TARG=tidy
CGOFILES=\
   main.go\

#CGO_LDFLAGS=/usr/local/lib/libtidy.a
CGO_OFILES=/Users/bendavies/Downloads/tidy/build/gmake/obj/*.o

include $(GOROOT)/src/Make.pkg