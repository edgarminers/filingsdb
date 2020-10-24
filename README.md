FilingsDB: Financial Statement and Notes Data Sets from sec.gov
===

> The Financial Statement and Notes Data Sets provide the text and detailed numeric information from all financial statements and their notes.  This data is extracted from exhibits to corporate financial reports filed with the SEC using eXtensible Business Reporting Language (XBRL). 

Description
---
A small golang script to download financial statements and notes to a local sqlite database from [sec.gov](https://www.sec.gov/dera/data/financial-statement-and-notes-data-set.html)

Requirements
---
- Go & Go modules
- Sqlite

Installation
--- 
Clone this repository, then run
```bash
$ go install
```
to compile the source to an executable in the `bin/` directory.

Usage
---
```
$ ./bin/filingsdb
Usage: filingsdb <year>
```
Give it a year and the script will download and store the data to a local `filings_$YEAR.db` sqlite database. This can take a while as there's a lot of data to ingest (the 2019 database clocks in at 16G) 

Database schema
---
The DB schema (tables, columns and types) follows the structure outlined in the [dataset official pdf documentation](https://www.sec.gov/files/aqfsn_1.pdf). The script also builds a convenient ticker <> cik table to make querying easier via join. Use this table with caution, as it's a snapshot of today's data. In the past a given ticker could potentially map to a different cik.

The submissions table (`data_sub`) contains one entry per submission. A filing's Accession Number (or `adsh`) is the main identifier used to join other facts tables.


Sample queries
---
Facebook's 10-K Accession Number
```sql
select data_tickers.*, data_subs.adsh, data_subs.form, data_subs.accepted
from data_subs 
left join data_tickers on data_subs.cik = data_tickers.cik
where data_tickers.ticker = "FB"
and form = '10-K';
```

Numbers from the filing
```sql
select *
from data_subs
left join data_nums on data_nums.adsh = data_subs.adsh
where data_subs.adsh = "0001326801-19-000009";
```

Identify companies in a given SIC (here biotech related)
```sql
select data_subs.adsh, data_subs.accepted, data_tickers.ticker , data_txts.tag, data_subs.form, data_subs.sic, data_subs.cik, data_txts.ddate, data_txts.value
from data_txts 
left join data_subs on data_txts.adsh = data_subs.adsh
left join data_tickers on data_tickers.cik = data_subs.cik
where data_txts.tag IN ("NatureOfOperations", "BusinessDescriptionAndBasisOfPresentationTextBlock", "BusinessDescriptionAndAccountingPoliciesTextBlock", "OrganizationConsolidationAndPresentationOfFinancialStatementsDisclosureTextBlock",
"OrganizationConsolidationAndPresentationOfFinancialStatementsDisclosureAndSignificantAccountingPoliciesTextBlock",
"OrganizationConsolidationBasisOfPresentationBusinessDescriptionAndAccountingPoliciesTextBlock"
) 
and data_subs.sic IN ("2834", "2835", "2836", "8071", "8731")
and data_subs.form = "10-K"
order by data_subs.accepted;
```
Welcome to the messy nature of XBRL. Different filers sometimes use different tags for the same thing.

License 
=== 
MIT