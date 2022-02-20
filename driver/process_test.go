package driver

import (
	"apiserv/read"
	"apiserv/transform"
	"reflect"
	"testing"
)

func TestTransformRecords(t *testing.T) {
	rt := transform.NewFileDecoder()
	for _, rc := range recordsCases {
		got, _ := rt.TransformRecords(rc.records)
		assertEqual(t, got, rc.want)
	}
}
func assertEqual(t testing.TB, got, want []transform.ExportRecords) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

var recordsCases = []struct {
	name    string
	records []read.Record
	want    []transform.ExportRecords
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
		want: []transform.ExportRecords{
			{
				"J",
				[]read.Record{{
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
				[]read.Record{{
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
				[]read.Record{{
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
				FirstName: "Thad",
				LastName:  "Feest",
				Email:     "Thad.Feest@cleve.name",
				Address:   "77881 Schaefer Loaf",
				Created:   "July 13, 2021",
				Balance:   "$1,950.71",
			},
		},
		want: []transform.ExportRecords{
			{
				"T",
				[]read.Record{{
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
		records: []read.Record{},
		want:    nil,
	},
}
