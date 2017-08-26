package main

import (
	"flag"
	"fmt"
	"github.com/leekchan/accounting"
	"log"
	"sort"
	"github.com/arduanov/cbr-bootcamp/cbrconverter"
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

	currencies, err := cbrconverter.GetCurrencies()
	check(err)
	rc, err := cbrconverter.Recalculate(currencies, *currencyIn, *valueIn)
	check(err)

	//sorting by code
	var keys []string
	for k := range rc {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	ac := accounting.Accounting{Precision: 2}
	for _, k := range keys {
		fmt.Printf("%s %s\n", k, ac.FormatMoneyBigFloat(rc[k]))
	}
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
