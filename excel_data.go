package main

type ExcelData struct {
	Type        []string      `json:"type"`
	Value       []interface{} `json:"value"`
	TitleHeader []string      `json:"titleHeader"`
	Number      []int         `json:"number"`
}
