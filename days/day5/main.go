package main

import (
	"fmt"
	"panko/aoc2024/pkg/input"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1: ", part1("input.txt"))
	fmt.Println("Part 2: ", part2("input.txt"))
}

func StoI(strs []string) []int {
	var ints []int
	for _, s := range strs {
		i, _ := strconv.Atoi(s)
		ints = append(ints, i)
	}
	return ints
}

func part1(name string) int {
	inp := input.ReadParagraphs(name)
	res := 0
	for _, line := range inp[1] {
		splitted := strings.Split(line, ",")
		update := StoI(splitted)
		breaking := false
		for i, u := range update {
			if i == 0 {
				continue
			}
			for _, r := range inp[0] {
				str_rules := strings.Split(r, "|")
				rules := StoI(str_rules)
				if slices.Contains(update, rules[0]) && slices.Contains(update, rules[1]) {
					for j := i; j >= 0; j-- {
						if u == rules[0] && rules[1] == update[j] {
							breaking = true
							break
						}
					}
				} else {
					continue
				}
				if breaking {
					break
				}
			}
			if breaking {
				break
			}
		}
		if breaking {
			continue
		}
		res += update[len(update)/2]

	}
	return res
}

func part2(name string) int {
	inp := input.ReadParagraphs(name)
	res := 0
	for _, line := range inp[1] {
		splitted := strings.Split(line, ",")
		update := StoI(splitted)
		breaking := false
		for i := 0; i < len(update); i++ {
			if i == 0 {
				continue
			}
			for r := 0; r < len(inp[0]); r++ {
				str_rules := strings.Split(inp[0][r], "|")
				rules := StoI(str_rules)
				if slices.Contains(update, rules[0]) && slices.Contains(update, rules[1]) {
					for j := i; j >= 0; j-- {
						if update[i] == rules[0] && rules[1] == update[j] {
							breaking = true
							update[i], update[j] = update[j], update[i]
							r = 0
							i = 0
						}
					}
				} else {
					continue
				}
			}
		}
		if breaking {
			res += update[len(update)/2]
		}
	}
	return res
}
