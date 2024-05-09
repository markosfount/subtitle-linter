package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"
)

var regChar = regexp.MustCompile("(^.+: )")
var regPar = regexp.MustCompile("(\\(.+\\))")

func breakLongLines(filePath string) {
	//dat, err := os.ReadFile(filePath)
	//detector := chardet.NewTextDetector()
	//result, err := detector.DetectBest(dat)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}
	//fmt.Println(result)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	//dec := transform.NewReader(file, charmap.Windows1252.NewDecoder())

	scanner := bufio.NewScanner(file)
	var modifiedLines []string
	for scanner.Scan() {
		line := scanner.Text()
		//line2 := scanner.Bytes()
		//ls := string(line2)
		//fmt.Println(ls)
		isSub := regChar.MatchString(line)
		isTime := strings.Contains(line, "TIME")
		if isSub || isTime {
			modifiedLines = append(modifiedLines, "")
		}
		line = regChar.ReplaceAllString(line, "")
		line = regPar.ReplaceAllString(line, "")

		if length(line) > 25 { //&& regChar.MatchString(line) {
			words := strings.Fields(line)
			var currentLine string
			for _, word := range words {
				if length(currentLine)+length(word)+1 <= 25 {
					currentLine += word + " "
				} else {
					modifiedLines = append(modifiedLines, strings.TrimSpace(currentLine))
					currentLine = word + " "
				}
			}
			modifiedLines = append(modifiedLines, strings.TrimSpace(currentLine))
		} else {
			modifiedLines = append(modifiedLines, line)
		}
		if isTime {
			modifiedLines = append(modifiedLines, "")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	outputFile, err := os.Create("output3.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for _, line := range modifiedLines {
		fmt.Fprintln(writer, line)
	}
	writer.Flush()
}

func length(line string) int {
	return utf8.RuneCountInString(line)
}

func main() {
	//filePath := "/home/markos/Documents/11th Draft.txt" // Change this to your file path
	filePath := "11th Draft.txt" // Change this to your file path
	breakLongLines(filePath)
}
