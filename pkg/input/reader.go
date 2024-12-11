package input

import (
	"bufio"
	"os"
)

func ReadParagraphs(filename string) [][]string {
	lines := ReadLines(filename)

	res := make([][]string, 0)
	var parag []string

	for _, line := range lines {
		if line == "" {
			res = append(res, parag)
			parag = nil
		} else {
			parag = append(parag, line)
		}
	}
	res = append(res, parag)
	return res
}

func ReadLines(filename string) []string {
	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines []string

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}
	return lines
}
