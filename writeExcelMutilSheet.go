package main

type SheetData struct {
	curId      int
	Sheets     []*DummySheetData
	sheetNames []string
}

type DummySheetData struct {
	SheetName string
	CurRow    int
	MaxRow    int
	MaxCol    int
	Data      [][]Value
	Header    []string
	Value     []Value
}

func (a *SheetData) GetSheetNames() []string {
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

func (a *SheetData) NextSheet() Sheet {
	if a.curId < len(a.Sheets) {
		p := a.Sheets[a.curId]
		a.curId++
		return p
	} else {
		return nil
	}
}

func (a *DummySheetData) NextRow() []Value {
	//append header
	if a.CurRow == a.MaxRow+1 {
		return nil
	}
	var results []Value

	if a.CurRow == 0 {
		var result Value
		result.Type = "string"
		for i := 0; i < len(a.Header); i++ {
			result.Value = a.Header[i]
			results = append(results, result)
		}
		a.CurRow++
		return results
	}
	if a.CurRow < len(a.Data) {
		val := a.Data[a.CurRow-1]
		for _, valCell := range val {
			var result Value
			result.Type = valCell.Type
			result.Value = valCell.Value
			results = append(results, result)
		}
	} else {
		return nil
	}
	a.CurRow++
	return results
}

func (a *DummySheetData) GetSheetName() string {
	return a.SheetName
}

func (a *DummySheetData) SetHeader(titles []string) {
	a.Header = titles
}

func (a *DummySheetData) SetValue(row, cell int, value interface{}, typeCell string) {
	if value != nil {
		var data Value
		data.Value = value
		data.Type = typeCell
		a.Value = append(a.Value, data)
	}
	if len(a.Value) == a.MaxCol {
		a.SetRowValue(row, a.Value)
		a.Value = nil
	}
}

func (a *DummySheetData) SetRowValue(row int, value []Value) {
	if row >= 0 {
		a.Data = append(a.Data, value)
	}
}
