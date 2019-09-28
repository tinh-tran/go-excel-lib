# go-excel-lib


For test 

listDataPerformance := make([]interface{}, len(listDataPer))
	for i, val := range listDataPer {
		listDataPerformance[i] = val
	}

	var sheetsSource = utils.PayrollSheetFetcher{}

	var sheet1 = &utils.PayrollSheet{
		MaxCol: 12,
		CurRow: 0,
		Name:   "Payroll",
		MaxRow: len(listData),
		Data:   listDataPayroll,
	}

	var sheet2 = &utils.PayrollSheet{
		MaxCol: 14,
		CurRow: 0,
		Name:   "Performance",
		MaxRow: len(listDataPerformance),
		Data:   listDataPerformance,
	}

	sheetsSource.Sheets = append(sheetsSource.Sheets, sheet1)
	sheetsSource.Sheets = append(sheetsSource.Sheets, sheet2)

	excel.ExportMultisheet(fmt.Sprintf("texxt2sheet_%v.xlsx", time.Now().Unix()), &sheetsSource)
