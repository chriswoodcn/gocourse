package test

import (
	"fmt"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, c int32
	}{
		{1, 2, 3},
		{5, 6, 11},
		{2, 2, 4},
		{math.MaxInt32, 1, math.MaxInt32},
	}
	for _, tt := range tests {
		if actual := Add(tt.a, tt.b); actual != tt.c {
			t.Errorf("test error v: %v \n", tt)
		}
	}
}
func TestAny(t *testing.T) {
	fmt.Println("test any func exec")
}
func TestAll(t *testing.T) {
	TestAdd(t)
	TestAny(t)
}
