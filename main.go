package main

import (
	"fmt"
	"time"
)

type User struct {
	Name     string
	Age      int
	BirthDay time.Time
}

func main() {

	var listDataPer = []User{}
	var data = User{
		Name:     "AAA",
		Age:      19,
		BirthDay: time.Now(),
	}
	for i := 0; i < 100; i++ {
		listDataPer = append(listDataPer, data)
	}
	listDataPerformance := make([]ExcelData, len(listDataPer))

	for i, val := range listDataPer {
		listDataPerformance[i].Value = []interface{}{val.Name, val.Age, val.BirthDay, fmt.Sprintf("SUM(B%v:B%v)", i, i+1)}
		listDataPerformance[i].Number = []int{1, 2, 3, 4}
		listDataPerformance[i].TitleHeader = []string{"Tên", "Tuổi", "Ngày sinh", "Test Formula"}
		listDataPerformance[i].Type = []string{"string", "int", "time.Time", "formula"}
	}
	var sheetsSource = SheetData{}
	var sheet1 = &DummySheetData{
		MaxCol: 12,
		CurRow: 0,
		Name:   "TestOk",
		MaxRow: len(listDataPer),
		Data:   listDataPerformance}

	//var sheet2 = &PayrollSheet{
	//	MaxCol: 14,
	//	CurRow: 0,
	//	Name:   "Performance",
	//	MaxRow: len(listDataPerformance),
	//	Data:   listDataPerformance,
	//}

	sheetsSource.Sheets = append(sheetsSource.Sheets, sheet1)
	//sheetsSource.Sheets = append(sheetsSource.Sheets, sheet2)

	ExportMultisheet(fmt.Sprintf("texxt2sheet_%v.xlsx", time.Now().Unix()), &sheetsSource)
}
