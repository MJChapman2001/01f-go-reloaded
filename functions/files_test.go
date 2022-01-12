package functions

import (
	"fmt"
	"testing"
)

func TestCommands(t *testing.T) {
	var tests = []struct {
		a string
		want string
	} {
		{"1E (hex)", "30"},
		{"10 (bin)", "2"},
		{"hello (up)", "HELLO"},
		{"Hello how are you? (up, 3)", "Hello HOW ARE YOU?"},
		{"HELLO (low)", "hello"},
		{"HELLO HOW ARE YOU? (low, 2)", "HELLO HOW are you?"},
		{"hello (cap)", "Hello"},
		{"hello how ARE yOu? (cap, 3)", "hello How Are You?"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := Commands(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestPunctuation(t *testing.T) {
	var tests = []struct {
		a string
		want string
	} {
		{"stop .", "stop."},
		{"stop   .", "stop."},
		{"pause ,", "pause,"},
		{"pause    ,", "pause,"},
		{"test!", "test!"},
		{"test   !", "test!"},
		{"why ?", "why?"},
		{"why   ?", "why?"},
		{"end  ;", "end;"},
		{"end     ;", "end;"},
		{"list:", "list:"},
		{"list  :", "list:"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := Punctuation(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestLetterA(t *testing.T) {
	var tests = []struct {
		a string
		want string
	} {
		{"a apple", "an apple"},
		{"a example", "an example"},
		{"a item", "an item"},
		{"a object", "an object"},
		{"a unboxing", "an unboxing"},
		{"a test", "a test"},
		{"A apple", "An apple"},
		{"A example", "An example"},
		{"A item", "An item"},
		{"A object", "An object"},
		{"A unboxing", "An unboxing"},
		{"A test", "A test"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := LetterA(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestQuote(t *testing.T) {
	var tests = []struct {
		a string
		want string
	} {
		{"This is ' a quote'", "This is 'a quote'"},
		{"This is '   a quote '", "This is 'a quote'"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := Quote(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}