package models

import "gorm.io/gorm"

// DataTAG is a Tag
func ParseDataDIM(tokens []string) DataDIM {
	dim := DataDIM{}
	dim.Dimh = tokens[0]
	dim.Segments = tokens[1]
	dim.Segt = tokens[2] == "1"
	return dim
}

type DataDIM struct {
	gorm.Model

	/**
	MD5 hash of the segments field text.
	Although MD5 is unsuitable for cryptographic use,
	it is used here merely to limit the size of the primary key.
	*/
	Dimh string `gorm:"index:idx_dim"`

	/**
	Concatenation of tag names representing the
	axis and members appearing in the XBRL segments.

	Tag names have their first characters "Statement",
	last 4 characters "Axis", and last 6 characters
	"Member" or "Domain" truncated where they appear.

	Namespaces and prefixes are ignored because EDGAR
	validation guarantees that the local-names are
	unique with a submission.

	Each dimension is represented as the pair "{axis}={member};"
	and the axes concatenated in lexical order.

	Example: "LegalEntity=Xyz;Scenario=Restated;" represents
	the XBRL segment with dimension LegalEntityAxis
	and member XyzMember, dimension StatementScenarioAxis
	and member RestatedMember.
	*/
	Segments string

	/**
	TRUE if the segments field would have been longer
	than 1024 characters had it not been truncated, else
	FALSE.
	*/
	Segt bool
}
