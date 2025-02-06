package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

// Add this function definition at the beginning of your file
func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestCleanEmptyLines(t *testing.T) {
	cases := []struct {
		input    []string
		expected []string
	}{
		{[]string{"hello", "", "world"}, []string{"hello", "world"}},
		{[]string{"", "", ""}, []string{}},
		{[]string{"a", "b", "c"}, []string{"a", "b", "c"}},
		{[]string{"", "a", "", "b", "", "c", ""}, []string{"a", "b", "c"}},
	}

	for _, c := range cases {
		result := cleanEmptyLines(c.input)
		if !equalSlices(result, c.expected) {
			t.Errorf("squeeze(%v) = %v; want %v", c.input, result, c.expected)
		}
	}
}

func TestBuildAsciiMap(t *testing.T) {
	// Test case 1: Basic mapping
	input := []string{
		"line1", "line2", "line3", "line4", "line5", "line6", "line7", "line8", // for space character
		"lineA1", "lineA2", "lineA3", "lineA4", "lineA5", "lineA6", "lineA7", "lineA8", // for '!' character
	}

	expected := map[rune][8]string{
		' ': {"line1", "line2", "line3", "line4", "line5", "line6", "line7", "line8"},
		'!': {"lineA1", "lineA2", "lineA3", "lineA4", "lineA5", "lineA6", "lineA7", "lineA8"},
	}

	result := buildAsciiMap(input)

	// Compare results
	if len(result) != len(expected) {
		t.Errorf("Expected map length %d, got %d", len(expected), len(result))
	}

	for k, v := range expected {
		if resultVal, exists := result[k]; !exists {
			t.Errorf("Expected key %q to exist in result", k)
		} else if resultVal != v {
			t.Errorf("For key %q, expected %v, got %v", k, v, resultVal)
		}
	}
}

func TestRenderWord(t *testing.T) {
	// Setup test map
	asciiMap = map[rune][8]string{
		'H': {"H1", "H2", "H3", "H4", "H5", "H6", "H7", "H8"},
		'I': {"I1", "I2", "I3", "I4", "I5", "I6", "I7", "I8"},
	}

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Test the function
	renderWord("HI")

	// Restore stdout
	w.Close()
	os.Stdout = oldStdout

	// Read captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)
	result := buf.String()

	expected := "H1I1\nH2I2\nH3I3\nH4I4\nH5I5\nH6I6\nH7I7\nH8I8\n"
	if result != expected {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expected, result)
	}
}

func TestDisplayWords(t *testing.T) {
	// Setup test map
	asciiMap = map[rune][8]string{
		'H': {"H1", "H2", "H3", "H4", "H5", "H6", "H7", "H8"},
		'I': {"I1", "I2", "I3", "I4", "I5", "I6", "I7", "I8"},
	}

	cases := []struct {
		input    []string
		expected string
	}{
		{
			[]string{"HI", "HI"},
			"H1I1\nH2I2\nH3I3\nH4I4\nH5I5\nH6I6\nH7I7\nH8I8\n" +
				"H1I1\nH2I2\nH3I3\nH4I4\nH5I5\nH6I6\nH7I7\nH8I8\n",
		},
		{
			[]string{"HI", "", "HI"},
			"H1I1\nH2I2\nH3I3\nH4I4\nH5I5\nH6I6\nH7I7\nH8I8\n" +
				"\n" +
				"H1I1\nH2I2\nH3I3\nH4I4\nH5I5\nH6I6\nH7I7\nH8I8\n",
		},
	}

	for _, c := range cases {
		// Capture stdout
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		// Test the function
		displayWords(c.input)

		// Restore stdout
		w.Close()
		os.Stdout = oldStdout

		// Read captured output
		var buf bytes.Buffer
		io.Copy(&buf, r)
		result := buf.String()

		if result != c.expected {
			t.Errorf("printWords(%v)\nExpected:\n%s\nGot:\n%s",
				c.input, c.expected, result)
		}
	}
}
