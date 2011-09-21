package tidy

const (
	NO   int32 = 0
	YES  int32 = 1
	AUTO int32 = 2

	KeepFirst int32 = 1 
	KeepLast  int32 = 2

	TidyClassic     int32 = 0
	Priority1Checks int32 = 1
	Priority2Checks int32 = 2
	Priority3Checks int32 = 3

	None  int32 = 0
	Alpha int32 = 1

	Raw      int32 = 1
	Ascii    int32 = 2
	Latin0   int32 = 3
	Latin1   int32 = 4
	Utf8     int32 = 5
	Iso2022  int32 = 6
	Mac      int32 = 7
	Win1252  int32 = 8
	Ibm858   int32 = 9
	Utf16le  int32 = 10
	Utf16be  int32 = 11
	Utf16    int32 = 12
	Big5     int32 = 13
	Shiftjis int32 = 14

	LF   int32 = 1
	CRLF int32 = 2
	CR   int32 = 3
)

type TidyOptions {
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

func New() *TidyOptions {
	return &TidyOptions {
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
