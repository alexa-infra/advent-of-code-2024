package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	mat := [][]byte{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		mat = append(mat, []byte(line))
	}
	n, m := len(mat), len(mat[0])
	type coord struct {
		x, y int
	}
	findInitial := func(ch byte) []coord {
		r := []coord{}
		for i, row := range mat {
			for j, val := range row {
				if val == ch {
					r = append(r, coord{i, j})
				}
			}
		}
		return r
	}
	find := func(pos []coord, ch byte, dx, dy int) []coord {
		r := []coord{}
		for _, p := range pos {
			next := coord{p.x + dx, p.y + dy}
			if next.x >= 0 && next.x < n && next.y >= 0 && next.y < m && mat[next.x][next.y] == ch {
				r = append(r, p)
			}
		}
		return r
	}
	p1 := 0
	xpos := findInitial('X')
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			mpos := find(xpos, 'M', dx, dy)
			apos := find(mpos, 'A', dx*2, dy*2)
			spos := find(apos, 'S', dx*3, dy*3)
			p1 += len(spos)
		}
	}
	fmt.Println("Part 1:", p1)
	pos := findInitial('A')
	p2 := 0
	xmas := func(ch1, ch2, ch3, ch4 byte) int {
		pos1 := find(pos, ch1, -1, -1)
		pos2 := find(pos1, ch2, 1, 1)
		pos3 := find(pos2, ch3, -1, 1)
		pos4 := find(pos3, ch4, 1, -1)
		return len(pos4)
	}
	p2 += xmas('M', 'S', 'M', 'S')
	p2 += xmas('S', 'M', 'M', 'S')
	p2 += xmas('M', 'S', 'S', 'M')
	p2 += xmas('S', 'M', 'S', 'M')
	fmt.Println("Part 2:", p2)
}
