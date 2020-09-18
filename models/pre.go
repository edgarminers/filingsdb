package models

import (
	"gorm.io/gorm"
)

// DataPRE is a Presentation
func ParseDataPRE(tokens []string) DataPRE {
	pre := DataPRE{}
	pre.Adsh = tokens[0]
	pre.Report = parseInt(tokens[1])
	pre.Line = parseInt(tokens[2])
	pre.Stmt = tokens[3]
	pre.Inpth = tokens[4]
	pre.Tag = tokens[5]
	pre.Version = tokens[6]
	pre.Prole = tokens[7]
	pre.Plabel = tokens[8]
	pre.Negating = tokens[9] == "1"
	return pre
}

type DataPRE struct {
	gorm.Model

	/**
	Accession Number. The 20-character string
	formed from the 18-digit number assigned by
	the Commission to each EDGAR submission.
	*/
	Adsh string `gorm:"index:idx_pre"`

	/**
	Represents the report grouping. The numeric
	value refers to the "R file" as computed by
	the renderer and posted on the EDGAR website.
	Note that in some situations the numbers skip.
	*/
	Report int `gorm:"index:idx_pre"`

	/**
	Represents the tag's presentation line order
	for a given report. Together with the statement
	and report field, presentation location,
	order and grouping can be derived.
	*/
	Line int `gorm:"index:idx_pre"`

	/**
	The financial statement location to which the value of the "report" field pertains.
	(
		CP = Cover Page, BS = Balance Sheet, IS = Income Statement,
		CF = Cash Flow, EQ = Equity,
		CI = Comprehensive Income, UN = Unclassifiable Statement).
	*/
	Stmt string

	/**
	1 indicates that the value was presented "parenthetically"
	instead of in fields within the financial statements.
	For example: Receivables (net of allowance for bad debts of $200 in 2012) $700.
	*/
	Inpth string

	/**
	The tag chosen by the filer for this line item.
	*/
	Tag string

	/**
	The taxonomy identifier if the tag is a standard tag, otherwise adsh.
	*/
	Version string

	/**
	The XBRL link "role" of the preferred label,
	using only the portion of the role URI after the last "/".
	*/
	Prole string

	/**
	The text presented on the line item, also known as a "preferred" label.
	*/
	Plabel string

	/**
	Flag to indicate whether the prole is treated as negating by the renderer.
	*/
	Negating bool
}
