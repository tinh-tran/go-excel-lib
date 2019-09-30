package main

import (
	"fmt"
	"go-excel-lib/excel"
	"reflect"
)

type PayrollSheetFetcher struct {
	curId      int
	Sheets     []*PayrollSheet
	sheetNames []string
}

type PayrollSheet struct {
	Name   string
	CurRow int
	MaxRow int
	MaxCol int
	Data   []interface{}
	Header []string
}

func (a *PayrollSheetFetcher) GetSheetNames() []string {
	if len(a.sheetNames) > 0 {
		return a.sheetNames
	} else {
		hasil := []string{}
		for _, c := range a.Sheets {
			hasil = append(hasil, c.GetSheetName())
		}
		return hasil
	}
}

func (a *PayrollSheetFetcher) NextSheet() excel.Sheet {
	if a.curId < len(a.Sheets) {
		p := a.Sheets[a.curId]
		a.curId++
		return p
	} else {
		return nil
	}
}

func (a *PayrollSheet) NextRow() []excel.Value {
	if a.CurRow == a.MaxRow {
		return nil
	}
	var results []excel.Value
	if a.CurRow == 0 {
		var result excel.Value
		result.Type = "string"
		for _, val := range a.Header {
			result.Value = val
			results = append(results, result)
		}
	} else {
		val := reflect.ValueOf(a.Data[a.CurRow])
		for i := 0; i < val.NumField(); i++ {
			var result excel.Value
			result.Type = fmt.Sprintf("%v", reflect.TypeOf(val.Field(i).Interface()))
			result.Value = val.Field(i).Interface()
			results = append(results, result)
		}
	}
	a.CurRow++
	return results
}

func (a *PayrollSheet) GetSheetName() string {
	return a.Name
}
