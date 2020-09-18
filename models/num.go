package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// DataNUM is a Number
func ParseDataNUM(tokens []string) DataNUM {
	num := DataNUM{}
	num.Adsh = tokens[0]
	num.Tag = tokens[1]
	num.Version = tokens[2]
	num.Ddate = tokens[3]
	num.Qtrs = parseInt(tokens[4])
	num.Uom = tokens[5]
	num.Dimh = tokens[6]
	num.Iprx = parseInt(tokens[7])
	num.Value = parseDecimal(tokens[8])
	num.Footnote = strOrNil(tokens[9])
	num.Footlen = parseInt(tokens[10])
	num.Dimn = parseInt(tokens[11])
	num.Coreg = strOrNil(tokens[12])
	num.Durp = *parseDecimal(tokens[13])
	num.Datp = *parseDecimal(tokens[14])
	num.Dcml = parseInt(tokens[15])
	return num
}

type DataNUM struct {
	gorm.Model

	/**
	Accession Number. The 20-character string
	formed from the 18-digit number assigned
	by the Commission to each EDGAR submission.
	*/
	Adsh string `gorm:"index:idx_num"`

	/**
	The unique identifier (name) for a
	tag in a specific taxonomy release.
	*/
	Tag string `gorm:"index:idx_num"`

	/**
	For a standard tag, an identifier for the
	taxonomy; otherwise the accession number
	where the tag was defined.
	*/
	Version string `gorm:"index:idx_num"`

	/**
	The end date for the data value, rounded
	to the nearest month end.
	*/
	Ddate string `gorm:"index:idx_num"`

	/**
	The count of the number of quarters
	represented by the data value, rounded to
	the nearest whole number. "0" indicates it
	is a point-in-time value.
	*/
	Qtrs int `gorm:"index:idx_num"`

	/**
	The unit of measure for the value.
	*/
	Uom string `gorm:"index:idx_num"`

	/**
	The 32-byte hexadecimal key for the
	dimensional information in the DIM data set.
	*/
	Dimh string `gorm:"index:idx_num"`

	/**
	A positive integer to distinguish different
	reported facts that otherwise would have the
	same primary key. For most purposes, data
	with iprx greater than 1 are not needed.
	The priority for the fact based on higher precision,
	closeness of the end date to a month end, and c
	loseness of the duration to a multiple of three months.
	See fields dcml, durp and datp below.
	*/
	Iprx int `gorm:"index:idx_num"`

	/**
	The value. This is not scaled, it is as found
	in the Interactive Data file, but is rounded
	to four digits to the right of the decimal point.
	*/
	Value *decimal.Decimal `sql:"type:decimal(20,8);"`

	/**
	The plain text of any superscripted footnotes
	on the value, if any, as shown on the statement
	page, truncated to 512 characters.
	*/
	Footnote *string

	/**
	Number of bytes in the plain text of the
	footnote prior to truncation; zero if no footnote.
	*/
	Footlen int

	/**
	Small integer representing the number of dimensions.
	Note that this value is a function of the dimension segments.
	*/
	Dimn int

	/**
	If specified, indicates a specific co-registrant, t
	he parent company, or other entity (e.g., guarantor).
	NULL indicates the consolidated entity. Note that this
	value is a function of the dimension segments.
	*/
	Coreg *string

	/**
	The difference between the reported fact duration
	and the quarter duration (qtrs), expressed as a
	fraction of 1. For example, a fact with duration of
	120 days rounded to a 91-day quarter has a durp
	value of 29/91 = +0.3187.
	*/
	Durp decimal.Decimal `sql:"type:decimal(20,8);"`

	/**
	The difference between the reported fact date and
	the month-end rounded date (ddate), expressed as a
	fraction of 1. For example, a fact reported for 29/Dec,
	with ddate rounded to 31/Dec, has a datp value of
	minus 2/31 = -0.0645.
	*/
	Datp decimal.Decimal `sql:"type:decimal(20,8);"`

	/**
	The value of the fact "decimals" attribute,
	with INF represented by 32767.
	*/
	Dcml int
}
