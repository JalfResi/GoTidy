package tidy

/*
#cgo CFLAGS: -I/usr/include
#include <tidy.h>
#include <buffio.h>
#include <errno.h>
*/
import "C"
import (
	"os"
	"unsafe"
)

const False int = 0
const True int = 1
const Auto int = 2

// HTML, XHTML, XML Options

func (this *Tidy) AddXmlDecl(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyXmlDecl, CBool(val))
}

func (this *Tidy) AddXmlSpace(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyXmlSpace, CBool(val))
}

func (this *Tidy) AltText(val string) (bool, os.Error) {
	return this.optSetString(C.TidyXmlSpace, (*_Ctypedef_tmbchar)(C.CString(val)))
}

func (this *Tidy) AnchorAsName(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyAnchorAsName, CBool(val))
}

func (this *Tidy) AssumeXmlProcins(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyXmlPIs, CBool(val))
}

func (this *Tidy) Bare(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyMakeBare, CBool(val))
}

func (this *Tidy) Clean(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyMakeClean, CBool(val))
}

func (this *Tidy) CssPrefix(val string) (bool, os.Error) {
	return this.optSetString(C.TidyCSSPrefix, (*_Ctypedef_tmbchar)(C.CString(val)))
}

func (this *Tidy) DecorateInferredUl(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyDecorateInferredUL, CBool(val))
}

func (this *Tidy) Doctype(val string) (bool, os.Error) {
	/*
	var dtmode C.TidyDoctypeModes
	switch val {
		case "auto":
			dtmode = C.TidyDoctypeAuto
		case "omit":
			dtmode = C.TidyDoctypeOmit
		case "strict":
			dtmode = C.TidyDoctypeStrict
		case "loose", "transitional":
			dtmode = C.TidyDoctypeLoose
		default:
			dtmode = val
	}
	*/
	return this.optSetString(C.TidyDoctype, (*_Ctypedef_tmbchar)(C.CString(val)))
}

func (this *Tidy) DropEmptyParas(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyDropEmptyParas, CBool(val))
}

func (this *Tidy) DropFontTags(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyDropFontTags, CBool(val))
}

func (this *Tidy) DropProprietaryAttributes(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyDropPropAttrs, CBool(val))
}

func (this *Tidy) EncloseBlockText(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyEncloseBlockText, CBool(val))
}

func (this *Tidy) EncloseText(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyEncloseBodyText, CBool(val))
}

func (this *Tidy) EscapeCdata(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyEscapeCdata, CBool(val))
}

func (this *Tidy) FixBackslash(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyFixBackslash, CBool(val))
}

func (this *Tidy) FixBadComments(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyFixComments, CBool(val))
}

func (this *Tidy) FixUri(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyFixUri, CBool(val))
}

func (this *Tidy) HideComments(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyHideComments, CBool(val))
}

func (this *Tidy) HideEndtags(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyHideEndTags, CBool(val))
}

func (this *Tidy) IndentCdata(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyIndentCdata, CBool(val))
}

func (this *Tidy) InputXml(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyXmlTags, CBool(val))
}

func (this *Tidy) JoinClasses(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyJoinClasses, CBool(val))
}

func (this *Tidy) JoinStyles(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyJoinStyles, CBool(val))
}

func (this *Tidy) LiteralAttributes(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyLiteralAttribs, CBool(val))
}

func (this *Tidy) LogicalEmphasis(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyLogicalEmphasis, CBool(val))
}

func (this *Tidy) LowerLiterals(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyLowerLiterals, CBool(val))
}

func (this *Tidy) MergeDivs(val int) (bool, os.Error) {
	return this.optSetAutoBool(C.TidyMergeDivs, (_Ctypedef_ulong)(val))
}

func (this *Tidy) MergeSpans(val int) (bool, os.Error) {
	return this.optSetAutoBool(C.TidyMergeSpans, (_Ctypedef_ulong)(val))
}

func (this *Tidy) Ncr(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyNCR, CBool(val))
}

func (this *Tidy) NewBlocklevelTags(val string) (bool, os.Error) {
	return this.optSetString(C.TidyBlockTags, (*_Ctypedef_tmbchar)(C.CString(val)))
}

func (this *Tidy) NewEmptyTags(val string) (bool, os.Error) {
	return this.optSetString(C.TidyEmptyTags, (*_Ctypedef_tmbchar)(C.CString(val)))
}

