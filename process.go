package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func ProcessFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	tabSize, err := GuessTabSize(file)
	if err != nil {
		return nil, err
	} else if tabSize == 0 { // file doesn't have space indentation, no need to rewrite
		return []string{}, nil
	}

	newLines, err := RebuildLines(file, tabSize)
	if err != nil {
		return nil, err
	}

	return newLines, nil
}

func GuessTabSize(file *os.File) (int, error) {
	defer file.Seek(0, io.SeekStart)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tabSize := len(line) - len(strings.TrimLeft(line, " "))
		if tabSize >= 2 {
			return tabSize, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return 0, nil // no leading spaces found
}

func RebuildLines(file *os.File, fileTabSize int) ([]string, error) {
	defer file.Seek(0, io.SeekStart)

	newLines := make([]string, 0, 30) // arbitrary buffer size
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineTabSize := len(line) - len(strings.TrimLeft(line, " "))
		if lineTabSize >= 2 {
			numTabs := (lineTabSize - (lineTabSize % fileTabSize)) / fileTabSize
			newLine := strings.Repeat("\t", numTabs) + line[numTabs*fileTabSize:]
			newLines = append(newLines, newLine)
		} else {
			newLines = append(newLines, line) // nothing to do
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return newLines, nil
}
