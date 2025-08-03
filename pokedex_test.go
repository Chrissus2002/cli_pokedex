package main

import(
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "   hello world   ",
			expected: []string{"hello", "world"},
		},
		{
			input: " RAICHU CHarizard squirtle",
			expected: []string{"raichu", "charizard", "squirtle"},
		},
		{
			input: "I am Jotaro Kujoh  ",
			expected: []string{"i", "am", "jotaro", "kujoh"},
		},
	}
	for _, c := range cases{
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected){
			t.Errorf("Expected len %d, got %d", len(c.expected), len(actual))
		}
		for i,_ := range actual{
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord{
				t.Errorf("Words at index %d, word: %s, expected: %s", i, word, expectedWord)
			}
		}
	}
}