package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var targetExtensions = []string{
	".html",
	".js",
	".json",
	".jsx",
	".svelte",
	".svg",
	".ts",
	".tsx",
	".vue",
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr,
			"ERROR: Please inform the target directory, like:\n"+
				"    ./tabify ~/foo/my-project")
		os.Exit(1)
	}

	if countScanned, countWritten, err := ProcessTargetDir(os.Args[1]); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	} else {
		fmt.Fprintf(os.Stdout,
			"Files scanned: %d; written: %d.\n",
			countScanned, countWritten)
	}
}

func ProcessTargetDir(targetDir string) (int, int, error) {
	countScanned := 0
	countWritten := 0

	err := filepath.Walk(targetDir,
		func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() && HasValidExtension(path) {
				newLines, err := ProcessFile(path)
				if err != nil {
					return err
				}
				countScanned++

				if len(newLines) > 0 {
					if err := ReSave(path, newLines); err != nil {
						return err
					}
					countWritten++
				}
			}
			return nil
		},
	)

	if err != nil {
		return 0, 0, err
	} else {
		return countScanned, countWritten, nil
	}
}

func HasValidExtension(path string) bool {
	fileExt := strings.ToLower(filepath.Ext(path))
	for _, ext := range targetExtensions {
		if ext == fileExt {
			return true
		}
	}
	return false
}

func ReSave(path string, lines []string) error {
	file, err := os.OpenFile(path, os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range lines {
		_, err := fmt.Fprintln(file, line)
		if err != nil {
			return err
		}
	}

	fmt.Fprintf(os.Stdout, "Written: %s (%d lines)\n", path, len(lines))
	return nil
}
