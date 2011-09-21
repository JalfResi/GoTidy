// Created by cgo - DO NOT EDIT

//line options.go:1
package tidy


import (
	"os"
	"unsafe"
)


func (this *Tidy) AddXmlDecl(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyXmlDecl, CBool(val))
}

func (this *Tidy) AddXmlSpace(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyXmlSpace, CBool(val))
}

func (this *Tidy) AltText(val string) (bool, os.Error) {
	return this.optSetString(_Cconst_TidyXmlSpace, (*_Ctypedef_tmbchar)(_Cfunc_CString(val)))
}

func (this *Tidy) AnchorAsName(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyAnchorAsName, CBool(val))
}

func (this *Tidy) AssumeXmlProcins(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyXmlPIs, CBool(val))
}

func (this *Tidy) Bare(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyMakeBare, CBool(val))
}

func (this *Tidy) Clean(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyMakeClean, CBool(val))
}

func (this *Tidy) CssPrefix(val string) (bool, os.Error) {
	return this.optSetString(_Cconst_TidyCSSPrefix, (*_Ctypedef_tmbchar)(_Cfunc_CString(val)))
}

func (this *Tidy) DecorateInferredUl(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyDecorateInferredUL, CBool(val))
}

func (this *Tidy) Doctype(val string) (bool, os.Error) {
	var dtmode _Ctypedef_TidyDoctypeModes
	switch val {
	case "auto":
		dtmode = _Cconst_TidyDoctypeAuto
	case "omit":
		dtmode = _Cconst_TidyDoctypeOmit
	case "strict":
		dtmode = _Cconst_TidyDoctypeStrict
	case "loose", "transitional":
		dtmode = _Cconst_TidyDoctypeLoose
	}
	return this.optSetInt(_Cconst_TidyDoctype, (_Ctypedef_ulong)(dtmode))
}

func (this *Tidy) DropEmptyParas(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyDropEmptyParas, CBool(val))
}

func (this *Tidy) DropFontTags(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyDropFontTags, CBool(val))
}

func (this *Tidy) DropProprietaryAttributes(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyDropPropAttrs, CBool(val))
}

func (this *Tidy) EncloseBlockText(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyEncloseBlockText, CBool(val))
}

func (this *Tidy) EncloseText(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyEncloseBodyText, CBool(val))
}

func (this *Tidy) EscapeCdata(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyEscapeCdata, CBool(val))
}

func (this *Tidy) FixBackslash(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyFixBackslash, CBool(val))
}

func (this *Tidy) FixBadComments(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyFixComments, CBool(val))
}

func (this *Tidy) FixUri(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyFixUri, CBool(val))
}

func (this *Tidy) HideComments(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyHideComments, CBool(val))
}

func (this *Tidy) HideEndtags(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyHideEndTags, CBool(val))
}

func (this *Tidy) IndentCdata(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyIndentCdata, CBool(val))
}

func (this *Tidy) InputXml(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyXmlTags, CBool(val))
}

func (this *Tidy) JoinClasses(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyJoinClasses, CBool(val))
}

func (this *Tidy) JoinStyles(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyJoinStyles, CBool(val))
}

func (this *Tidy) LiteralAttributes(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyLiteralAttribs, CBool(val))
}

func (this *Tidy) LogicalEmphasis(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyLogicalEmphasis, CBool(val))
}

func (this *Tidy) LowerLiterals(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyLowerLiterals, CBool(val))
}

func (this *Tidy) Ncr(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyNCR, CBool(val))
}

func (this *Tidy) NumericEntities(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyNumEntities, CBool(val))
}

func (this *Tidy) OutputHtml(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyHtmlOut, CBool(val))
}

func (this *Tidy) OutputXhtml(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyXhtmlOut, CBool(val))
}

func (this *Tidy) OutputXml(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyXmlOut, CBool(val))
}

func (this *Tidy) PreserveEntities(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyPreserveEntities, CBool(val))
}

func (this *Tidy) QuoteAmpersand(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyQuoteAmpersand, CBool(val))
}

func (this *Tidy) QuoteMarks(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyQuoteMarks, CBool(val))
}

func (this *Tidy) QuoteNbsp(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyQuoteNbsp, CBool(val))
}

func (this *Tidy) ReplaceColor(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyReplaceColor, CBool(val))
}

func (this *Tidy) UppercaseAttributes(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyUpperCaseAttrs, CBool(val))
}

