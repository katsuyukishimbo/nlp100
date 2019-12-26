package main

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		actual   string
		expected string
	}{
		{"stressed", "desserts"},
		{"level", "level"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s to %s", test.actual, test.expected), func(t *testing.T) {
			if got := Reverse(test.actual); got != test.expected {
				t.Errorf("got %s, to %s", got, test.expected)
			}
		})
	}
}
