package utils

import (
	"fmt"
	"hrm-generate-excel/components/generate_payroll/models"
	"reflect"
)

type DummyDataFetcher struct {
	CurRow int
	MaxRow int
	Data   []models.PayrollDetail
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
