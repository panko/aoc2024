package main

import (
	"fmt"
	"panko/aoc2024/pkg/input"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/stat/combin"
)

func main() {
	fmt.Println("Part 1: ", part1("input.txt"))
	fmt.Println("Part 2: ", part2("input.txt"))
}

func part1(filename string) int {
	res := 0
	equations := input.ReadLines(filename)
	for _, l := range equations {
		splitted := strings.Split(l, ":")
		line_res, _ := strconv.Atoi(splitted[0])
		splitted_parts := strings.Split(splitted[1], " ")
		splitted_parts = splitted_parts[1:]
		parts := make([]int, len(splitted_parts))
		for i, n := range splitted_parts {
			n, _ := strconv.Atoi(n)
			parts[i] = n
		}
		fmt.Println(parts)
		fmt.Println("Generate list:")
		n := 3
		k := 2
		list := combin.Permutations(n, k)
		for i, v := range list {
			fmt.Println(i, v)
		}

		fmt.Println(line_res)
	}

	return res
}

func part2(filename string) int {
	return 0
}
