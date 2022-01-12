package functions

import (
	"fmt"
	"testing"
)

func TestPower(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	} {
		{2, 2, 4},
		{4, 3, 64},
		{16, 1, 16},
		{10, 0, 1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := Power(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestHex(t *testing.T) {
	var tests = []struct {
		a string
		want string
	} {
		{"1E", "30"},
		{"4F", "79"},
		{"A2", "162"},
		{"9C", "156"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := Hex(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestBin(t *testing.T) {
	var tests = []struct {
		a string
		want string
	} {
		{"1001", "9"},
		{"101010", "42"},
		{"1101", "13"},
		{"11111", "31"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := Bin(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}