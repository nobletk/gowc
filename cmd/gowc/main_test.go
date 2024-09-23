package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.T) {
	tests := []struct {
		name     string
		flags    Flags
		expected string
	}{
		{
			name: "Print all count",
			flags: Flags{
				ByteFlag: true,
				WordFlag: true,
				LineFlag: true,
				CharFlag: false,
			},
			expected: "7145  58164  342190  ./testdata/test.txt",
		},
		{
			name: "Print number of byte count",
			flags: Flags{
				ByteFlag: true,
				WordFlag: false,
				LineFlag: false,
				CharFlag: false,
			},
			expected: "342190  ./testdata/test.txt",
		},
		{
			name: "Print the word count",
			flags: Flags{
				ByteFlag: false,
				WordFlag: true,
				LineFlag: false,
				CharFlag: false,
			},
			expected: "58164  ./testdata/test.txt",
		},
		{
			name: "Print the newline count",
			flags: Flags{
				ByteFlag: false,
				WordFlag: false,
				LineFlag: true,
				CharFlag: false,
			},
			expected: "7145  ./testdata/test.txt",
		},
		{
			name: "Print the character count",
			flags: Flags{
				ByteFlag: false,
				WordFlag: false,
				LineFlag: false,
				CharFlag: true,
			},
			expected: "339292  ./testdata/test.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fPath := "./testdata/test.txt"

			totalCount, err := GetCount(fPath)
			require.NoError(t, err)

			actual := printTotalCount(tt.flags, totalCount, fPath)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
