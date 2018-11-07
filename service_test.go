package main

import "testing"

func TestUppercase(t *testing.T) {
	tt := []struct {
		in             string
		out            string
		expectingError bool
	}{
		{
			"",
			"HELLO",
			true,
		},
		{
			"hello",
			"HELLO",
			false,
		},
		{
			"World",
			"WORLD",
			false,
		},
	}

	ss := stringService{}
	for _, tc := range tt {
		res, err := ss.Uppercase(tc.in)
		if err != nil {
			if !tc.expectingError {
				t.Fatal("expecting no errors")
			}
			continue
		}

		if err == nil && tc.expectingError {
			t.Fatal("expecting errors")
		}

		if res != tc.out {
			t.Errorf("expected %s, found %s", tc.out, res)
		}
	}
}

func TestCount(t *testing.T) {
	tt := []struct {
		in    string
		count int
	}{
		{
			"",
			0,
		},
		{
			"hello",
			5,
		},
		{
			"Hello World",
			11,
		},
	}

	ss := stringService{}
	for _, tc := range tt {
		res := ss.Count(tc.in)
		if res != tc.count {
			t.Errorf("expected count %d, received %d", tc.count, res)
		}
	}
}
