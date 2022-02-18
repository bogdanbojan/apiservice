package transform

import (
	"apiserv/read"
	"reflect"
	"testing"
)

// TODO: assert errors in the rest of the tests.
// TODO: group assertEqual under an interface that holds the types of data structures used in the transform step
// TODO: minimize duplicate structs used for tests. DRY.

func TestProcessExport(t *testing.T) {}

//func TestSortRecords(t *testing.T) {
//	for _, dc := range sortCases {
//		t.Run(dc.name, func(t *testing.T) {
//			got := sortRecords(dc.records)
//			want := dc.want
//			assertDuplicate(t, got, want)
//		})
//	}
//}
func TestFormatExport(t *testing.T) {
	for _, fc := range formatCases {
		t.Run(fc.name, func(t *testing.T) {
			got := formatExport(fc.mappedExports)
			want := fc.want
			assertFormatting(t, got, want)
		})
	}
}

func assertFormatting(t testing.TB, got, want []ExportRecords) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

var formatCases = []struct {
	name          string
	mappedExports map[string][]read.Record
	want          []ExportRecords
}{
	{
		name: "multiple entries",
		mappedExports: map[string][]read.Record{
			"J": {{FirstName: "Justice"}},
			"S": {{FirstName: "Sabrina"}},
			"T": {{FirstName: "Thad"}, {FirstName: "Tyson"}},
		},
		want: []ExportRecords{
			{
				"J",
				[]read.Record{{FirstName: "Justice"}},
				1,
			},
			{
				"S",
				[]read.Record{{FirstName: "Sabrina"}},
				1,
			},
			{
				"T",
				[]read.Record{{FirstName: "Thad"}, {FirstName: "Tyson"}},
				2,
			},
		},
	},
}

func TestMapExport(t *testing.T) {
	for _, mc := range mapCases {
		t.Run(mc.name, func(t *testing.T) {
			got := mapExport(mc.records)
			want := mc.want
			assertMapping(t, got, want)
		})
	}
}

func assertMapping(t testing.TB, got, want map[string][]read.Record) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

var mapCases = []struct {
	name    string
	records []read.Record
	want    map[string][]read.Record
}{
	{
		name: "multiple entries",
		records: []read.Record{
			{
				FirstName: "Thad",
			},
			{
				FirstName: "Tyson",
			},
			{
				FirstName: "Justice",
			},
			{
				FirstName: "Sabrina",
			},
		},
		want: map[string][]read.Record{
			"J": {{FirstName: "Justice"}},
			"S": {{FirstName: "Sabrina"}},
			"T": {{FirstName: "Thad"}, {FirstName: "Tyson"}},
		},
	}, {
		name: "all entries with the same letter",
		records: []read.Record{
			{
				FirstName: "Thad",
			},
			{
				FirstName: "Tyson",
			},
			{
				FirstName: "Thames",
			},
			{
				FirstName: "Thar",
			},
		},
		want: map[string][]read.Record{
			"T": {{FirstName: "Thad"}, {FirstName: "Tyson"}, {FirstName: "Thames"}, {FirstName: "Thar"}},
		},
	},
}

func TestRemoveDuplicates(t *testing.T) {
	for _, dc := range duplicatesCases {
		t.Run(dc.name, func(t *testing.T) {
			got := removeDuplicates(dc.records)
			want := dc.want
			assertDuplicate(t, got, want)
		})
	}

}

func assertDuplicate(t testing.TB, got, want []read.Record) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

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
