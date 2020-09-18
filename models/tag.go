package models

import (
	"gorm.io/gorm"
)

// DataTAG is a Tag
func ParseDataTAG(tokens []string) DataTAG {
	tag := DataTAG{}
	tag.Tag = tokens[0]
	tag.Version = tokens[1]
	tag.Custom = tokens[2] == "1"
	tag.Abstract = tokens[3] == "1"
	tag.Datatype = strOrNil(tokens[4])
	tag.Iord = strOrNil(tokens[5])
	tag.Crdr = strOrNil(tokens[6])
	tag.Tlabel = strOrNil(tokens[7])
	tag.Doc = strOrNil(tokens[8])
	return tag
}

type DataTAG struct {
	gorm.Model
	/**
	The unique identifier (name) for a tag
	in a specific taxonomy release.
	*/
	Tag string `gorm:"index:idx_tag"`

	/**
	For a standard tag, an identifier for the
	taxonomy; otherwise the accession number
	where the tag was defined.
	*/
	Version string `gorm:"index:idx_tag"`

	/**
	1 if tag is custom (version=adsh), 0 if it is
	standard. Note: This flag is technically
	redundant with the version and adsh fields.
	*/
	Custom bool

	/**
	1 if the tag is not used to represent a numeric fact.
	*/
	Abstract bool

	/**
	If abstract=1, then NULL, otherwise the data
	type (e.g., monetary) for the tag.
	*/
	Datatype *string

	/**
	If abstract=1, then NULL; otherwise, "I" if
	the value is a point in time, or "D" if the value
	is a duration.
	*/
	Iord *string

	/**
	If datatype = monetary, then the tag's natural
	accounting balance from the perspective of the
	balance sheet or income statement (debit or credit);
	if not defined, then NULL.
	*/
	Crdr *string

	/**
	If a standard tag, then the label text provided
	by the taxonomy, otherwise the text provided by
	the filer. A tag which had neither would
	have a NULL value here.
	*/
	Tlabel *string

	/**
	The detailed definition for the tag, truncated
	to 2048 characters. If a standard tag, then the
	text provided by the taxonomy, otherwise the text
	assigned by the filer. Some tags have neither,
	in which case this field is NULL.
	*/
	Doc *string
}
