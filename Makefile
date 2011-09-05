include $(GOROOT)/src/Make.inc

TARG=tidy
CGOFILES=\
   main.go\

CGO_LDFLAGS=/usr/local/lib/libtidy.a

include $(GOROOT)/src/Make.pkg