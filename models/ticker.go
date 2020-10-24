package models

type DataTicker struct {
	Cik       int    `json:"cik_str" gorm:"-"`
	CikString string `gorm:"column:cik" gorm:"index:idx_tickers_cik"`
	Ticker    string `json:"ticker" gorm:"index:idx_tickers_ticker`
	Name      string `json:"title"`
}

type DataTickers map[string]DataTicker
