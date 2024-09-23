package main

import (
	"bufio"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

type TotalCount struct {
	BytesTotal int
	WordsTotal int
	LinesTotal int
	CharsTotal int
}

func GetCount(filePath string) (TotalCount, error) {
	var reader *bufio.Reader
	var byteCount, wordCount, lineCount, charCount int

	if filePath == "" {
		reader = bufio.NewReader(os.Stdin)
	} else {
		file, err := os.Open(filePath)
		if err != nil {
			return TotalCount{}, err
		}
		defer file.Close()
		reader = bufio.NewReader(file)
	}

	var remainingRunes []byte
	const bufferSize = 4096
	inWord := false

	for {
		buf := make([]byte, bufferSize)
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return TotalCount{}, err
		}

		byteCount += len(buf[:n])
		lineCount += countLines(buf[:n])
		wordCount, inWord = countWords(buf[:n], inWord, wordCount)

		chunk := append(remainingRunes, buf[:n]...)
		charCount, remainingRunes = countChars(chunk, charCount)
	}

	return TotalCount{
		BytesTotal: byteCount,
		LinesTotal: lineCount,
		WordsTotal: wordCount,
		CharsTotal: charCount,
	}, nil
}

func countLines(chunk []byte) int {
	lineCount := 0
	for _, b := range chunk {
		if b == '\n' {
			lineCount++
		}
	}
	return lineCount
}

func countWords(chunk []byte, inWord bool, wordCount int) (int, bool) {
	for _, b := range chunk {
		if unicode.IsSpace(rune(b)) {
			if inWord {
				wordCount++
				inWord = false
			}
		} else {
			inWord = true
		}
	}

	return wordCount, inWord
}

func countChars(buf []byte, charCount int) (int, []byte) {
	for len(buf) > 0 {
		r, sz := utf8.DecodeRune(buf)
		if r == utf8.RuneError && sz == 1 {
			// Handle invalid UTF-8
			return charCount, buf
		}
		charCount++
		buf = buf[sz:]
	}

	return charCount, nil
}
