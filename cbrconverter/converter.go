package cbrconverter

import (
	"errors"
	"gopkg.in/kolomiichenko/cbr-currency-go.v1"
	"math/big"
)


func Recalculate(currencies map[string]float64, currencyIn string, valueIn float64) (rc map[string]*big.Float, err error) {
	bigValueIn := big.NewFloat(valueIn)

	rc = make(map[string]*big.Float, 30)
	currencyVal, ok := currencies[currencyIn]
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

func GetCurrencies() (currencies map[string]float64, err error) {
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

