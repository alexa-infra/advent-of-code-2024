package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

var (
	keypadNumeric     = []string{
		"789",
		"456",
		"123",
		".0A",
	}
	keypadDirectional  = []string{
		".^A",
		"<v>",
	}
)

func findPosition(key string, keypad []string) (int, int) {
	for r, row := range keypad {
		for c, char := range row {
			if string(char) == key {
				return r, c
			}
		}
	}
	panic("Key not found")
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func buildPathBetween(key1, key2 string, keypad []string) []string {
	r1, c1 := findPosition(key1, keypad)
	r2, c2 := findPosition(key2, keypad)
	rGap, cGap := findPosition(".", keypad)
	dr, dc := r2 - r1, c2 - c1

	var rowMoves, colMoves string
	if dr >= 0 {
		rowMoves = strings.Repeat("v", abs(dr))
	} else {
		rowMoves = strings.Repeat("^", abs(dr))
	}
	if dc >= 0 {
		colMoves = strings.Repeat(">", abs(dc))
	} else {
		colMoves = strings.Repeat("<", abs(dc))
	}

	if dr == 0 && dc == 0 {
		return []string{""}
	}
	if dr == 0 {
		return []string{colMoves}
	}
	if dc == 0 {
		return []string{rowMoves}
	}
	if (r1 == rGap && c2 == cGap) {
		return []string{rowMoves + colMoves}
	}
	if (r2 == rGap && c1 == cGap) {
		return []string{colMoves + rowMoves}
	}
	return []string{rowMoves + colMoves, colMoves + rowMoves}
}

func buildSequencePath(seq string, keypad []string) [][]string {
	var res [][]string
	seq2 := "A" + seq
	for i := 1; i < len(seq2); i++ {
		key1 := string(seq2[i - 1])
		key2 := string(seq2[i])
		shortestPaths := buildPathBetween(key1, key2, keypad)
		s := []string{}
		for _, sp := range shortestPaths {
			s = append(s, sp + "A")
		}
		res = append(res, s)
	}
	return res
}

func main() {
	data := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	
	type args struct {
		seq string
		depth int
	}
	cache := map[args]int{}

	var solve func(seq string, depth int) int
	solve = func(seq string, depth int) int {
		if val, ok := cache[args{ seq, depth }]; ok {
			return val
		}
		
		if depth == 0 {
			cache[args{ seq, depth }] = len(seq)
			return len(seq)
		}

		var keypad []string
		if strings.ContainsAny(seq, "012345679") {
			keypad = keypadNumeric
		} else {
			keypad = keypadDirectional
		}

		sumPath := 0
		shortestPathsList := buildSequencePath(seq, keypad)
		for _, shortestPaths := range shortestPathsList {
			shortest := -1
			for _, sp := range shortestPaths {
				length := solve(sp, depth-1)
				if shortest == -1 || length < shortest {
					shortest = length
				}
			}
			if shortest != -1 {
				sumPath += shortest
			}
		}
		cache[args{ seq, depth }] = sumPath
		return sumPath
	}

	p1 := 0
	for _, code := range data {
		intVal, _ := strconv.Atoi(code[:3])
		p1 += solve(code, 1+2) * intVal
	}
	fmt.Println("Part 1:", p1)

	p2 := 0
	for _, code := range data {
		intVal, _ := strconv.Atoi(code[:3])
		p2 += solve(code, 1+25) * intVal
	}
	fmt.Println("Part 2:", p2)
}
