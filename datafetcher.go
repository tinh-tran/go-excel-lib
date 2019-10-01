package main

// Fetch next row, if no more row exists return nil
type RowFetcher interface {
	NextRow() []Value
}

// Fetch next sheet, if no more row exists return nil
type SheetFetcher interface {
	NextSheet() Sheet
	GetSheetNames() []string
}
type SheetNamer interface {
	GetSheetName() string
}
type Sheet interface {
	RowFetcher
	SheetNamer
}
