package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func scanFile(textfile string) {
	f, err := os.OpenFile(textfile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer f.Close()
	res := 0
	do := true
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		szemet := sc.Text()
		// fmt.Println(szemet)

		r, _ := regexp.Compile(`mul\(([0-9][0-9]?[0-9]?),([0-9][0-9]?[0-9]?)\)|do\(\)|don\'t\(\)`)
		matches := r.FindAllStringSubmatch(szemet, -1)
		for _, v := range matches {
			if v[0] == "do()" {
				do = true
			} else if v[0] == "don't()" {
				do = false
			} else {
				if do {

					o, _ := strconv.Atoi(v[1])
					t, _ := strconv.Atoi(v[2])
					res += o * t
				}
			}
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return
	}
	fmt.Println(res)
}

func main() {
	scanFile("har.txt")
}
