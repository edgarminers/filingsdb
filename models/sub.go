package models

import (
	"strings"

	"github.com/shopspring/decimal"
)

// DataSUB is a Submission
func ParseDataSUB(tokens []string) DataSUB {
	sub := DataSUB{}
	sub.Adsh = tokens[0]
	sub.Cik = strings.TrimLeft(tokens[1], "0")
	sub.Name = tokens[2]
	sub.Sic = tokens[3]
	sub.Countryba = tokens[4]
	sub.Stprba = strOrNil(tokens[5])
	sub.Cityba = tokens[6]
	sub.Zipba = strOrNil(tokens[7])
	sub.Bas1 = strOrNil(tokens[8])
	sub.Bas2 = strOrNil(tokens[9])
	sub.Baph = strOrNil(tokens[10])
	sub.Countryma = strOrNil(tokens[11])
	sub.Stprma = strOrNil(tokens[12])
	sub.Cityma = strOrNil(tokens[13])
	sub.Zipma = strOrNil(tokens[14])
	sub.Mas1 = strOrNil(tokens[15])
	sub.Mas2 = strOrNil(tokens[16])
	sub.Countryinc = tokens[17]
	sub.Stprinc = strOrNil(tokens[18])
	sub.Ein = strOrNil(tokens[19])
	sub.Former = strOrNil(tokens[20])
	sub.Changed = strOrNil(tokens[21])
	sub.Afs = strOrNil(tokens[22])
	sub.Wksi = tokens[23] == "1"
	sub.Fye = tokens[24]
	sub.Form = tokens[25]
	sub.Period = tokens[26]
	sub.Fy = tokens[27]
	sub.Fp = tokens[28]
	sub.Filed = tokens[29]
	sub.Accepted = tokens[30]
	sub.Prevrpt = tokens[31] == "1"
	sub.Detail = tokens[32] == "1"
	sub.Instance = tokens[33]
	sub.Nciks = parseInt(tokens[34])
	sub.Aciks = strOrNil(tokens[35])
	sub.Pubfloatusd = parseDecimal(tokens[36])
	sub.Floatdate = strOrNil(tokens[37])
	sub.Floataxis = strOrNil(tokens[38])
	sub.Floatmems = parseOptInt(tokens[39])
	return sub
}

