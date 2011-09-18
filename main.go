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

const (
	NO = 0
	YES = 1
	AUTO = 2

	KeepFirst = 1 
	KeepLast = 2

	TidyClassic = 0
	Priority1Checks = 1
	Priority2Checks = 2
	Priority3Checks = 3

	None = 0
	Alpha = 1

	Raw = 1
	Ascii = 2
	Latin0 = 3
	Latin1 = 4
	Utf8 = 5
	Iso2022 = 6
	Mac = 7
	Win1252 = 8
	Ibm858 = 9
	Utf16le = 10
	Utf16be = 11
	Utf16 = 12
	Big5 = 13
	Shiftjis = 14

	LF = 1
	CRLF = 2
	CR = 3
)

type tidyOptions {
	// HTML, XHTML, XML Options
	AddXmlDecl bool,
	AddXmlSpace bool,
	AltText string,
	AnchorAsName bool,
	AssumeXmlProcins bool,
	Bare bool,
	Clean bool,
	CssPrefix string,
	DecorateInferredUl bool,
	Doctype string,
	DropEmptyParas bool,
	DropFontTags bool,
	DropProprietaryAttributes bool,
	EncloseBlockText bool,
	EncloseText bool,
	EscapeCdata bool,
	FixBackslash bool,
	FixBadComments bool,
	FixUri bool,
	HideComments bool,
	HideEndtags bool,
	IndentCdata bool,
	InputXml bool,
	JoinClasses bool,
	JoinStyles bool,
	LiteralAttributes bool,
	LogicalEmphasis bool,
	LowerLiterals bool,
	MergeDivs int,
	MergeSpans int,
	Ncr bool,
	NewBlocklevelTags string,
	NewEmptyTags string,
	NewInlineTags string,
	NewPreTags string,
	NumericEntities bool,
	OutputHtml bool,
	OutputXhtml bool,
	OutputXml bool,
	PreserveEntities bool,
	QuoteAmpersand bool,
	QuoteMarks bool,
	QuoteNbsp bool,
	RepeatedAttributes int,
	ReplaceColor bool,
	ShowBodyOnly int,
	UppercaseAttributes bool,
	UppercaseTags bool,
	Word2000 bool,
	// Diagnostics Options
	AccessibilityCheck int,
	ShowErrors int,
	ShowWarnings bool,
	// Pretty Print Options
	BreakBeforeBr bool,
	Indent int,
	IndentAttributes bool,
	IndentSpaces int,
	Markup bool,
	PunctuationWrap bool,
	SortAttributes int,
	Split bool,
	TabSize int,
	VerticalSpace bool,
	Wrap int,
	WrapAsp bool,
	WrapAttributes bool,
	WrapJste bool,
	WrapPhp bool,
	WrapScriptLiterals bool,
	WrapSections bool,
	// Character Encoding Options
	AsciiChars bool,
	CharEncoding int,
	InputEncoding int,
	Language string,
	Newline int,
	OutputBom int,
	OutputEncoding int,
	// Miscellaneous Options
	ErrorFile string,
	ForceOutput bool,
	GnuEmacs bool,
	GnuEmacsFile string,
	KeepTime bool,
	OutputFile string,
	Quiet bool,
	SlideStyle string,
	TidyMark bool,
	WriteBack bool
}

func New() *tidyOptions {
	return &tidyOptions {
		// HTML, XHTML, XML Options
		AddXmlDecl: 0,
		AddXmlSpace: 0,
		AltText: "",
		AnchorAsName: 1,
		AssumeXmlProcins: 0,
		Bare: 0,
		Clean: 0,
		CssPrefix: "",
		DecorateInferredUl: 0,
		Doctype: "auto",
		DropEmptyParas: 1,
		DropFontTags: 0,
		DropProprietaryAttributes: 0,
		EncloseBlockText: 0,
		EncloseText: 0,
		EscapeCdata: 0,
		FixBackslash: 1,
		FixBadComments: 1,
		FixUri: 1,
		HideComments: 0,
		HideEndtags: 0,
		IndentCdata: 0,
		InputXml: 0,
		JoinClasses: 0,
		JoinStyles: 1,
		LiteralAttributes: 0,
		LogicalEmphasis: 0,
		LowerLiterals: 1,
		MergeDivs: AUTO,
		MergeSpans: AUTO,
		Ncr: 1,
		NewBlocklevelTags: "",
		NewEmptyTags: "",
		NewInlineTags: "",
		NewPreTags: "",
		NumericEntities: 0,
		OutputHtml: 0,
		OutputXhtml: 0,
		OutputXml: 0,
		PreserveEntities: 0,
		QuoteAmpersand: 1,
		QuoteMarks: 0,
		QuoteNbsp: 1,
		RepeatedAttributes: KeepLast,
		ReplaceColor: 0,
		ShowBodyOnly: NO,
		UppercaseAttributes: 0,
		UppercaseTags: 0,
		Word2000: 0,
		// Diagnostics Options
		AccessibilityCheck: TidyClassic,
		ShowErrors: 6,
		ShowWarnings: 1,
		// Pretty Print Options
		BreakBeforeBr: 0,
		Indent: AUTO,
		IndentAttributes: 0,
		IndentSpaces: 2,
		Markup: 1,
		PunctuationWrap: 0,
		SortAttributes: None,
		Split: 0,
		TabSize: 8,
		VerticalSpace: 0,
		Wrap: 68,
		WrapAsp: 1,
		WrapAttributes: 0,
		WrapJste: 1,
		WrapPhp: 1,
		WrapScriptLiterals: 0,
		WrapSections: 1,
		// Character Encoding Options
		AsciiChars: 0,
		CharEncoding: Ascii,
		InputEncoding: Latin1,
		Language: "",
		Newline: LF,
		OutputBom: Auto,
		OutputEncoding: Ascii,
		// Miscellaneous Options
		ErrorFile: "",
		ForceOutput: 0,
		GnuEmacs: 0,
		GnuEmacsFile: "",
		KeepTime: 0,
		OutputFile: "",
		Quiet: 0,
		SlideStyle: "",
		TidyMark: 1,
		WriteBack: 0
	}
}

func Tidy(htmlSource string) (string, os.Error) {
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