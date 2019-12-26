package main

import (
	"fmt"
	"testing"
)

func TestAppendEven(t *testing.T) {
	tests := []struct {
		actual   string
		expected string
	}{
		{"パタトクカシーー", "パトカー"},
		{"atourofgo", "aorfo"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s to %s", test.actual, test.expected), func(t *testing.T) {
			if got := appendEven(test.actual); got != test.expected {
				t.Errorf("got %s, to %s", got, test.expected)
			}
		})
	}
}
