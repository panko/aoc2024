package main

import (
	"fmt"
	"panko/aoc2024/pkg/input"
)

func main() {
	fmt.Println("Part 1: ", part1("input.txt"))
	fmt.Println("Part 2: ", part2("input.txt"))
}

func part1(name string) int {
	lines := input.ReadLines(name)
	xmas := "XMAS"
	res := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		for j := 0; j < len(line); j++ {
			found1 := 0
			found2 := 0
			found3 := 0
			found4 := 0
			found5 := 0
			found6 := 0
			found7 := 0
			found8 := 0
			for x := 0; x < len(xmas); x++ {
				if j+x < len(line) && line[j+x] == xmas[x] {
					found1 += 1
				}
				if found1 == 4 {
					res += 1
				}
				if j-x >= 0 && line[j-x] == xmas[x] {
					found2 += 1
				}
				if found2 == 4 {
					res += 1
				}
				if i+x < len(lines) && lines[i+x][j] == xmas[x] {
					found3 += 1
				}
				if found3 == 4 {
					res += 1
				}
				if i-x >= 0 && lines[i-x][j] == xmas[x] {
					found4 += 1
				}
				if found4 == 4 {
					res += 1
				}
				if i+x < len(lines) && j+x < len(line) && lines[i+x][j+x] == xmas[x] {
					found5 += 1
				}
				if found5 == 4 {
					res += 1
				}
				if i+x < len(lines) && j-x >= 0 && lines[i+x][j-x] == xmas[x] {
					found6 += 1
				}
				if found6 == 4 {
					res += 1
				}
				if i-x >= 0 && j-x >= 0 && lines[i-x][j-x] == xmas[x] {
					found7 += 1
				}
				if found7 == 4 {
					res += 1
				}
				if i-x >= 0 && j+x < len(line) && lines[i-x][j+x] == xmas[x] {
					found8 += 1
				}
				if found8 == 4 {
					res += 1
				}
			}
		}
	}
	return res
}

func part2(name string) int {
	lines := input.ReadLines(name)
	mas := "MAS"
	res := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		for j := 0; j < len(line); j++ {
			founda := false
			foundb := false
			found1 := 0
			found2 := 0
			found3 := 0
			found4 := 0
			if i < 1 || i > len(lines)-2 || j < 1 || j > len(line)-2 {
				continue
			}
			for x := 0; x < len(mas); x++ {
				if lines[i+x-1][j+x-1] == mas[x] {
					found1 += 1
				}
				if lines[i-x+1][j-x+1] == mas[x] {
					found2 += 1
				}
				if found1 == 3 || found2 == 3 {
					founda = true
				}
				if lines[i-x+1][j+x-1] == mas[x] {
					found3 += 1
				}
				if lines[i+x-1][j-x+1] == mas[x] {
					found4 += 1
				}
				if found3 == 3 || found4 == 3 {
					foundb = true
				}
				if founda && foundb {
					res += 1
				}
			}
		}
	}
	return res
}
