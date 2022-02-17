package transform

import (
	"apiserv/read"
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	t.Run("testing multiple entries", func(t *testing.T) {
		records := []read.Record{
			{
				FirstName: "Thad",
				LastName:  "Feest",
				Email:     "Thad.Feest@cleve.name",
				Address:   "77881 Schaefer Loaf",
				Created:   "July 13, 2021",
				Balance:   "$1,950.71",
			},
			{FirstName: "Thad",
				LastName: "Feest",
				Email:    "Thad.Feest@cleve.name",
				Address:  "77881 Schaefer Loaf",
				Created:  "July 13, 2021",
				Balance:  "$1,950.71",
			},
			{
				FirstName: "Thad",
				LastName:  "Adam",
				Email:     "Thad.Feest@cleve.name",
				Address:   "77881 Schaefer Loaf",
				Created:   "July 13, 2021",
				Balance:   "$1,950.71"},
			{
				FirstName: "Sabrina",
				LastName:  "Kuphal",
				Email:     "whiterabbit70@gmail.com",
				Address:   "03491 Howard Vista",
				Created:   "August 29, 2018",
				Balance:   "$6,996.45"},
		}
		got := removeDuplicates(records)
		want := []read.Record{
			{FirstName: "Thad",
				LastName: "Feest",
				Email:    "Thad.Feest@cleve.name",
				Address:  "77881 Schaefer Loaf",
				Created:  "July 13, 2021",
				Balance:  "$1,950.71",
			},
			{
				FirstName: "Thad",
				LastName:  "Adam",
				Email:     "Thad.Feest@cleve.name",
				Address:   "77881 Schaefer Loaf",
				Created:   "July 13, 2021",
				Balance:   "$1,950.71"},
			{
				FirstName: "Sabrina",
				LastName:  "Kuphal",
				Email:     "whiterabbit70@gmail.com",
				Address:   "03491 Howard Vista",
				Created:   "August 29, 2018",
				Balance:   "$6,996.45"},
		}
		assertDuplicates(t, got, want)
	})
}

func assertDuplicates(t testing.TB, got, want []read.Record) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}
