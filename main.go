// GoTidy is a cGo wrapper around libtidy
package tidy

/*
#cgo CFLAGS: -I/usr/include
#cgo LDFLAGS: -ltidy -L/usr/local/lib
#include <tidy.h>
#include <buffio.h>
#include <errno.h>
*/
import "C"
import (
	"errors"
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

func (this *Tidy) Tidy(htmlSource string) (string, error) {
	input := C.CString(htmlSource)
	defer C.free(unsafe.Pointer(input))

	var output C.TidyBuffer
	defer C.tidyBufFree(&output)

	var rc C.int = -1

	rc = C.tidySetErrorBuffer(this.tdoc, &this.errbuf) // Capture diagnostics

	if rc >= 0 {
		rc = C.tidyParseString(this.tdoc, (*C.tmbchar)(input)) // Parse the input
	}

	if rc >= 0 {
		rc = C.tidyCleanAndRepair(this.tdoc) // Tidy it up!	
	}

	if rc >= 0 {
		rc = C.tidyRunDiagnostics(this.tdoc) // Kvetch
	}

	if rc > 1 { // If error, force output.
		if C.tidyOptSetBool(this.tdoc, C.TidyForceOutput, C.yes) == 0 {
			rc = -1
		}
	}

	if rc >= 0 {
		rc = C.tidySaveBuffer(this.tdoc, &output) // Pretty Print
	}

	if rc >= 0 {
		if rc > 0 {
			err := errors.New(C.GoStringN((*C.char)(unsafe.Pointer(this.errbuf.bp)), C.int(this.errbuf.size)))
			return C.GoStringN((*C.char)(unsafe.Pointer(output.bp)), C.int(output.size)), err
		}
		return C.GoStringN((*C.char)(unsafe.Pointer(output.bp)), C.int(output.size)), nil
	}
	return "", os.NewSyscallError(fmt.Sprintf("A severe error (%d) occurred.\n", int(rc)), errors.New(string(rc)))
}
