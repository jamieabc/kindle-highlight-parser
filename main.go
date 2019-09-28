package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

const (
	inputFileName      = "My Clippings.txt"
	outputFileName     = "kindle.txt"
	highlightSeparator = "=========="
	specialChar        = "\ufeff"
	positionPrefix     = "- "
	sentenceStartIndex = 3
	newLineChars       = "\r\n"

	// "2018年5月3日星期四 下午2:59:59"
	timeFormat = "(20\\d\\d).*(\\d+).*(\\d+).* (.*)(\\d+):(\\d+):(\\d+)"
)

type highlight struct {
	sentence string
	date     string
}

type highlights map[string][]highlight

func main() {
	bytes, err := rawBytes(inputFileName)
	if nil != err {
		fmt.Printf("get highlight with error: %s\n", err)
		return
	}
	hs := parse(bytes)
	writeToOutputFile(hs)
}

func rawBytes(file string) ([]byte, error) {
	handle, err := os.Open(file)
	if nil != err {
		return nil, err
	}
	defer handle.Close()

	bytes, err := ioutil.ReadAll(handle)
	if nil != err {
		return nil, err
	}

	return bytes, nil
}

func parse(raw []byte) highlights {
	paragraphs := strings.Split(string(raw), highlightSeparator)
	hs := make(highlights)

	for _, p := range paragraphs {
		truncated := removeFirstEmptyLines(p)
		if 2 >= len(truncated) {
			continue
		}
		bookName, timeStr, sentence := parseParagraph(truncated)
		combineHighlight(hs, bookName, highlight{
			sentence: sentence,
			date:     timeStr,
		})
	}

	return hs
}

func removeFirstEmptyLines(line string) string {
	for i := 0; i < len(line); i = i + 2 {
		if newLineChars != line[i:i+2] {
			return line[i:]
		}
	}
	return ""
}

func parseParagraph(paragraph string) (bookName string, timeStr string, sentence string) {
	lines := strings.Split(paragraph, "\r\n")
	bookName = parseBook(lines[0])
	timeStr, err := parseTime(lines[1])
	if nil != err {
		fmt.Printf("parse time of %s with error: %s\n", lines[1], err)
		return
	}
	sentence = strings.Join(lines[sentenceStartIndex:], "\n")
	return
}

func parseBook(line string) string {
	return strings.ReplaceAll(line, specialChar, "")
}

func parseTime(line string) (string, error) {
	r, _ := regexp.Compile(timeFormat)
	matches := r.FindAllString(line, -1)
	if 1 != len(matches) {
		return "", errors.New("error parse time format")
	}
	return matches[0], nil
}

func combineHighlight(hs highlights, book string, h highlight) {
	if _, ok := hs[book]; !ok {
		hs[book] = make([]highlight, 0)
	}
	hs[book] = append(hs[book], h)
}

func writeToOutputFile(hs highlights) {
	f, err := os.Create(outputFileName)
	if nil != err {
		fmt.Printf("create file %s with error: %s\n", outputFileName, err)
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	bookNumber := 1
	for bookName, sentences := range hs {
		w.WriteString(fmt.Sprintf("%d. %s\n\n", bookNumber, bookName))
		for _, s := range sentences {
			if emptyLine(s.sentence) {
				continue
			}
			str := fmt.Sprintf("\t%s%s\n", positionPrefix, s.sentence)
			w.WriteString(str)
		}
		bookNumber++
	}
	w.Flush()
}

func emptyLine(line string) bool {
	truncated := strings.ReplaceAll(line, "\n", "")
	truncated = strings.ReplaceAll(truncated, " ", "")
	return 0 == len(truncated)
}
