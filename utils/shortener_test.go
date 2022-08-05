package utils

import (
	"regexp"
	"testing"
)

func TestMakeString(t *testing.T) {

	tests := []struct {
		length     int
		charset    string
		wantlength int
		wantchars  bool
	}{
		{length: 5, wantlength: 5, wantchars: true},
		{length: 8, wantlength: 8, wantchars: true},
		{length: 7, wantlength: 7, wantchars: true},
		{length: 4, wantlength: 4, wantchars: true},
	}

	r, _ := regexp.Compile("^[A-Za-z0-9_-]*$")

	for _, tc := range tests {
		got := MakeString(tc.length)
		match := r.MatchString(got)
		if !match {
			t.Fatalf("expected: %v, got: %v", tc.wantchars, match)
		}
	}

}