func (this *Tidy) UppercaseTags(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyUpperCaseTags, CBool(val))
}

func (this *Tidy) Word2000(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyWord2000, CBool(val))
}


func (this *Tidy) ShowWarnings(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyShowWarnings, CBool(val))
}


func (this *Tidy) BreakBeforeBr(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyBreakBeforeBR, CBool(val))
}

func (this *Tidy) IndentAttributes(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyIndentAttributes, CBool(val))
}

func (this *Tidy) Markup(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyShowMarkup, CBool(val))
}

func (this *Tidy) PunctuationWrap(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyPunctWrap, CBool(val))
}

func (this *Tidy) Split(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyBurstSlides, CBool(val))
}

func (this *Tidy) VerticalSpace(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyVertSpace, CBool(val))
}

func (this *Tidy) WrapAsp(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyWrapAsp, CBool(val))
}

func (this *Tidy) WrapAttributes(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyWrapAttVals, CBool(val))
}

func (this *Tidy) WrapJste(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyWrapJste, CBool(val))
}

func (this *Tidy) WrapPhp(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyWrapPhp, CBool(val))
}

func (this *Tidy) WrapScriptLiterals(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyWrapScriptlets, CBool(val))
}

func (this *Tidy) WrapSections(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyWrapSection, CBool(val))
}


func (this *Tidy) AsciiChars(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyAsciiChars, CBool(val))
}

func (this *Tidy) Language(val string) (bool, os.Error) {
	return this.optSetString(_Cconst_TidyLanguage, (*_Ctypedef_tmbchar)(_Cfunc_CString(val)))
}


func (this *Tidy) ErrorFile(val string) (bool, os.Error) {
	return this.optSetString(_Cconst_TidyErrFile, (*_Ctypedef_tmbchar)(_Cfunc_CString(val)))
}

func (this *Tidy) ForceOutput(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyForceOutput, CBool(val))
}

func (this *Tidy) GnuEmacs(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyEmacs, CBool(val))
}

func (this *Tidy) GnuEmacsFile(val string) (bool, os.Error) {
	return this.optSetString(_Cconst_TidyEmacsFile, (*_Ctypedef_tmbchar)(_Cfunc_CString(val)))
}

func (this *Tidy) KeepTime(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyKeepFileTimes, CBool(val))
}

func (this *Tidy) OutputFile(val string) (bool, os.Error) {
	return this.optSetString(_Cconst_TidyOutFile, (*_Ctypedef_tmbchar)(_Cfunc_CString(val)))
}

func (this *Tidy) Quiet(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyQuiet, CBool(val))
}

func (this *Tidy) SlideStyle(val string) (bool, os.Error) {
	return this.optSetString(_Cconst_TidySlideStyle, (*_Ctypedef_tmbchar)(_Cfunc_CString(val)))
}

func (this *Tidy) TidyMark(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyMark, CBool(val))
}

func (this *Tidy) WriteBack(val bool) (bool, os.Error) {
	return this.optSetBool(_Cconst_TidyWriteBack, CBool(val))
}


func (this *Tidy) optSetString(opt _Ctypedef_TidyOptionId, val *_Ctypedef_tmbchar) (bool, os.Error) {
	if _Cfunc_tidyOptSetValue(this.tdoc, opt, val) == 1 {
		return false, nil
	}
	return true, nil
}

func (this *Tidy) optSetInt(opt _Ctypedef_TidyOptionId, val _Ctypedef_ulong) (bool, os.Error) {
	if _Cfunc_tidyOptSetInt(this.tdoc, opt, val) == 1 {
		return false, nil
	}
	return true, nil
}

func (this *Tidy) optSetBool(opt _Ctypedef_TidyOptionId, val _Ctypedef_Bool) (bool, os.Error) {
	var rc _Ctype_int = -1
	if _Cfunc_tidyOptSetBool(this.tdoc, opt, val) == 1 {
		rc = _Cfunc_tidySetErrorBuffer(this.tdoc, &this.errbuf)
		if rc != 0 {
			return false, os.NewError(_Cfunc_GoStringN((*_Ctype_char)(unsafe.Pointer(this.errbuf.bp)), _Ctype_int(this.errbuf.size)))
		}
	}
	return true, nil
}

func CBool(val bool) _Ctypedef_Bool {
	var v uint32 = 0
	if val {
		v = 1
	}
	return _Ctypedef_Bool(v)
}
