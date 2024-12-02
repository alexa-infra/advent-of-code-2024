package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	mat := [][]int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		row := []int{}
		for _, part := range parts {
			val, _ := strconv.Atoi(part)
			row = append(row, val)
		}
		mat = append(mat, row)
	}
	checkSafe := func(row []int) bool {
		d1 := row[1] - row[0]
		if d1 == 0 || abs(d1) > 3 {
			return false
		}
		for j := 2; j < len(row); j++ {
			a, b := row[j-1], row[j]
			d2 := b - a
			if d1*d2 <= 0 || abs(d2) > 3 {
				return false
			}
		}
		return true
	}
	p1 := 0
	for i := 0; i < len(mat); i++ {
		row := mat[i]
		if checkSafe(row) {
			p1++
		}
	}
	fmt.Println("Part 1:", p1)
	p2 := 0
	for i := 0; i < len(mat); i++ {
		row := mat[i]
		if checkSafe(row) {
			p2++
		} else {
			for badIndex := 0; badIndex < len(row); badIndex++ {
				newRow := []int{}
				newRow = append(newRow, row[:badIndex]...)
				newRow = append(newRow, row[badIndex+1:]...)
				if checkSafe(newRow) {
					p2++
					break
				}
			}
		}
	}
	fmt.Println("Part 2:", p2)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
