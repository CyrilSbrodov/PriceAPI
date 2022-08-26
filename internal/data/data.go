package data

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"strconv"
)

func ImportXML(filename string) PriceTest { //функция импорта из xml
	file := "internal/data/" + filename + ".xml"
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var price PriceTest

	result, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	err = xml.Unmarshal(result, &price)
	if err != nil {
		panic(err)
	}

	return price
}

type Data struct {
	Name     string  `json:"name"`
	Unit     string  `json:"unit"`
	Sum      float64 `json:"sum"`
	PriceBuy float64 `json:"price_buy"`
	//PriceSell float64 `json:"price_sell"`
	//WarningMessage string  `json:"warning_message"`
}

type Price struct {
	ProductMSK []Data `json:"price_msk"`
	ProductSH  []Data `json:"price_sh"`
}

type PriceTest struct {
	Worksheet struct {
		Table struct {
			Row []struct {
				Cell []struct {
					Data struct {
						Text string `xml:",chardata"`
					} `xml:"Data"`
				} `xml:"Cell"`
			} `xml:"Row"`
		} `xml:"Table"`
	} `xml:"Worksheet"`
}

func CollectProduct(product, price PriceTest) []Data { //функция объединения
	var products []Data
	for i := 0; i < len(price.Worksheet.Table.Row); i++ {
		if len(price.Worksheet.Table.Row[i].Cell) > 3 {
			var product Data
			product.Name = price.Worksheet.Table.Row[i].Cell[1].Data.Text
			product.PriceBuy, _ = strconv.ParseFloat(price.Worksheet.Table.Row[i].Cell[2].Data.Text, 64)
			product.Unit = price.Worksheet.Table.Row[i].Cell[3].Data.Text
			products = append(products, product)
		}
	}

	for i := 0; i < len(product.Worksheet.Table.Row); i++ {
		if len(product.Worksheet.Table.Row[i].Cell) > 3 {
			for j := 0; j < len(products); j++ {
				if products[j].Name == product.Worksheet.Table.Row[i].Cell[0].Data.Text {
					products[j].Sum, _ = strconv.ParseFloat(product.Worksheet.Table.Row[i].Cell[3].Data.Text, 64)
				}
			}
		}
	}

	return products
}
