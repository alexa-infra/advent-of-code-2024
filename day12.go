package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	data := [][]byte{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data = append(data, []byte(scanner.Text()))
	}
	m, n := len(data), len(data[0])
	type pair struct {
		x, y int
	}
	type tuple struct {
		p pair
		k int
	}
	inBound := func(p pair) bool {
		return p.x >= 0 && p.x < m && p.y >= 0 && p.y < n
	}
	visited := map[pair]bool{}
	p1, p2 := 0, 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			p := pair{i, j}
			if _, ok := visited[p]; ok {
				continue
			}
			ch := data[i][j]
			next := []pair{p}
			seen := map[pair]bool{}
			seen[p] = true
			area := 0
			borders := map[tuple]bool{}
			for len(next) > 0 {
				p := next[0]
				next = next[1:]
				visited[p] = true
				area += 1
				dx, dy := 1, 0
				for k := 0; k < 4; k++ {
					pp := pair{p.x + dx, p.y + dy}
					if inBound(pp) {
						ch2 := data[pp.x][pp.y]
						if ch == ch2 {
							if _, ok := seen[pp]; !ok {
								seen[pp] = true
								next = append(next, pp)
							}
						} else {
							borders[tuple{p, k}] = false
						}
					} else {
						borders[tuple{p, k}] = false
					}
					dx, dy = -dy, dx
				}
			}
			perimeter := len(borders)
			p1 += perimeter * area
			sides := 0
			for border, counted := range borders {
				if counted {
					continue
				}
				sides += 1
				p := border.p
				dx, dy := 1, 0
				for k := 0; k < 4; k++ {
					for times := 1; times < max(n, m); times++ {
						pp := pair{p.x + dx*times, p.y + dy*times}
						t := tuple{pp, border.k}
						if _, ok := borders[t]; ok {
							borders[t] = true
						} else {
							break
						}
					}
					for times := 1; times < max(n, m); times++ {
						pp := pair{p.x - dx*times, p.y - dy*times}
						t := tuple{pp, border.k}
						if _, ok := borders[t]; ok {
							borders[t] = true
						} else {
							break
						}
					}
					dx, dy = -dy, dx
				}
			}
			p2 += sides * area
		}
	}
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
