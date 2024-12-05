package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	type pair struct {
		x, y int
	}
	less := map[pair]bool{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, "|")
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		p := pair{a, b}
		less[p] = true
	}
	printsets := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		set := []int{}
		for _, part := range parts {
			a, _ := strconv.Atoi(part)
			set = append(set, a)
		}
		printsets = append(printsets, set)
	}
	check := func(set []int) bool {
		has := map[int]bool{}
		for _, x := range set {
			has[x] = true
		}
		require := map[int]bool{}
		for _, x := range set {
			if _, ok := require[x]; ok {
				delete(require, x)
			}
			for y := range has {
				p1 := pair{x, y}
				if _, ok := less[p1]; ok {
					require[y] = true
				}
			}
		}
		return len(require) == 0
	}
	p1, p2 := 0, 0
	for _, set := range printsets {
		if check(set) {
			p1 += set[len(set)/2]
		} else {
			sort.Slice(set, func(i, j int) bool {
				x, y := set[i], set[j]
				p1 := pair{x, y}
				_, ok := less[p1]
				return ok
			})
			if check(set) {
				p2 += set[len(set)/2]
			} else {
				fmt.Println("FAIL")
			}
		}
	}
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