func (this *Tidy) NewInlineTags(val string) (bool, os.Error) {
	return this.optSetString(C.TidyInlineTags, (*_Ctypedef_tmbchar)(C.CString(val)))
}

func (this *Tidy) NewPreTags(val string) (bool, os.Error) {
	return this.optSetString(C.TidyPreTags, (*_Ctypedef_tmbchar)(C.CString(val)))
}

func (this *Tidy) NumericEntities(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyNumEntities, CBool(val))
}

func (this *Tidy) OutputHtml(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyHtmlOut, CBool(val))
}

func (this *Tidy) OutputXhtml(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyXhtmlOut, CBool(val))
}

func (this *Tidy) OutputXml(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyXmlOut, CBool(val))
}

func (this *Tidy) PreserveEntities(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyPreserveEntities, CBool(val))
}

func (this *Tidy) QuoteAmpersand(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyQuoteAmpersand, CBool(val))
}

func (this *Tidy) QuoteMarks(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyQuoteMarks, CBool(val))
}

func (this *Tidy) QuoteNbsp(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyQuoteNbsp, CBool(val))
}

const KeepFirst int = 0
const KeepLast int = 1

func (this *Tidy) RepeatedAttributes(val int) (bool, os.Error) {
	switch val {
		case KeepFirst, KeepLast:
			return this.optSetInt(C.TidyDuplicateAttrs, (_Ctypedef_ulong)(val))
	}
	return false, os.NewError("Argument val int is out of range (0,1)")
}

func (this *Tidy) ReplaceColor(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyReplaceColor, CBool(val))
}

func (this *Tidy) ShowBodyOnly(val int) (bool, os.Error) {
	return this.optSetAutoBool(C.TidyBodyOnly, (_Ctypedef_ulong)(val))
}

func (this *Tidy) UppercaseAttributes(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyUpperCaseAttrs, CBool(val))
}

func (this *Tidy) UppercaseTags(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyUpperCaseTags, CBool(val))
}

func (this *Tidy) Word2000(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyWord2000, CBool(val))
}

// Diagnostics Options

const TidyClassic int = 0
const Priority1Checks int = 1
const Priority2Checks int = 2
const Priority3Checks int = 3

func (this *Tidy) AccessibilityCheck(val int) (bool, os.Error) {
	switch val {
		case TidyClassic, Priority1Checks, Priority2Checks, Priority3Checks:
			return this.optSetInt(C.TidyAccessibilityCheckLevel, (_Ctypedef_ulong)(val))
	}
	return false, os.NewError("Argument val int is out of range (0,1,2,3)")
}

func (this *Tidy) ShowWarnings(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyShowWarnings, CBool(val))
}

func (this *Tidy) ShowErrors(val int) (bool, os.Error) {
	return this.optSetInt(C.TidyShowErrors, (_Ctypedef_ulong)(val))
}

// Pretty Print Options

func (this *Tidy) BreakBeforeBr(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyBreakBeforeBR, CBool(val))
}

func (this *Tidy) Indent(val int) (bool, os.Error) {
	return this.optSetAutoBool(C.TidyIndentContent, (_Ctypedef_ulong)(val))
}

func (this *Tidy) IndentAttributes(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyIndentAttributes, CBool(val))
}

func (this *Tidy) IndentSpaces(val int) (bool, os.Error) {
	return this.optSetInt(C.TidyIndentSpaces, (_Ctypedef_ulong)(val))
}

func (this *Tidy) Markup(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyShowMarkup, CBool(val))
}

func (this *Tidy) PunctuationWrap(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyPunctWrap, CBool(val))
}

const None int = 0
const Alpha int = 1

func (this *Tidy) SortAttributes(val int) (bool, os.Error) {
	switch val {
		case None, Alpha:
			return this.optSetInt(C.TidySortAttributes, (_Ctypedef_ulong)(val))
	}
	return false, os.NewError("Argument val int is out of range (0,1)")
}

func (this *Tidy) Split(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyBurstSlides, CBool(val))
}

func (this *Tidy) TabSize(val int) (bool, os.Error) {
	return this.optSetInt(C.TidyTabSize, (_Ctypedef_ulong)(val))
}

func (this *Tidy) VerticalSpace(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyVertSpace, CBool(val))
}

func (this *Tidy) Wrap(val int) (bool, os.Error) {
	return this.optSetInt(C.TidyWrapLen, (_Ctypedef_ulong)(val))
}

