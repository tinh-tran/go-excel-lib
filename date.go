package main

import "time"

const (
	MJD_0      float64 = 2400000.5
	MJD_JD2000 float64 = 51544.5

	secondsInADay = float64((24 * time.Hour) / time.Second)
	nanosInADay   = float64((24 * time.Hour) / time.Nanosecond)
)

var (
	timeLocationUTC, _ = time.LoadLocation("UTC")

	unixEpoc = time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
	// In 1900 mode, Excel takes dates in floating point numbers of days starting with Jan 1 1900.
	// The days are not zero indexed, so Jan 1 1900 would be 1.
	// Except that Excel pretends that Feb 29, 1900 occurred to be compatible with a bug in Lotus 123.
	// So, this constant uses Dec 30, 1899 instead of Jan 1, 1900, so the diff will be correct.
	// http://www.cpearson.com/excel/datetime.htm
	excel1900Epoc = time.Date(1899, time.December, 30, 0, 0, 0, 0, time.UTC)
	excel1904Epoc = time.Date(1904, time.January, 1, 0, 0, 0, 0, time.UTC)
	// Days between epocs, including both off by one errors for 1900.
	daysBetween1970And1900 = float64(unixEpoc.Sub(excel1900Epoc) / (24 * time.Hour))
	daysBetween1970And1904 = float64(unixEpoc.Sub(excel1904Epoc) / (24 * time.Hour))
)

func TimeToExcelTime(t time.Time, date1904 bool) float64 {
	// Get the number of days since the unix epoc
	daysSinceUnixEpoc := float64(t.Unix()) / secondsInADay
	// Get the number of nanoseconds in days since Unix() is in seconds.
	nanosPart := float64(t.Nanosecond()) / nanosInADay
	// Add both together plus the number of days difference between unix and Excel epocs.
	var offsetDays float64
	if date1904 {
		offsetDays = daysBetween1970And1904
	} else {
		offsetDays = daysBetween1970And1900
	}
	daysSinceExcelEpoc := daysSinceUnixEpoc + offsetDays + nanosPart
	return daysSinceExcelEpoc
}
