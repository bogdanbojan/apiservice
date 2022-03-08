package main

import (
	"apiserv/records"
	"reflect"
	"testing"
)

func TestTransformRecords(t *testing.T) {
	rt := records.NewFileDecoder()
	for _, rc := range recordsCases {
		got, _ := rt.TransformRecords(rc.records)
		assertEqual(t, got, rc.want)
	}
}
func assertEqual(t testing.TB, got, want []records.ExportRecords) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

var recordsCases = []struct {
	name    string
	records []records.Record
	want    []records.ExportRecords
}{
	{
		name: "multiple entries with one duplicate",
		records: []records.Record{
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
		want: []records.ExportRecords{
			{
				"J",
				[]records.Record{{
					FirstName: "Justice",
					LastName:  "Keebler",
					Email:     "magentarabbit07@gmail.com",
					Address:   "173 Hyatt Crossroad",
					Created:   "April 12, 2014",
					Balance:   "$1,359.61",
				},
				},
				1,
			},
			{
				"S",
				[]records.Record{{
					FirstName: "Sabrina",
					LastName:  "Kuphal",
					Email:     "whiterabbit70@gmail.com",
					Address:   "03491 Howard Vista",
					Created:   "August 29, 2018",
					Balance:   "$6,996.45",
				},
				},
				1,
			},
			{
				"T",
				[]records.Record{{
					FirstName: "Thad",
					LastName:  "Feest",
					Email:     "Thad.Feest@cleve.name",
					Address:   "77881 Schaefer Loaf",
					Created:   "July 13, 2021",
					Balance:   "$1,950.71",
				},
				},
				1,
			},
		},
	},
	{
		name: "entries that are only duplicates",
		records: []records.Record{
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
				FirstName: "Thad",
				LastName:  "Feest",
				Email:     "Thad.Feest@cleve.name",
				Address:   "77881 Schaefer Loaf",
				Created:   "July 13, 2021",
				Balance:   "$1,950.71",
			},
		},
		want: []records.ExportRecords{
			{
				"T",
				[]records.Record{{
					FirstName: "Thad",
					LastName:  "Feest",
					Email:     "Thad.Feest@cleve.name",
					Address:   "77881 Schaefer Loaf",
					Created:   "July 13, 2021",
					Balance:   "$1,950.71",
				},
				},
				1,
			},
		},
	},
	{
		name:    "empty entry",
		records: []records.Record{},
		want:    nil,
	},
}