func (this *Tidy) WrapAsp(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyWrapAsp, CBool(val))
}

func (this *Tidy) WrapAttributes(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyWrapAttVals, CBool(val))
}

func (this *Tidy) WrapJste(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyWrapJste, CBool(val))
}

func (this *Tidy) WrapPhp(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyWrapPhp, CBool(val))
}

func (this *Tidy) WrapScriptLiterals(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyWrapScriptlets, CBool(val))
}

func (this *Tidy) WrapSections(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyWrapSection, CBool(val))
}

// Character Encoding Options

func (this *Tidy) AsciiChars(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyAsciiChars, CBool(val))
}

func (this *Tidy) Language(val string) (bool, os.Error) {
	return this.optSetString(C.TidyLanguage, (*_Ctypedef_tmbchar)(C.CString(val)))
}

const LF int = 0
const CRLF int = 1
const CR int = 2

func (this *Tidy) Newline(val int) (bool, os.Error) {
	switch val {
		case LF, CRLF, CR:
			return this.optSetInt(C.TidyNewline, (_Ctypedef_ulong)(val))
	}
	return false, os.NewError("Argument val int is out of range (0,1,2)")
}

func (this *Tidy) OutputBom(val int) (bool, os.Error) {
	return this.optSetAutoBool(C.TidyOutputBOM, (_Ctypedef_ulong)(val))
}

// Miscellaneous Options

func (this *Tidy) ErrorFile(val string) (bool, os.Error) {
	return this.optSetString(C.TidyErrFile, (*_Ctypedef_tmbchar)(C.CString(val)))
}

func (this *Tidy) ForceOutput(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyForceOutput, CBool(val))
}

func (this *Tidy) GnuEmacs(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyEmacs, CBool(val))
}

func (this *Tidy) GnuEmacsFile(val string) (bool, os.Error) {
	return this.optSetString(C.TidyEmacsFile, (*_Ctypedef_tmbchar)(C.CString(val)))
}

func (this *Tidy) KeepTime(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyKeepFileTimes, CBool(val))
}

func (this *Tidy) OutputFile(val string) (bool, os.Error) {
	return this.optSetString(C.TidyOutFile, (*_Ctypedef_tmbchar)(C.CString(val)))
}

func (this *Tidy) Quiet(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyQuiet, CBool(val))
}

func (this *Tidy) SlideStyle(val string) (bool, os.Error) {
	return this.optSetString(C.TidySlideStyle, (*_Ctypedef_tmbchar)(C.CString(val)))
}

func (this *Tidy) TidyMark(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyMark, CBool(val))
}

func (this *Tidy) WriteBack(val bool) (bool, os.Error) {
	return this.optSetBool(C.TidyWriteBack, CBool(val))
}

const Autobool_false _Ctypedef_ulong = 0
const Autobool_true _Ctypedef_ulong = 1
const Autobool_auto _Ctypedef_ulong = 2

func (this *Tidy) optSetAutoBool(opt C.TidyOptionId, val _Ctypedef_ulong) (bool, os.Error) {
	switch val {
		case Autobool_false, Autobool_true, Autobool_auto:
			return this.optSetInt(opt, val)
	}
	return false, os.NewError("Argument val int is out of range (0,1,2)")
}


func (this *Tidy) optSetString(opt C.TidyOptionId, val *_Ctypedef_tmbchar) (bool, os.Error) {
	if C.tidyOptSetValue(this.tdoc, opt, val) == 1 {
		return false, nil	
	}
	return true, nil		
}

func (this *Tidy) optSetInt(opt C.TidyOptionId, val _Ctypedef_ulong) (bool, os.Error) {
	if C.tidyOptSetInt(this.tdoc, opt, val) == 1 {
		return false, nil	
	}
	return true, nil		
}

func (this *Tidy) optSetBool(opt C.TidyOptionId, val C.Bool) (bool, os.Error) {
	var rc C.int = -1
	if C.tidyOptSetBool(this.tdoc, opt, val) == 1 {
		rc = C.tidySetErrorBuffer(this.tdoc, &this.errbuf)	// Capture diagnostics
		if rc != 0 {
			return false, os.NewError( C.GoStringN( (*C.char)(unsafe.Pointer(this.errbuf.bp)), C.int(this.errbuf.size) ))
		}
	}
	return true, nil
}

func CBool(val bool) C.Bool {
	var v uint32 = 0
	if val {
		v = 1
	}
	return C.Bool(v)
}
