package main

import (
	"bufio"
	"fmt"
	pq "github.com/emirpasic/gods/queues/priorityqueue"
	"os"
)

func main() {
	data := [][]byte{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row := []byte(scanner.Text())
		data = append(data, row)
	}
	m, n := len(data), len(data[0])
	type pair struct {
		x, y int
	}
	type pos struct {
		p, dir pair
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
	var start, end pos
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ch := data[i][j]
			p := pair{i, j}
			if ch == 'S' {
				start = pos{p, pair{0, 1}}
			} else if ch == 'E' {
				end = pos{p, pair{0, 0}}
			}
		}
	}
	h := func(a pos) int {
		distA := distance(a.p, end.p) + 1000
		if a.dir.x == 1 || a.dir.y == -1 {
			distA += 1000
		}
		return distA
	}
	cameFrom := map[pos][]pos{}
	gScore := map[pos]int{}
	gScore[start] = 0
	fScore := map[pos]int{}
	fScore[start] = h(start)
	byFScore := func(a, b interface{}) int {
		ap, bp := a.(pos), b.(pos)
		return fScore[ap] - fScore[bp]
	}
	queue := pq.NewWith(byFScore)
	queue.Enqueue(start)
	minVal := 0
	found := false
	addNeighbor := func(current, next pos, scores int) {
		if data[next.p.x][next.p.y] == '#' {
			return
		}
		tentative_gScore := gScore[current] + scores
		g, ok := gScore[next]
		if !ok || tentative_gScore < g {
			cameFrom[next] = []pos{current}
			gScore[next] = tentative_gScore
			fScore[next] = tentative_gScore + h(next)
			queue.Enqueue(next)
		} else if ok && tentative_gScore == g {
			cameFrom[next] = append(cameFrom[next], current)
		}
	}
	for !queue.Empty() {
		a, _ := queue.Dequeue()
		current := a.(pos)
		at, dir := current.p, current.dir
		if at == end.p {
			found = true
			minVal = gScore[current]
			continue
		}
		if found && gScore[current] > minVal {
			continue
		}
		f := pos{pair{at.x + dir.x, at.y + dir.y}, dir}
		addNeighbor(current, f, 1)
		left := pair{-dir.y, dir.x}
		l := pos{pair{at.x + left.x, at.y + left.y}, left}
		addNeighbor(current, l, 1001)
		right := pair{dir.y, -dir.x}
		r := pos{pair{at.x + right.x, at.y + right.y}, right}
		addNeighbor(current, r, 1001)
	}
	fmt.Println("Part 1:", minVal)
	uniq := map[pair]bool{}
	uniq[start.p] = true
	uniq[end.p] = true
	var mark func(pos)
	mark = func(p pos) {
		prev, ok := cameFrom[p]
		if !ok {
			return
		}
		for _, pp := range prev {
			uniq[pp.p] = true
			mark(pp)
		}
	}
	for c := range cameFrom {
		if c.p == end.p {
			mark(c)
		}
	}
	fmt.Println("Part 2:", len(uniq))
}
