package main

import (
	"flag"
	"errors"
	"log"
	"fmt"
	"math/big"
	"github.com/leekchan/accounting"
	"gopkg.in/kolomiichenko/cbr-currency-go.v1"
	"sort"
)

var currencyIn = flag.String("currency", "", "Currency name")
var valueIn = flag.Float64("value", 0, "Currency value")

func main() {
	flag.Parse()

	if len(*currencyIn) != 3 {
		log.Println("Invalid Currency")
		return
	}
	if *valueIn <= 0 {
		log.Println("Invalid Value")
		return
	}

	currencies, err := getCurrencies()
	check(err)
	rc, err := recalculate(currencies, *currencyIn, *valueIn)
	check(err)


	//sorting by code
	var keys []string
	for k := range rc {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	ac := accounting.Accounting{Precision: 2}
	for _,k := range keys {
		fmt.Printf("%s %s\n", k, ac.FormatMoneyBigFloat(rc[k]))
	}
}

func recalculate(currencies map[string]float64, currencyIn string, valueIn float64) (rc map[string]*big.Float, err error) {
	bigValueIn := big.NewFloat(valueIn)

	rc = make(map[string]*big.Float, 30)
	currencyVal, ok := currencies[currencyIn];
	if !ok {
		err = errors.New("Invalid Currency")
		return
	}

	bigCurrencyVal := big.NewFloat(currencyVal)
	for k, v := range currencies {
		if k == currencyIn {
			continue
		}

		bigV := big.NewFloat(v)
		z := new(big.Float)
		rc[k] = z.Quo(bigCurrencyVal, bigV).Mul(z, bigValueIn)
	}
	return
}

func getCurrencies() (currencies map[string]float64, err error) {
	cbr.UpdateCurrencyRates()
	rates := cbr.GetCurrencyRates()
	if len(rates) == 0 {
		err = errors.New("Failed to get currencies from cbr")
		return
	}

	currencies = make(map[string]float64, 30)
	currencies["RUB"] = 1
	for _, v := range rates {
		currencies[v.ISOCode] = v.Value
	}
	return
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
