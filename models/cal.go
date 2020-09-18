package models

import (
	"gorm.io/gorm"
)

// DataCAL is a Calculations
func ParseDataCAL(tokens []string) DataCAL {
	cal := DataCAL{}
	cal.Adsh = tokens[0]
	cal.Grp = parseInt(tokens[1])
	cal.Arc = parseInt(tokens[2])
	cal.Negative = tokens[3] == "-1"
	cal.Ptag = tokens[4]
	cal.Pversion = tokens[5]
	cal.Ctag = tokens[6]
	cal.Cversion = tokens[7]

	return cal
}

type DataCAL struct {
	gorm.Model

	/**
	Accession Number. The 20-character string
	formed from the 18-digit number assigned by
	the Commission to each EDGAR submission.
	*/
	Adsh string `gorm:"index:idx_cal"`

	/**
	Sequential number for grouping arcs in a submission.
	*/
	Grp int `gorm:"index:idx_cal"`

	/**
	Sequential number for arcs within a
	group in a submission.
	*/
	Arc int `gorm:"index:idx_cal"`

	/**
	Indicates a weight of -1
	(TRUE if the arc is negative), but typically +1 (FALSE).
	*/
	Negative bool

	/**
	The tag for the parent of the arc
	*/
	Ptag string

	/**
	The version of the tag for the parent of the arc
	*/
	Pversion string

	/**
	The tag for the child of the arc
	*/
	Ctag string

	/**
	The version of the tag for the child of the arc
	*/
	Cversion string
}
