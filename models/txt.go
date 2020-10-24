package models

import (
	"github.com/shopspring/decimal"
)

// DataTAG is a Tag
func ParseDataTXT(tokens []string) DataTXT {
	txt := DataTXT{}
	txt.Adsh = tokens[0]
	txt.Tag = tokens[1]
	txt.Version = tokens[2]
	txt.Ddate = tokens[3]
	txt.Qtrs = parseInt(tokens[4])
	txt.Iprx = parseInt(tokens[5])
	txt.Lang = tokens[6]
	txt.Dcml = parseInt(tokens[7])
	txt.Durp = *parseDecimal(tokens[8])
	txt.Datp = *parseDecimal(tokens[9])
	txt.Dimh = tokens[10]
	txt.Dimn = parseOptInt(tokens[11])
	txt.Coreg = strOrNil(tokens[12])
	txt.Escaped = tokens[13] == "1"
	txt.Srclen = parseInt(tokens[14])
	txt.Txtlen = parseOptInt(tokens[15])
	txt.Footnote = strOrNil(tokens[16])
	txt.Footlen = parseOptInt(tokens[17])
	txt.Context = tokens[18]
	txt.Value = strOrNil(tokens[19])
	return txt
}

type DataTXT struct {

	/**
	Accession Number. The 20-character string
	formed from the 18-digit number assigned by
	the Commission to each EDGAR submission.
	*/
	Adsh string `gorm:"index:idx_txts_adsh"`

	/**
	The unique identifier (name) for a tag in a
	 specific taxonomy release.
	*/
	Tag string `gorm:"index:idx_txts_tag"`

	/**
	For a standard tag, an identifier for the
	taxonomy; otherwise the accession number
	where the tag was defined. For example,
	"invest/2013" indicates that the tag is
	defined in the 2013 INVEST taxonomy.
	*/
	Version string //`gorm:"index:idx_txt"`

	/**
	The end date for the data value,
	rounded to the nearest month end.
	*/
	Ddate string //`gorm:"index:idx_txt"`

	/**
	The count of the number of quarters
	represented by the data value,
	rounded to the nearest whole number.
	A point in time value is represented by 0.
	*/
	Qtrs int //`gorm:"index:idx_txt"`

	/**
	A positive integer to distinguish different
	reported facts that otherwise would have
	the same primary key. For most purposes,
	data with iprx greater than 1 are not needed.
	The priority for the fact based on higher precision,
	closeness of the end date to a month end, and
	closeness of the duration to a multiple of
	three months. See fields dcml, durp and datp below.
	*/
	Iprx int //`gorm:"index:idx_txt"`

	/**
	The ISO language code of the fact content.
	*/
	Lang string

	/**
	The value of the fact "xml:lang" attribute,
	en-US represented by 32767, other "en"
	dialects having lower values, and other
	languages lower still.
	*/
	Dcml int

	/**
	The difference between the reported fact
	duration and the quarter duration (qtrs),
	expressed as a fraction of 1. For example,
	a fact with duration of 120 days rounded
	to a 91-day quarter has a durp
	value of 29/91 = +0.3187.
	*/
	Durp decimal.Decimal `sql:"type:decimal(20,8);"`

	/**
	The difference between the reported fact date
	and the month-end rounded date (ddate),
	expressed as a fraction of 1. For example,
	a fact reported for 29/Dec, with ddate rounded t
	o 31/Dec, has a datp value of minus 2/31 = - 0.0645.
	*/
	Datp decimal.Decimal `sql:"type:decimal(20,8);"`

	/**
	The 32-byte hexadecimal key for the
	dimensional information in the DIM data set.
	*/
	Dimh string `gorm:"index:idx_txts_dimh"`

	/**
	Small integer representing the number of dimensions,
	useful for sorting. Note that this value is
	function of the dimension segments.
	*/
	Dimn *int

	/**
	If specified, indicates a specific co-registrant,
	the parent company, or other entity (e.g., guarantor).
	NULL indicates the consolidated entity. Note that
	this value is a function of the dimension segments.
	*/
	Coreg *string

	/**
	Flag indicating whether the value has had tags removed.
	*/
	Escaped bool

	/**
	Number of bytes in the original, unprocessed value.
	Zero indicates a NULL value.
	*/
	Srclen int

	/**
	The original length of the whitespace normalized value,
	which may have been greater than 8192.
	*/
	Txtlen *int

	/**
	The plain text of any superscripted footnotes on
	the value, as shown on the page, truncated to
	512 characters, or if there is no footnote,
	then this field will be blank.
	*/
	Footnote *string

	/**
	Number of bytes in the plain text of the f
	ootnote prior to truncation.
	*/
	Footlen *int

	/**
	The value of the contextRef attribute in
	the source XBRL document, which can be
	used to recover the original HTML tagging if desired.
	*/
	Context string

	/**
	The value, with all whitespace normalized,
	that is, all sequences of line feeds,
	carriage returns, tabs, non-breaking spaces,
	and spaces having been collapsed to a single space,
	and no leading or trailing spaces.
	Escaped XML that appears in EDGAR "Text Block" tags
	is processed to remove all
	mark-up (comments, processing instructions, elements, attributes).
	The value is truncated to a maximum number of bytes.
	The resulting text is not intended for end user
	display but only for text analysis applications.
	*/
	Value *string
}
