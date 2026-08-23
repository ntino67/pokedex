package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCleanInput(t *testing.T) {
	tests := map[string]struct {
		input string
		want []string
	}{
		"hello world": {input: "hello world", want: []string{"hello", "world"}},
		"caps": {input: "HELLO WORLD", want: []string{"hello", "world"}},
		"trailing whitespace": {input: "           hello world          ", want: []string{"hello", "world"}},
		"inside whitespace": {input: "hello             world", want: []string{"hello", "world"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := cleanInput(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
