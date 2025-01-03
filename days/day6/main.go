package main

import (
	"fmt"
	"panko/aoc2024/pkg/input"
	"slices"
	"strings"
)

func main() {
	fmt.Println("Part 1: ", part1("input.txt"))
	fmt.Println("Part 2: ", part2("input.txt"))
}

func isGuardOnMap(mymap []string) bool {
	guard := []rune{'v', '^', '<', '>'}
	for _, l := range mymap {
		for _, c := range l {
			if slices.Contains(guard, c) {
				return true
			}
		}
	}
	return false
}

func replaceAtIndex(str string, replacement rune, index int) string {
	return str[:index] + string(replacement) + str[index+1:]
}

func isObstacle(r rune) bool {
	return r == '#'
}

func guardWalk(mymap *[]string) {
	guard := []rune{'v', '^', '<', '>'}
	for i, l := range *mymap {
		for j, c := range l {
			if slices.Contains(guard, c) {
				if c == '^' {
					nextLineRunes := []rune((*mymap)[i-1])
					next := nextLineRunes[j]
					if isObstacle(next) {
						(*mymap)[i] = replaceAtIndex(l, '>', j)
					} else {
						(*mymap)[i-1] = replaceAtIndex((*mymap)[i-1], '^', j)
						(*mymap)[i] = replaceAtIndex((*mymap)[i], 'X', j)
					}
					return
				}
				if c == '>' {
					nextLineRunes := []rune((*mymap)[i])
					next := nextLineRunes[j+1]
					if isObstacle(next) {
						(*mymap)[i] = replaceAtIndex(l, 'v', j)
					} else {
						(*mymap)[i] = replaceAtIndex((*mymap)[i], '>', j+1)
						(*mymap)[i] = replaceAtIndex((*mymap)[i], 'X', j)
					}
					return
				}
				if c == 'v' {
					if i+1 == len(*mymap) {
						(*mymap)[i] = replaceAtIndex((*mymap)[i], 'X', j)
						return
					}
					nextLineRunes := []rune((*mymap)[i+1])
					next := nextLineRunes[j]
					if isObstacle(next) {
						(*mymap)[i] = replaceAtIndex(l, '<', j)
					} else {
						(*mymap)[i+1] = replaceAtIndex((*mymap)[i+1], 'v', j)
						(*mymap)[i] = replaceAtIndex((*mymap)[i], 'X', j)
					}
					return
				}
				if c == '<' {
					if j-1 < 0 {
						(*mymap)[i] = replaceAtIndex((*mymap)[i], 'X', j)
						return
					}
					nextLineRunes := []rune((*mymap)[i])
					next := nextLineRunes[j-1]
					if isObstacle(next) {
						(*mymap)[i] = replaceAtIndex(l, '^', j)
					} else {
						(*mymap)[i] = replaceAtIndex((*mymap)[i], '<', j-1)
						(*mymap)[i] = replaceAtIndex((*mymap)[i], 'X', j)
					}
					return
				}

			}
		}
	}
}

type Cord struct {
	x int
	y int
	c rune
}

func isInLoop(mymap []string, cords map[Cord]int) bool {
	for _, cordv := range cords {
		if cordv > 1 {
			return true
		}
	}
	return false
}

func guardWalk2(mymap *[]string, turnHistory *map[Cord]int) {
	guard := []rune{'v', '^', '<', '>'}
	for i, l := range *mymap {
		for j, c := range l {
			if slices.Contains(guard, c) {
				if c == '^' {
					if i-1 < 0 {
						(*mymap)[i] = replaceAtIndex((*mymap)[i], 'X', j)
						return
					}
					nextLineRunes := []rune((*mymap)[i-1])
					next := nextLineRunes[j]
					if isObstacle(next) {
						(*mymap)[i] = replaceAtIndex(l, '>', j)
						(*turnHistory)[Cord{j, i, c}] += 1
					} else {
						(*mymap)[i-1] = replaceAtIndex((*mymap)[i-1], '^', j)
						(*mymap)[i] = replaceAtIndex((*mymap)[i], 'X', j)
					}
					return
				}
				if c == '>' {
					nextLineRunes := []rune((*mymap)[i])
					if j+1 == len(nextLineRunes) {
						(*mymap)[i] = replaceAtIndex((*mymap)[i], 'X', j)
						return
					}
					next := nextLineRunes[j+1]
					if isObstacle(next) {
						(*mymap)[i] = replaceAtIndex(l, 'v', j)
						(*turnHistory)[Cord{j, i, c}] += 1
					} else {
						(*mymap)[i] = replaceAtIndex((*mymap)[i], '>', j+1)
						(*mymap)[i] = replaceAtIndex((*mymap)[i], 'X', j)
					}
				}
				if c == 'v' {
					if i+1 == len(*mymap) {
						(*mymap)[i] = replaceAtIndex((*mymap)[i], 'X', j)
						return
					}
					nextLineRunes := []rune((*mymap)[i+1])
					next := nextLineRunes[j]
					if isObstacle(next) {
						(*mymap)[i] = replaceAtIndex(l, '<', j)
						(*turnHistory)[Cord{j, i, c}] += 1
					} else {
						(*mymap)[i+1] = replaceAtIndex((*mymap)[i+1], 'v', j)
						(*mymap)[i] = replaceAtIndex((*mymap)[i], 'X', j)
					}
				}
				if c == '<' {
					if j-1 < 0 {
						(*mymap)[i] = replaceAtIndex((*mymap)[i], 'X', j)
						return
					}
					nextLineRunes := []rune((*mymap)[i])
					next := nextLineRunes[j-1]
					if isObstacle(next) {
						(*mymap)[i] = replaceAtIndex(l, '^', j)
						(*turnHistory)[Cord{j, i, c}] += 1
					} else {
						(*mymap)[i] = replaceAtIndex((*mymap)[i], '<', j-1)
						(*mymap)[i] = replaceAtIndex((*mymap)[i], 'X', j)
					}
					return
				}

			}
		}
	}
}

func printMap(mymap []string) {
	for _, l := range mymap {
		fmt.Println(l)
	}
	fmt.Println()
}

func evaluateMap(mymap []string) int {
	res := 0
	for _, s := range mymap {
		res += strings.Count(s, "X")
	}
	return res
}

func part1(filename string) int {
	mymap := input.ReadLines(filename)
	for isGuardOnMap(mymap) {
		guardWalk(&mymap)
		// printMap(mymap)
	}

	return evaluateMap(mymap)
}

func part2(filename string) int {
	origimymap := input.ReadLines(filename)
	max_x, max_y := len(origimymap[0]), len(origimymap)
	res := 0
	for y := 0; y < max_y; y++ {
		for x := 0; x < max_x; x++ {
			mymap := make([]string, len(origimymap))
			copy(mymap, origimymap)
			mymap[y] = replaceAtIndex(mymap[y], '#', x)
			turnHistory := make(map[Cord]int, 0)
			for isGuardOnMap(mymap) {
				guardWalk2(&mymap, &turnHistory)
				if isInLoop(mymap, turnHistory) {
					res += 1
					fmt.Println(res)
					break
				}
			}
		}
	}

	return res
}
