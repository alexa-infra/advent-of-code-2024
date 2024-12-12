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
	type pair struct {
		x, y int
	}
	starts := []pair{}
	m, n := len(data), len(data[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if data[i][j] == '0' {
				starts = append(starts, pair{i, j})
			}
		}
	}
	inBounds := func(p pair) bool {
		return p.x >= 0 && p.x < m && p.y >= 0 && p.y < n
	}
	follow := func(start pair) int {
		next := []pair{start}
		visited := map[pair]bool{start: true}
		r := 0
		for len(next) > 0 {
			cur := next[0]
			next = next[1:]
			val := data[cur.x][cur.y]
			dx, dy := 1, 0
			for i := 0; i < 4; i++ {
				d := pair{cur.x + dx, cur.y + dy}
				if inBounds(d) {
					nextVal := data[d.x][d.y]
					if nextVal == val+byte(1) {
						if _, ok := visited[d]; !ok {
							visited[d] = true
							if nextVal == '9' {
								r++
							} else {
								next = append(next, d)
							}
						}
					}
				}
				dx, dy = -dy, dx
			}
		}
		return r
	}
	p1 := 0
	for _, p := range starts {
		p1 += follow(p)
	}
	fmt.Println("Part 1:", p1)

	follow2 := func(start pair) int {
		next := []pair{start}
		r := 0
		for len(next) > 0 {
			cur := next[0]
			next = next[1:]
			val := data[cur.x][cur.y]
			dx, dy := 1, 0
			for i := 0; i < 4; i++ {
				d := pair{cur.x + dx, cur.y + dy}
				if inBounds(d) {
					nextVal := data[d.x][d.y]
					if nextVal == val+byte(1) {
						if nextVal == '9' {
							r++
						} else {
							next = append(next, d)
						}
					}
				}
				dx, dy = -dy, dx
			}
		}
		return r
	}
	p2 := 0
	for _, p := range starts {
		p2 += follow2(p)
	}
	fmt.Println("Part 2:", p2)
}
