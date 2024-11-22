package tools

import (
	"testing"
)

func Test_ConvertNonNegativeFloatToCents(t *testing.T) {
	for _, tCase := range []struct {
		name      string
		in        string
		out       int64
		haveError bool
	}{
		{
			name:      "1",
			in:        "100",
			out:       int64(10000),
			haveError: false,
		},
		{
			name:      "2",
			in:        "-100",
			haveError: true,
		},
		{
			name:      "3",
			in:        "100.1",
			out:       10010,
			haveError: false,
		},
		{
			name:      "4",
			in:        "100.10",
			out:       10010,
			haveError: false,
		},
		{
			name:      "5",
			in:        "100.01",
			out:       10001,
			haveError: false,
		},
		{
			name:      "6",
			in:        "100.00",
			out:       10000,
			haveError: false,
		},
		{
			name:      "7",
			in:        "-100.00",
			out:       10000,
			haveError: true,
		},
		{
			name:      "8",
			in:        "9223372036854775807",
			out:       10000,
			haveError: true,
		},
		{
			name:      "8",
			in:        "92233720368547758.07",
			out:       9223372036854775807,
			haveError: false,
		},
		{
			name:      "9",
			in:        "92233720368547758.08",
			out:       0,
			haveError: true,
		},
		{
			name:      "10",
			in:        "str",
			out:       0,
			haveError: true,
		},
	} {
		t.Run(tCase.name, func(t *testing.T) {
			actualOut, actualErr := ConvertNonNegativeFloatToCents(tCase.in)
			if tCase.haveError == (actualErr == nil) {
				t.Errorf("invalid error")
				return

			}
			if !tCase.haveError {
				if tCase.out != actualOut {
					t.Errorf("invalid out")
					return
				}
			}
		})
	}
}

func Test_CentsPrettyPrinted(t *testing.T) {
	for _, tCase := range []struct {
		name string
		in   int64
		out  string
	}{
		{
			name: "1",
			in:   100,
			out:  "1.00",
		},
		{
			name: "2",
			in:   101,
			out:  "1.01",
		},
		{
			name: "3",
			in:   0,
			out:  "0.00",
		},
		{
			name: "4",
			in:   1,
			out:  "0.01",
		},
		{
			name: "5",
			in:   11,
			out:  "0.11",
		},
		{
			name: "6",
			in:   9223372036854775807,
			out:  "92233720368547758.07",
		},
	} {
		t.Run(tCase.name, func(t *testing.T) {
			actualOut := CentsPrettyPrinted(tCase.in)
			if tCase.out != actualOut {
				t.Errorf("invalid out: %s", actualOut)
				return
			}
		})
	}
}
