//testing code is in same package it uses

package main

import (
	"fmt"
	"testing"
)

// testing int min
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//Typically, the code we’re testing would be in a source file named something like intutils.go
// and the test file for it would then be named intutils_test.go.

// test function
// error reporst & continues
// fatal aborts
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// table driven test for inputs and outputs
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	//subset for each table entry
	for _, tt := range tests {

		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// Benchmark tests typically go in _test.go files and are named beginning with Benchmark.
// The testing runner executes each benchmark function several times, increasing b.N on each run until it collects a precise measurement.
func BenchmarkIntMin(b *testing.B) {
	//benchmarked func runs  a loop b.N times
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
