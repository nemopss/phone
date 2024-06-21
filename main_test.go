package main

import "testing"

func TestNormalize(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{"1234567890", "1234567890"},
		{"123123 456 7891", "1231234567891"},
		{"123(123) 456 7892", "1231234567892"},
		{"123(123) 456-7893", "1231234567893"},
		{"123123-456-7894", "1231234567894"},
		{"123123-456-7890", "1231234567890"},
		{"1231234567892", "1231234567892"},
		{"123(123)456-7892", "1231234567892"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actual := normalize(tc.input)
			if actual != tc.want {
				t.Errorf("got %s; want %s", actual, tc.want)
			}
		})
	}
}
