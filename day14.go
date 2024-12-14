package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	re1 := regexp.MustCompile(`(\d+) (\d+)`)
	re2 := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	match := re1.FindStringSubmatch(scanner.Text())
	m, _ := strconv.Atoi(match[1])
	n, _ := strconv.Atoi(match[2])
	type pair struct {
		x, y int
	}
	type pos struct {
		p, v pair
	}
	data := []pos{}
	for scanner.Scan() {
		match = re2.FindStringSubmatch(scanner.Text())
		px, _ := strconv.Atoi(match[1])
		py, _ := strconv.Atoi(match[2])
		vx, _ := strconv.Atoi(match[3])
		vy, _ := strconv.Atoi(match[4])
		r := pos{pair{px, py}, pair{vx, vy}}
		data = append(data, r)
	}
	mod := func(a, b int) int {
		r := a % b
		if r < 0 {
			return r + b
		}
		return r
	}
	run := func(steps int) map[pair]bool {
		g := map[pair]bool{}
		for _, d := range data {
			x := mod(d.p.x+d.v.x*steps, m)
			y := mod(d.p.y+d.v.y*steps, n)
			pp := pair{x, y}
			g[pp] = true
		}
		return g
	}
	g := run(100)
	var r1, r2, r3, r4 int
	for p := range g {
		x, y := p.x, p.y
		if x == m/2 || y == n/2 {
			continue
		}
		if x < m/2 && y < n/2 {
			r1++
		}
		if x < m/2 && y > n/2 {
			r2++
		}
		if x > m/2 && y < n/2 {
			r3++
		}
		if x > m/2 && y > n/2 {
			r4++
		}
	}
	p1 := r1 * r2 * r3 * r4
	fmt.Println("Part 1:", p1)

	hasBox := func(mm map[pair]bool) bool {
		for p := range mm {
			next := []pair{p}
			uniq := map[pair]bool{p: true}
			visited := map[pair]bool{}
			for len(next) > 0 {
				cur := next[0]
				next = next[1:]
				visited[cur] = true
				dx, dy := 1, 0
				for i := 0; i < 4; i++ {
					pp := pair{cur.x + dx, cur.y + dy}
					if _, ok := mm[pp]; ok {
						if _, ok := uniq[pp]; !ok {
							uniq[pp] = true
							next = append(next, pp)
						}
					}
					dx, dy = -dy, dx
				}
			}
			if len(visited) > 200 {
				return true
			}
		}
		return false
	}
	//printMap := func(mm map[pair]bool) {
	//	for i := 0; i < m; i++ {
	//		for j := 0; j < n; j++ {
	//			pp := pair{i, j}
	//			if _, ok := mm[pp]; ok {
	//				fmt.Print("x")
	//			} else {
	//				fmt.Print(".")
	//			}
	//		}
	//		fmt.Println()
	//	}
	//}
	for steps := 0; steps < 4000000; steps++ {
		mm := run(steps)
		if hasBox(mm) {
			//printMap(mm)
			fmt.Println("Part 2:", steps)
			break
		}
	}
}
