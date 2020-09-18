package models

import (
	"gorm.io/gorm"
)

// DataPRE is a Presentation
func ParseDataREN(tokens []string) DataREN {
	ren := DataREN{}
	ren.Adsh = tokens[0]
	ren.Report = tokens[1]
	ren.Rfile = tokens[2]
	ren.Menucat = strOrNil(tokens[3])
	ren.Shortname = tokens[4]
	ren.Longname = tokens[5]
	ren.Roleuri = strOrNil(tokens[6])
	ren.Parentroleuri = strOrNil(tokens[7])
	ren.Parentreport = strOrNil(tokens[8])
	ren.Ultparentrpt = strOrNil(tokens[9])
	return ren
}

type DataREN struct {
	gorm.Model

	/**
	Accession Number. The 20-character string
	formed from the 18-digit number assigned by
	the Commission to each EDGAR submission.
	*/
	Adsh string `gorm:"index:idx_ren"`

	/**
	Represents the report grouping. The numeric
	value refers to the "R file" as computed by
	the renderer and posted on the EDGAR website.
	Note that in some situations the numbers skip.
	*/
	Report string `gorm:"index:idx_ren"`

	/**
	The type of interactive data file rendered
	on the EDGAR website, H = .htm file, X = .xml file.
	*/
	Rfile string

	/**
	If available, one of the menu categories as computed by the renderer:
	C=Cover, S=Statements, N=Notes, P=Policies,
	T=Tables, D=Details, O=Other, and U=Uncategorized.
	*/
	Menucat *string

	/**
	The portion of the long name used in the renderer menu.
	*/
	Shortname string

	/**
	The space-normalized text of the XBRL l
	ink "definition" element content.
	*/
	Longname string

	/**
	The XBRL "roleuri" of the role.
	*/
	Roleuri *string

	/**
	The XBRL roleuri of a role for which this
	role has a matching shortname prefix and a
	higher level menu category, as computed by
	the renderer.
	*/
	Parentroleuri *string

	/**
	The value of the report field for the
	role where roleuri equals this parentroleuri.
	*/
	Parentreport *string

	/**
	The highest ancestor report reachable by
	following parentreport relationships.
	A note (menucat = N) is its own ultimate parent.
	*/
	Ultparentrpt *string
}
