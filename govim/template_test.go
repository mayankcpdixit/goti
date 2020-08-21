package main

import (
	"testing"
)

func MyFunc_Test(t *testing.T) {
	tcs := []struct {
		Input  int
		Output int
	}{{
		Input:  1,
		Output: 1,
	}}
	for _, tc := range tcs {
		if MyFunc(tc.Input) != tc.Output {
			t.Errorf("You messed up")
		}
	}
}
