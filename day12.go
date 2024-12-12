package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	x, y int
}
type tuple struct {
	p pair
	k int
}

func main() {
	data := [][]byte{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data = append(data, []byte(scanner.Text()))
	}
	m, n := len(data), len(data[0])
	getData := func(p pair) (byte, bool){
		if p.x >= 0 && p.x < m && p.y >= 0 && p.y < n {
			return data[p.x][p.y], true
		}
		return byte(0), false
	}
	visited := map[pair]bool{}
	p1, p2 := 0, 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			p := pair{i, j}
			if _, ok := visited[p]; ok {
				continue
			}
			seen, borders := explore(p, getData, func(p pair) {
				visited[p] = true
			})
			area := len(seen)
			perimeter := len(borders)
			p1 += perimeter * area
			sides := bordersToSides(borders)
			p2 += sides * area
		}
	}
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}

func explore(p pair, getData func (p pair) (byte, bool), visit func(p pair)) (map[pair]bool, map[tuple]bool) {
	ch, _ := getData(p)
	next := []pair{p}
	seen := map[pair]bool{}
	seen[p] = true
	borders := map[tuple]bool{}
	for len(next) > 0 {
		p := next[0]
		next = next[1:]
		visit(p)
		dx, dy := 1, 0
		for k := 0; k < 4; k++ {
			pp := pair{p.x + dx, p.y + dy}
			if ch2, ok := getData(pp); ok && ch == ch2 {
				if _, ok := seen[pp]; !ok {
					seen[pp] = true
					next = append(next, pp)
				}
			} else {
				borders[tuple{p, k}] = false
			}
			dx, dy = -dy, dx
		}
	}
	return seen, borders
}

func bordersToSides(borders map[tuple]bool) int {
	sides := 0
	for border, counted := range borders {
		if counted {
			continue
		}
		sides += 1
		p := border.p
		dx, dy := 1, 0
		for k := 0; k < 4; k++ {
			for times := 1; ; times++ {
				pp := pair{p.x + dx*times, p.y + dy*times}
				t := tuple{pp, border.k}
				if _, ok := borders[t]; ok {
					borders[t] = true
				} else {
					break
				}
			}
			for times := 1; ; times++ {
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
	return sides
}
