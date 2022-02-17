package transform

import (
	"apiserv/read"
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	for _, dc := range duplicatesCases {
		t.Run(dc.title, func(t *testing.T) {
			got := removeDuplicates(dc.records)
			want := dc.want
			assertDuplicates(t, got, want)
		})
	}

}

func assertDuplicates(t testing.TB, got, want []read.Record) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

var duplicatesCases = []struct {
	title   string
	records []read.Record
	want    []read.Record
}{
	{
		title: "multiple entries with one duplicate",
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
		title: "all duplicate entries",
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