type DataSUB struct {
	/**
	Accession Number.
	The 20  -character string formed from the
	18-digit number assigned by the Commission
	to each EDGAR submission
	*/
	Adsh string `gorm:"index:idx_subs_adsh"`

	/**
	Central Index Key (CIK).
	Ten digit number assigned by the Commission to
	each registrant that submits filings.
	*/
	Cik string `gorm:"index:idx_subs_cik"`

	/**
	Name of registrant.
	This corresponds to the name of the
	legal entity as recorded in EDGAR as
	of the filing date.
	*/
	Name string

	/**
	Standard Industrial Classification (SIC).
	Four digit code assigned by the Commission
	as of the filing date, indicating the registrant's
	type of business.
	*/
	Sic string `gorm:"index:idx_subs_sic"`

	/**
	The ISO 3166-1 country of the registrant's
	business address.
	*/
	Countryba string

	/**
	The state or province of the registrant's
	business address, if field countryba is US or CA
	*/
	Stprba *string

	/**
	The city of the registrant's business address.
	*/
	Cityba string

	/**
	The zip code of the registrant's business address.
	*/
	Zipba *string

	/**
	The first line of the street of the registrant's
	business address.
	*/
	Bas1 *string

	/**
	The second line of the street of the registrant's
	business address.
	*/
	Bas2 *string

	/**
	The phone number of the registrant's business address
	*/
	Baph *string

	/**
	The ISO 3166-1 country of the registrant's mailing address.
	*/
	Countryma *string

	/**
	The state or province of the registrant's
	mailing address, if field countryma is US or CA
	*/
	Stprma *string

	/**
	The city of the registrant's mailingaddress.
	*/
	Cityma *string

	/**
	The zip code of the registrant's mailingaddress.
	*/
	Zipma *string

	/**
	The first line of the street of the registrant's
	mailing address.
	*/
	Mas1 *string

	/**
	The second line of the street of the registrant's
	mailing address.
	*/
	Mas2 *string

	/**
	The country of incorporation for the registrant.
	*/
	Countryinc string

	/**
	The state or province of incorporation for the registrant, if countryinc is US or CA, otherwise NULL.
	*/
	Stprinc *string

	/**
	Employee Identification Number
	9 digit identification number assigned by
	the Internal Revenue Service to business
	entities operating in the United States
	*/
	Ein *string

	/**
	Most recent former name of the registrant, if any.
	*/
	Former *string

	/**
	Date of change from the former name, if any.
	*/
	Changed *string

	/**
	Filer status with the Commission at the time
	of submission:
	1-LAF=Large Accelerated,
	2-ACC=Accelerated,
	3-SRA=Smaller Reporting Accelerated,
	4-NON=Non-Accelerated,
	5-SML=Smaller Reporting Filer,
	NULL=not assigned.
	*/
	Afs *string

	/**
	Well Known Seasoned Issuer (WKSI).
	An issuer that meets specific Commission
	requirements at some point during a 60-day
	period preceding the date the issuer satisfies
	its obligation to update its shelf registration
	statement.
	*/
	Wksi bool

	/**
	Fiscal Year End Date.
	*/
	Fye string

	/**
	The submission type of the registrant's filing
	*/
	Form string

	/**
	Balance Sheet Date.
	*/
	Period string

	/**
	Fiscal Year Focus(as defined in EFM C h. 6)
	*/
	Fy string

	/**
	Fiscal Period Focus (as defined in EFM C h. 6)
	within Fiscal Year.
	The 10 -Q for the 1st, 2nd and 3rd quarters
	would have a fiscal period focus of Q1, Q2(or H1),
	and Q3 (or M9)  respectively, and a 10-K
	would have a fiscal period focus of FY.
	*/
	Fp string

	/**
	The date of the registrant's filing
	with the Commission.
	*/
	Filed string

	/**
	The acceptance date and time of the registrant's
	filing with the Commission. Filings accepted after
	5:30pm EST are considered filed on the following
	business day
	*/
	Accepted string

	/**
	Previous Report
	–TRUE indicates that the submission
	information was subsequentlyamended prior to the
	end cutoff date of the data set.
	*/
	Prevrpt bool

	/**
	TRUE indicates that the XBRL submission contains
	quantitative disclosures within the footnotes and
	schedules at the required detail level (e.g.,
	each amount)
	*/
	Detail bool

	/**
	The name of the submitted XBRL Instance Document
	(EX-101.INS) type data file. The name often begins
	with the company ticker symbol
	*/
	Instance string

	/**
	Number of Central Index Keys (CIK) of registrants
	(i.e., business units) included in the consolidating
	entity's submitted filing.
	*/
	Nciks int

	/**
	Additional CIKs of co-registrants included in a
	consolidating entity's EDGAR submission, separated
	by spaces. If there are no other co-registrants
	(i.e., nciks=1) , the value of aciks is  NULL.
	For a very small number of filers, the list of
	co-registrants is too long to fit in the field.
	Where this is the case, PARTIAL will appear at the
	end of the list indicating  that  not all
	co-registrants' CIKs are included in the field;
	users should refer to the complete submission file
	for all CIK information
	*/
	Aciks *string

	/**
	Public float, in USD, if provided in this submission
	*/
	Pubfloatusd *decimal.Decimal

	/**
	Date on which the public float was measured by the
	filer.
	*/
	Floatdate *string

	/**
	If the public float value was computed by summing
	across several tagged values, this indicates the
	nature of the summation.
	*/
	Floataxis *string

	/**
	If the public float was computed, the number of
	terms in the summation.
	*/
	Floatmems *int
}
