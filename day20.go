package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inData := [][]byte{}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	threshold, _ := strconv.Atoi(scanner.Text())
	for scanner.Scan() {
		inData = append(inData, []byte(scanner.Text()))
	}
	m, n := len(inData), len(inData[0])
	type pair struct {
		x, y int
	}
	var start, end pair
	walls := []pair{}
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			ch := inData[i][j]
			p := pair{i, j}
			if ch == 'S' {
				start = p
			} else if ch == 'E' {
				end = p
			} else if ch == '#' {
				walls = append(walls, p)
			}
		}
	}
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	distance := func(a, b pair) int {
		return abs(a.x-b.x) + abs(a.y-b.y)
	}
	inBound := func(a pair) bool {
		return a.x > 0 && a.x < m-1 && a.y > 0 && a.y < n-1
	}
	data := map[pair]bool{}
	for _, wall := range walls {
		data[wall] = true
	}
	noWall := func(a pair) bool {
		_, ok := data[a]
		return !ok
	}
	findDejkstra := func(start pair) map[pair]int {
		visited := map[pair]int{start: 0}
		queue := []pair{start}
		step := 1
		for len(queue) > 0 {
			n := len(queue)
			for i := 0; i < n; i++ {
				current := queue[i]
				dx, dy := 1, 0
				for i := 0; i < 4; i++ {
					next := pair{current.x + dx, current.y + dy}
					if inBound(next) && noWall(next) {
						if _, ok := visited[next]; !ok {
							visited[next] = step
							queue = append(queue, next)
						}
					}
					dx, dy = -dy, dx
				}
			}
			queue = queue[n:]
			step++
		}
		return visited
	}
	fromStart := findDejkstra(start)
	fromEnd := findDejkstra(end)
	noCheats := fromStart[end]
	iterWithinManhattanDist := func(p pair, dd int) func(func(pair)bool) {
		return func(yield func(pair) bool) {
			for dx := -dd; dx <= dd; dx++ {
				for dy := -dd; dy <= dd; dy++ {
					pp := pair{p.x + dx, p.y + dy}
					if distance(pp, p) <= dd {
						if !yield(pp) {
							return
						}
					}
				}
			}
		}
	}
	iterEmptySpaces := func() func(func(pair)bool) {
		return func(yield func(pair) bool) {
			for i := 1; i < m-1; i++ {
				for j := 1; j < n-1; j++ {
					p := pair{i, j}
					if noWall(p) {
						if !yield(p) {
							return
						}
					}
				}
			}
		}
	}
	solve := func(dd int) int {
		p1 := 0
		for p := range iterEmptySpaces() {
			beforeCheat, ok := fromStart[p]
			if !ok || beforeCheat >= noCheats {
				continue
			}
			for pp := range iterWithinManhattanDist(p, dd) {
				if p != pp && inBound(pp) && noWall(pp) {
					cheatTime := distance(pp, p)
					if beforeCheat+cheatTime >= noCheats {
						continue
					}
					afterCheat, ok := fromEnd[pp]
					if !ok {
						continue
					}
					newTime := afterCheat + cheatTime + beforeCheat
					if newTime+threshold <= noCheats {
						p1 += 1
					}
				}
			}
		}
		return p1
	}
	fmt.Println("Part 1:", solve(2))
	fmt.Println("Part 2:", solve(20))
}
