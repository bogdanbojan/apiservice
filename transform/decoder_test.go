package transform

import (
	"apiserv/read"
	"reflect"
	"testing"
)

func TestProcessExport(t *testing.T) {}
func TestMapExport(t *testing.T)     {}
func TestFormatExport(t *testing.T)  {}

//func TestSortRecords(t *testing.T) {
//	for _, dc := range sortCases {
//		t.Run(dc.name, func(t *testing.T) {
//			got := sortRecords(dc.records)
//			want := dc.want
//			assertEqual(t, got, want)
//		})
//	}
//}

func TestRemoveDuplicates(t *testing.T) {
	for _, dc := range duplicatesCases {
		t.Run(dc.name, func(t *testing.T) {
			got := removeDuplicates(dc.records)
			want := dc.want
			assertEqual(t, got, want)
		})
	}

}

func assertEqual(t testing.TB, got, want []read.Record) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

var ExportCases = []struct {
	name    string
	records []read.Record
	want    map[string][]read.Record
}{}

//
//var sortCases = []struct {
//	name    string
//	records []read.Record
//	want    []read.Record
//}{
//	{
//		name: "out of order records",
//		records: []read.Record{
//			{
//				FirstName: "Thad",
//				LastName:  "Feest",
//				Email:     "Thad.Feest@cleve.name",
//				Address:   "77881 Schaefer Loaf",
//				Created:   "July 13, 2021",
//				Balance:   "$1,950.71",
//			},
//			{
//				FirstName: "Sabrina",
//				LastName:  "Kuphal",
//				Email:     "whiterabbit70@gmail.com",
//				Address:   "03491 Howard Vista",
//				Created:   "August 29, 2018",
//				Balance:   "$6,996.45",
//			},
//			{
//				FirstName: "Justice",
//				LastName:  "Keebler",
//				Email:     "magentarabbit07@gmail.com",
//				Address:   "173 Hyatt Crossroad",
//				Created:   "April 12, 2014",
//				Balance:   "$1,359.61"},
//		},
//		want: []read.Record{
//			{
//				FirstName: "Justice",
//				LastName:  "Keebler",
//				Email:     "magentarabbit07@gmail.com",
//				Address:   "173 Hyatt Crossroad",
//				Created:   "April 12, 2014",
//				Balance:   "$1,359.61",
//			},
//			{
//				FirstName: "Sabrina",
//				LastName:  "Kuphal",
//				Email:     "whiterabbit70@gmail.com",
//				Address:   "03491 Howard Vista",
//				Created:   "August 29, 2018",
//				Balance:   "$6,996.45",
//			},
//			{
//				FirstName: "Thad",
//				LastName:  "Feest",
//				Email:     "Thad.Feest@cleve.name",
//				Address:   "77881 Schaefer Loaf",
//				Created:   "July 13, 2021",
//				Balance:   "$1,950.71",
//			},
//		},
//	},
//}

var duplicatesCases = []struct {
	name    string
	records []read.Record
	want    []read.Record
}{
	{
		name: "multiple entries with one duplicate",
		records: []read.Record{
			{
				FirstName: "Thad",
				LastName:  "Feest",
				Email:     "Thad.Feest@cleve.name",
				Address:   "77881 Schaefer Loaf",
				Created:   "July 13, 2021",
				Balance:   "$1,950.71",
			},
			{
				FirstName: "Thad",
				LastName:  "Feest",
				Email:     "Thad.Feest@cleve.name",
				Address:   "77881 Schaefer Loaf",
				Created:   "July 13, 2021",
				Balance:   "$1,950.71",
			},
			{
				FirstName: "Justice",
				LastName:  "Keebler",
				Email:     "magentarabbit07@gmail.com",
				Address:   "173 Hyatt Crossroad",
				Created:   "April 12, 2014",
				Balance:   "$1,359.61"},
			{
				FirstName: "Sabrina",
				LastName:  "Kuphal",
				Email:     "whiterabbit70@gmail.com",
				Address:   "03491 Howard Vista",
				Created:   "August 29, 2018",
				Balance:   "$6,996.45"},
		},
		want: []read.Record{
			{
				FirstName: "Thad",
				LastName:  "Feest",
				Email:     "Thad.Feest@cleve.name",
				Address:   "77881 Schaefer Loaf",
				Created:   "July 13, 2021",
				Balance:   "$1,950.71",
			},
			{
				FirstName: "Justice",
				LastName:  "Keebler",
				Email:     "magentarabbit07@gmail.com",
				Address:   "173 Hyatt Crossroad",
				Created:   "April 12, 2014",
				Balance:   "$1,359.61"},
			{
				FirstName: "Sabrina",
				LastName:  "Kuphal",
				Email:     "whiterabbit70@gmail.com",
				Address:   "03491 Howard Vista",
				Created:   "August 29, 2018",
				Balance:   "$6,996.45"},
		},
	},
	{
		name: "all duplicate entries",
		records: []read.Record{
			{
				FirstName: "Thad",
				LastName:  "Feest",
				Email:     "Thad.Feest@cleve.name",
				Address:   "77881 Schaefer Loaf",
				Created:   "July 13, 2021",
				Balance:   "$1,950.71",
			},
			{
				FirstName: "Thad",
				LastName:  "Feest",
				Email:     "Thad.Feest@cleve.name",
				Address:   "77881 Schaefer Loaf",
				Created:   "July 13, 2021",
				Balance:   "$1,950.71",
			},
		},
		want: []read.Record{
			{
				FirstName: "Thad",
				LastName:  "Feest",
				Email:     "Thad.Feest@cleve.name",
				Address:   "77881 Schaefer Loaf",
				Created:   "July 13, 2021",
				Balance:   "$1,950.71",
			},
		},
	},
}
