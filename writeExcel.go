package main

import (
	"fmt"
	"reflect"
)

type DummyDataFetcher struct {
	CurRow int
	MaxRow int
	Data   []User
	Header []string
}

func (d *DummyDataFetcher) NextRow() []string {
	if d.CurRow < d.MaxRow {
		res := []string{}
		v := reflect.ValueOf(d.Data[d.CurRow])
		for i := 0; i < v.NumField(); i++ {
			res = append(res, fmt.Sprintf("%v", v.Field(i).Interface()))
		}
		d.CurRow++
		return res

	}
	return nil
}
