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

	//var listDataPer = []User{}

	var data = User{
		Name:     "AAA",
		Age:      19,
		BirthDay: time.Now(),
	}
	var sheet1 = &DummySheetData{
		MaxCol:    4,
		CurRow:    0,
		SheetName: "TestOk",
		MaxRow:    100}
	count := 0
	for i := 0; i < 100; i++ {
		count++
		sheet1.SetValue(count, 1, data.Name, "string")
		sheet1.SetValue(count, 2, data.Age, "int")
		sheet1.SetValue(count, 3, data.BirthDay, "time.Time")
		sheet1.SetValue(count, 4, fmt.Sprintf("=SUM(A%v+B%v", count+1, count+1), "formula")
	}

	var title = []string{"Tên", "Tuổi", "Ngày sinh", "Test Formula"}
	sheet1.SetHeader(title)

	var sheetsSource = SheetData{}

	sheetsSource.Sheets = append(sheetsSource.Sheets, sheet1)
	//sheetsSource.Sheets = append(sheetsSource.Sheets, sheet2)

	ExportMultisheet(fmt.Sprintf("texxt2sheet_%v.xlsx", time.Now().Unix()), &sheetsSource)
}
