package mutil

import (
	"testing"
)

func Abs_Test(t *testing.T) {
	tcs := []struct {
		Input  int
		Output int
	}{{
		Input:  1,
		Output: 1,
	}}
	for _, tc := range tcs {
		if Abs(tc.Input) != tc.Output {
			t.Errorf("You messed up")
		}
	}
}
