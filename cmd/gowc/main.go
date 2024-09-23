package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
)

type Flags struct {
	ByteFlag bool
	WordFlag bool
	LineFlag bool
	CharFlag bool
}

func main() {
	var f Flags

	pflag.BoolVarP(&f.ByteFlag, "bytes", "c", false, "print the byte counts")
	pflag.BoolVarP(&f.WordFlag, "words", "w", false, "print the word counts")
	pflag.BoolVarP(&f.LineFlag, "lines", "l", false, "print the newline counts")
	pflag.BoolVarP(&f.CharFlag, "chars", "m", false, "print the character counts")

	pflag.Usage = func() {
		var buf bytes.Buffer

		buf.WriteString("Usage:\n")
		buf.WriteString("  gowc [OPTION]... [FILE]...\n")
		buf.WriteString("  cat [FILE]... | gowc [OPTION]...\n")
		buf.WriteString("Options:\n")

		fmt.Fprintf(os.Stderr, buf.String())
		pflag.PrintDefaults()
	}

	pflag.Parse()

	if len(pflag.Args()) > 1 {
		pflag.Usage()
		os.Exit(2)
	}

	fPath := pflag.CommandLine.Arg(0)
	totalCount, err := GetCount(fPath)
	if err != nil {
		fmt.Println("Error getting count: ", err)
		os.Exit(1)
	}

	fmt.Println("  " + printTotalCount(f, totalCount, fPath))
}

func printTotalCount(f Flags, totalCount TotalCount, fPath string) string {
	if !f.ByteFlag && !f.WordFlag && !f.LineFlag && !f.CharFlag {
		f.ByteFlag = true
		f.WordFlag = true
		f.LineFlag = true
	}

	var out []string
	if f.CharFlag {
		out = append(out, strconv.Itoa(totalCount.CharsTotal))
	}

	if f.LineFlag {
		out = append(out, strconv.Itoa(totalCount.LinesTotal))
	}

	if f.WordFlag {
		out = append(out, strconv.Itoa(totalCount.WordsTotal))
	}

	if f.ByteFlag {
		out = append(out, strconv.Itoa(totalCount.BytesTotal))
	}

	out = append(out, fPath)

	return strings.Join(out, "  ")
}
