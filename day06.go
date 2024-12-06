package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	area := [][]byte{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row := []byte(scanner.Text())
		area = append(area, row)
	}
	type pos struct {
		x, y int
	}
	n, m := len(area), len(area[0])
	obst := map[pos]bool{}
	var start pos
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ch := area[i][j]
			p := pos{i, j}
			if ch == '^' {
				start = p
			} else if ch == '#' {
				obst[p] = true
			}
		}
	}
	current := start
	visited := map[pos]bool{}
	dx, dy := -1, 0
	for current.x >= 0 && current.x < m && current.y >= 0 && current.y < n {
		visited[current] = true
		next := pos{current.x + dx, current.y + dy}
		if _, ok := obst[next]; ok {
			dx, dy = dy, -dx
		} else {
			current = next
		}
	}
	p1 := len(visited)
	fmt.Println("Part 1:", p1)
	type tuple struct {
		x, y, d int
	}
	p2 := 0
	for wp := range visited {
		if wp == start {
			continue
		}
		obst[wp] = true
		current = start
		visited2 := map[tuple]bool{}
		dx, dy = -1, 0
		dir := 0
		for current.x >= 0 && current.x < m && current.y >= 0 && current.y < n {
			t := tuple{current.x, current.y, dir}
			if _, ok := visited2[t]; ok {
				p2++
				break
			}
			visited2[t] = true
			next := pos{current.x + dx, current.y + dy}
			if _, ok := obst[next]; ok {
				dx, dy = dy, -dx
				dir = (dir + 1) % 4
			} else {
				current = next
			}
		}
		delete(obst, wp)
	}
	fmt.Println("Part 2:", p2)
}
