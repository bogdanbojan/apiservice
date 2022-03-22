package records

import (
	"context"
	"reflect"
	"testing"
)

var ctx = context.Background()

func TestGetAdditionalRecords(t *testing.T) {
	records := recordsCases[0].records
	t.Run("nr of records is enough", func(t *testing.T) {
		got, _ := getAdditionalRecords(ctx, records, 3, "")
		want := records
		assertEqual(t, got, want)

	})
	t.Run("need to fetch additional records", func(t *testing.T) {
		_, err := getAdditionalRecords(ctx, records, 10, "")
		if err == nil {
			t.Errorf("expected an error")
		}
	})
}

func TestValidateRecords(t *testing.T) {
	records := recordsCases[0].records
	t.Run("nr of records is less than the desired user set nr", func(t *testing.T) {
		_, err := validateRecordsNr(ctx, records, 4, "")
		// TODO: Should i have a custom error so I can test it against assertError?
		if err == nil {
			t.Errorf("expected an error")
		}
	})
	t.Run("nr of records is the same as the desired user set nr", func(t *testing.T) {
		got, _ := validateRecordsNr(ctx, records, 3, "")
		want := []Record{
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
				Balance:   "$6,996.45",
			},
		}
		assertEqual(t, got, want)

	})
	t.Run("nr of records is bigger than the desired user set nr", func(t *testing.T) {
		got, _ := validateRecordsNr(ctx, records, 2, "")
		want := []Record{
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
		}
		assertEqual(t, got, want)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertEqual(t testing.TB, got, want []Record) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

var recordsCases = []struct {
	name    string
	records []Record
}{
	{
		name: "multiple entries",
		records: []Record{
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
		name: "one entry",
		records: []Record{
			{
				FirstName: "Abdiel",
				LastName:  "Jenkins",
				Email:     "Abdiel.Jenkins@tate.info",
				Address:   "0196 Edmond Falls",
				Created:   "May 7, 2016",
				Balance:   "$5,424.08",
			},
		},
	},
}
