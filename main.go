package main

import (
	"fmt"
	"go-excel-lib/excel"
	"time"
)

type User struct {
	Name string
	Age  int
}

func main() {

	var listDataPer = []User{}
	var data = User{
		Name: "AAA",
		Age:  1,
	}
	for i := 0; i < 100; i++ {
		listDataPer = append(listDataPer, data)
	}
	listDataPerformance := make([]interface{}, len(listDataPer))

	for i, val := range listDataPer {
		listDataPerformance[i] = val
	}
	var sheetsSource = PayrollSheetFetcher{}
	var header = []string{"Tên", "Tuổi"}
	var sheet1 = &PayrollSheet{
		MaxCol: 12,
		CurRow: 0,
		Name:   "Payroll",
		MaxRow: len(listDataPer),
		Data:   listDataPerformance,
		Header: header,
	}

	//var sheet2 = &PayrollSheet{
	//	MaxCol: 14,
	//	CurRow: 0,
	//	Name:   "Performance",
	//	MaxRow: len(listDataPerformance),
	//	Data:   listDataPerformance,
	//}

	sheetsSource.Sheets = append(sheetsSource.Sheets, sheet1)
	//sheetsSource.Sheets = append(sheetsSource.Sheets, sheet2)

	excel.ExportMultisheet(fmt.Sprintf("texxt2sheet_%v.xlsx", time.Now().Unix()), &sheetsSource)
}
