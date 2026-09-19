package main

import "testing"

func TestCapitalizeTitle(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"capiTalIze tHe titLe", "Capitalize The Title"},
		{"a bb ccc dddd", "a bb Ccc Dddd"},
	}

	for _, tc := range cases {
		t.Run(tc.in, func(t *testing.T) {
			if got := capitalizeTitle(tc.in); got != tc.want {
				t.Fatalf("capitalizeTitle(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
