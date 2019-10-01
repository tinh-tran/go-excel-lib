package main

type SheetData struct {
	curId      int
	Sheets     []*DummySheetData
	sheetNames []string
}

type DummySheetData struct {
	Name   string
	CurRow int
	MaxRow int
	MaxCol int
	Data   []ExcelData
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
	if a.CurRow == a.MaxRow {
		return nil
	}
	var results []Value
	//val := reflect.ValueOf(a.Data[a.CurRow].Value)
	//valType := reflect.ValueOf(a.Data[a.CurRow].Type)

	for i := 0; i < len(a.Data[a.CurRow].Value); i++ {
		var result Value
		if a.CurRow == 0 {
			result.Type = a.Data[a.CurRow].Type[i]
			result.Value = a.Data[a.CurRow].TitleHeader[i]
		} else {
			result.Type = a.Data[a.CurRow].Type[i]
			result.Value = a.Data[a.CurRow].Value[i]
		}
		results = append(results, result)
	}
	a.CurRow++
	return results
}

func (a *DummySheetData) GetSheetName() string {
	return a.Name
}
