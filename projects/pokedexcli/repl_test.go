package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{"hello world", []string{"hello", "world"}},
		{"  hello   world  ", []string{"hello", "world"}},
		{"HELLO WORLD", []string{"hello", "world"}},
		{"hello 123", []string{"hello", "123"}},
		{"hello! world?", []string{"hello!", "world?"}},
		{"", []string{}},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf(
				"The lengthd are not equal: %v VS %v",
				len(actual),
				len(cs.expected),
			)
			continue
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]

			if actualWord != expectedWord {
				t.Errorf(
					"Expected '%s' but got '%s'",
					expectedWord,
					actualWord,
				)
			}
		}
	}
}
