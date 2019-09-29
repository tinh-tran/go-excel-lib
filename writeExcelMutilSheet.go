package main

import (
	"database/sql"
	"fmt"
	"go-excel-lib/excel"
	"reflect"
	"strconv"
	"time"
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

func (a *PayrollSheet) NextRow() []string {
	if a.CurRow == a.MaxRow {
		return nil
	}
	results := []string{}
	val := reflect.ValueOf(a.Data[a.CurRow])
	if a.CurRow == 0 {
		results = append(results, a.Header...)
	} else {
		for i := 0; i < val.NumField(); i++ {
			var value interface{}
			switch v := val.Field(i).Interface().(type) {
			case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
				value = fmt.Sprintf("%d", val.Field(i).Interface())
			case float32:
				value = fmt.Sprintf("%.2f", val.Field(i).Interface())
			case float64:
				value = fmt.Sprintf("%.2f", val.Field(i).Interface())
			case string:
				value = fmt.Sprintf("%v", val.Field(i).Interface())
			case []byte:
				value = fmt.Sprintf("%#v", val.Field(i).Interface())
			case time.Duration:
				value = strconv.FormatFloat(v.Seconds()/86400.0, 'f', -1, 32)
			case time.Time:
				value = fmt.Sprintf("%v", val.Field(i).Interface())
			case nil:
				value = ""
			case sql.NullString:
				var s sql.NullString
				if err := s.Scan(val.Field(i).Interface()); err != nil {
					//logger.Errorf(fmt.Sprintf("%v", err))
				}
				if s.Valid {
					value = fmt.Sprintf("%v", s.String)
				}
			default:
				value = fmt.Sprintf("%v", val.Field(i).Interface())
			}
			results = append(results, value.(string))
		}
	}
	a.CurRow++
	return results
}

func (a *PayrollSheet) GetSheetName() string {
	return a.Name
}
