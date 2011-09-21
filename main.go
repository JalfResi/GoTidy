package tidy

/*
#cgo CFLAGS: -I/usr/include
#include <tidy.h>
#include <buffio.h>
#include <errno.h>
*/
import "C"
import (
	"fmt"
	"os"
	"unsafe"
)

type Tidy struct {
	tdoc   C.TidyDoc
	errbuf C.TidyBuffer
}

func New() *Tidy {
	t := &Tidy{}
	t.tdoc = C.tidyCreate()
	return t
}

func (this *Tidy) Free() {
	C.tidyBufFree(&this.errbuf)
	C.tidyRelease(this.tdoc)
}

func (this *Tidy) Tidy(htmlSource string) (string, os.Error) {
	input := C.CString(htmlSource)
	defer C.free(unsafe.Pointer(input))

	var output C.TidyBuffer
	defer C.tidyBufFree(&output)
	
	var rc C.int = -1
	
	if rc >= 0 {
    	rc = C.tidyParseString( this.tdoc, (*_Ctypedef_tmbchar)(input) )	// Parse the input
    }

	if rc >= 0 {
	    rc = C.tidyCleanAndRepair( this.tdoc )	// Tidy it up!	
	}

	if rc >= 0 {
    	rc = C.tidyRunDiagnostics( this.tdoc )	// Kvetch
    }

	if rc > 1 {		// If error, force output.
		if C.tidyOptSetBool(this.tdoc, C.TidyForceOutput, C.yes) == 0 {
			rc = -1
		}
	}

	if rc >= 0 {
    	rc = C.tidySaveBuffer( this.tdoc, &output )	// Pretty Print
    }

	if rc >= 0 {
    	if rc > 0 {
    		err := os.NewError(C.GoStringN((*C.char)(unsafe.Pointer(this.errbuf.bp)), C.int(this.errbuf.size)))
    		return C.GoStringN((*C.char)(unsafe.Pointer(output.bp)), C.int(output.size)), err
      	}
		return C.GoStringN((*C.char)(unsafe.Pointer(output.bp)), C.int(output.size)), nil
  	}
    return "", os.NewSyscallError(fmt.Sprintf( "A severe error (%d) occurred.\n", int(rc) ), int(rc))		
}

func TidyHtml(htmlSource string) (string, os.Error) {
	input := C.CString(htmlSource)
	defer C.free(unsafe.Pointer(input))

	var output C.TidyBuffer
	defer C.tidyBufFree( &output )

	var errbuf C.TidyBuffer
  	defer C.tidyBufFree( &errbuf )

	var rc C.int = -1
	var ok C.Bool

	var tdoc C.TidyDoc = C.tidyCreate()	// Initialize "document"
	defer C.tidyRelease( tdoc )

	ok = C.tidyOptSetBool( tdoc, C.TidyXhtmlOut, C.yes )  // Convert to XHTML

	if ok == 1 {
		rc = C.tidySetErrorBuffer( tdoc, &errbuf )	// Capture diagnostics
	}

	if rc >= 0 {
    	rc = C.tidyParseString( tdoc, (*_Ctypedef_tmbchar)(input) )	// Parse the input
    }

	if rc >= 0 {
	    rc = C.tidyCleanAndRepair( tdoc )	// Tidy it up!	
	}

	if rc >= 0 {
    	rc = C.tidyRunDiagnostics( tdoc )	// Kvetch
    }

	if rc > 1 {		// If error, force output.
		if C.tidyOptSetBool(tdoc, C.TidyForceOutput, C.yes) == 0 {
			rc = -1
		}
	}

	if rc >= 0 {
    	rc = C.tidySaveBuffer( tdoc, &output )	// Pretty Print
    }

	if rc >= 0 {
    	if rc > 0 {
    		return C.GoStringN((*C.char)(unsafe.Pointer(output.bp)), C.int(output.size)), os.NewError(C.GoStringN((*C.char)(unsafe.Pointer(errbuf.bp)), C.int(errbuf.size)))
      	}
		return C.GoStringN((*C.char)(unsafe.Pointer(output.bp)), C.int(output.size)), nil
  	}
    return "", os.NewSyscallError(fmt.Sprintf( "A severe error (%d) occurred.\n", int(rc) ), int(rc))
}