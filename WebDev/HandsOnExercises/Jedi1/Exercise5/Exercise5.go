package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"text/template"
	"time"
)

var tpl *template.Template

type pass [][]string

type recordtype struct {
	Date     time.Time
	Open     float64
	High     float64
	Low      float64
	Close    float64
	Volume   float64
	AdjClose float64
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.html"))
}

func main() {
	descriptor, descerr := os.Open("table.csv")
	if descerr != nil {
		log.Fatalln(descerr)
	}

	reader := csv.NewReader(bufio.NewReader(descriptor))
	record, _ := reader.Read()
	// passer := pass{record}

	records := make([]recordtype, 0, len(record))

	for {
		record, recerr := reader.Read()
		if recerr != nil {
			if recerr == io.EOF {
				break
			}
		}

		date, _ := time.Parse("02-01-2006", record[0])
		open, _ := strconv.ParseFloat(record[1], 64)
		high, _ := strconv.ParseFloat(record[2], 64)
		low, _ := strconv.ParseFloat(record[3], 64)
		close, _ := strconv.ParseFloat(record[4], 64)
		volume, _ := strconv.ParseFloat(record[5], 64)
		adjclose, _ := strconv.ParseFloat(record[6], 64)

		records = append(records, recordtype{
			Date:     date,
			Open:     open,
			High:     high,
			Low:      low,
			Close:    close,
			Volume:   volume,
			AdjClose: adjclose,
		})
	}

	// for {
	// 	record, recerr := reader.Read()
	// 	if recerr != nil {
	// 		if recerr == io.EOF {
	// 			break
	// 		}
	// 	} else {
	// 		log.Fatalln(recerr)
	// 	}

	// 	for _, v := range record {
	// 		fmt.Printf("%s ", v)
	// 	}
	// 	fmt.Printf("\n")

	// }

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.html", records)
	if err != nil {
		log.Fatalln("An error occurred")
	}
}
