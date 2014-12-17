// +build !darwin
package tidy

import (
	"C"
	"errors"
)

// This option controls the deletion or addition of the name attribute in elements where it can serve as anchor. If set to "true", a name attribute, if not already existing, is added along an existing id attribute if the DTD allows it. If set to "false", any existing name attribute is removed if an id attribute exists or has been added.
func (this *Tidy) AnchorAsName(val bool) (bool, error) {
	// C.TidyAnchorAsName = 87 it doesn't exists in tidyp
	return this.optSetBool(87, cBool(val))
}

// Can be used to modify behavior of -c (--clean yes) option. This option specifies if Tidy should merge nested <span> such as "<span><span>...</span></span>". The algorithm is identical to the one used by --merge-divs.
func (this *Tidy) MergeSpans(val int) (bool, error) {
	// C.TidyMergeSpans = 86 it doesn't exists in tidyp
	return this.optSetAutoBool(86, (C.ulong)(val))
}

// This option specifies if Tidy should preserve the well-formed entitites as found in the input.
func (this *Tidy) PreserveEntities(val bool) (bool, error) {
	// C.TidyPreserveEntities = 85 it doesn't exists in tidyp
	return this.optSetBool(85, cBool(val))
}

// This option specifies that tidy should sort attributes within an element using the specified sort algorithm. If set to "alpha", the algorithm is an ascending alphabetic sort.
func (this *Tidy) SortAttributes(val int) (bool, error) {
	switch val {
	case None, Alpha:
		// C.TidySortAttributes = 84 it doesn't exists in tidyp
		return this.optSetInt(84, (C.ulong)(val))
	}
	return false, errors.New("Argument val int is out of range (0,1)")
}
