package cli

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hi  there how are you   ",
			expected: []string{"hi", "there", "how", "are", "you"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf(
				"The resulted slice has %d elements, but must have %d",
				len(actual),
				len(c.expected),
			)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf(
					"Expected word \"%s\", but got \"%s\"",
					expectedWord,
					word,
				)
			}
		}
	}
}
