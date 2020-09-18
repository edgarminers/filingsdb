package models

import (
	"log"
	"strconv"

	"github.com/shopspring/decimal"
)

func parseOptInt(str string) *int {
	if str == "" {
		return nil
	}
	v, err := strconv.Atoi(str)
	if err != nil {
		log.Printf("cannot parse `%s` to int", str)
		def := 0
		return &def
	}
	return &v
}

func parseInt(str string) int {
	if str == "" {
		return -1
	}
	v, err := strconv.Atoi(str)
	if err != nil {
		log.Printf("cannot parse `%s` to *int", str)
		return 0
	}
	return v
}

func strOrNil(str string) *string {
	if str == "" {
		return nil
	}
	return &str
}

func parseDecimal(str string) *decimal.Decimal {
	if str == "" {
		return nil
	}
	v, err := decimal.NewFromString(str)
	if err != nil {
		log.Printf("cannot parse `%v` to decimal", str)
		return nil
	}
	return &v
}
