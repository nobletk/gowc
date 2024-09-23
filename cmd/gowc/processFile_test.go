package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProcessFile(t *testing.T) {
	expected := TotalCount{
		BytesTotal: 342190,
		LinesTotal: 7145,
		WordsTotal: 58164,
		CharsTotal: 339292,
	}
	fPath := "./testdata/test.txt"
	actual, err := GetCount(fPath)
	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestCountChars(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected int
	}{
		{
			name:     "Single english character",
			input:    []byte("z"),
			expected: 1,
		},
		{
			name:     "Single german character",
			input:    []byte("ë"),
			expected: 1,
		},
		{
			name:     "Two arabic characters",
			input:    []byte("أَ"),
			expected: 2,
		},
		{
			name:     "Multi-byte character split",
			input:    []byte("Hello, world! 界"),
			expected: 15,
		},
		{
			name:     "Invalid UTF-8 characters",
			input:    []byte("Invalid UTF-8: \xFF\xFF"),
			expected: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var count int

			count, _ = countChars(tt.input, count)
			assert.Equal(t, tt.expected, count)
		})
	}
}
