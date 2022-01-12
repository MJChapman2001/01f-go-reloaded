package functions

import (
	"fmt"
	"testing"
)

func TestUp(t *testing.T) {
	var tests = []struct {
		a string
		want string
	} {
		{"hello", "HELLO"},
		{"How", "HOW"},
		{"ARE", "ARE"},
		{"yoU?", "YOU?"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := Up(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestLow(t *testing.T) {
	var tests = []struct {
		a string
		want string
	} {
		{"hello", "hello"},
		{"How", "how"},
		{"ARE", "are"},
		{"yoU?", "you?"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := Low(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestCap(t *testing.T) {
	var tests = []struct {
		a string
		want string
	} {
		{"hello", "Hello"},
		{"How", "How"},
		{"ARE", "Are"},
		{"yoU?", "You?"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := Cap(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}