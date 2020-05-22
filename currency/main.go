package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type stock struct {
	XMLName     xml.Name `xml:"STOCK"`
	SYMBOL      string   `xml:"SYMBOL"`
	DESC        string   `xml:"DESC"`
	LAST        string   `xml:"LAST"`
	PERNC       string   `xml:"PERNC"`
	PERNCNUMBER string   `xml:"PERNC_NUMBER"`
	LASTMOD     string   `xml:"LAST_MOD"`
}

type icpiyasa struct {
	XMLName xml.Name `xml:"ICPIYASA"`
	Stocks  []stock  `xml:"STOCK"`
}

func (s stock) String() string {
	return fmt.Sprintf("\nSymbol: %s\nDesc: %s\nLast: %s\nPernc: %s\nPernc numer: %s\nLast mod: %s\n", s.SYMBOL, s.DESC, s.LAST, s.PERNC, s.PERNCNUMBER, s.LASTMOD)
}

func main() {
	response, err := http.Get("http://realtime.paragaranti.com/asp/xml/icpiyasa.asp")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	//fmt.Println(string(data))
	var ip icpiyasa
	xml.Unmarshal(data, &ip)
	fmt.Println(ip.Stocks)

}

/*
<ICPIYASA>
<STOCK>
<SYMBOL>XU100</SYMBOL>
<DESC>BIST 100</DESC>
<LAST>103024</LAST>
<PERNC>0,37</PERNC>
<PERNC_NUMBER>0.37</PERNC_NUMBER>
<LAST_MOD>22.05.2020 18:10:10</LAST_MOD>
</STOCK>
...
*/
