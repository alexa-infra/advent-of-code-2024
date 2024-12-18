package main

import (
	"bufio"
	"fmt"
	pq "github.com/emirpasic/gods/queues/priorityqueue"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	m, _ := strconv.Atoi(parts[0])
	n, _ := strconv.Atoi(parts[1])
	k, _ := strconv.Atoi(parts[2])
	type pair struct {
		x, y int
	}
	inData := []pair{}
	for scanner.Scan() {
		parts = strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		p := pair{x, y}
		inData = append(inData, p)
	}
	start, end := pair{0, 0}, pair{m, n}
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	distance := func(a, b pair) int {
		return abs(a.x-b.x) + abs(a.y-b.y)
	}
	h := func(a pair) int {
		return distance(a, end)
	}
	inBound := func(a pair) bool {
		return a.x >= 0 && a.x <= m && a.y >= 0 && a.y <= n
	}
	findWay := func(k int, anyWay bool) int {
		data := map[pair]bool{}
		for i := 0; i < k; i++ {
			p := inData[i]
			data[p] = true
		}
		gScore := map[pair]int{}
		gScore[start] = 0
		fScore := map[pair]int{}
		fScore[start] = h(start)
		byFScore := func(a, b interface{}) int {
			ap, bp := a.(pair), b.(pair)
			return fScore[bp] - fScore[ap]
		}
		queue := pq.NewWith(byFScore)
		queue.Enqueue(start)
		minSteps := -1
		for !queue.Empty() {
			a, _ := queue.Dequeue()
			current := a.(pair)
			if current == end {
				ms := gScore[current]
				if minSteps == -1 || ms < minSteps {
					minSteps = ms
				}
				if anyWay {
					break
				}
			}
			dx, dy := 1, 0
			for i := 0; i < 4; i++ {
				next := pair{current.x + dx, current.y + dy}
				if _, ok := data[next]; !ok && inBound(next) {
					tentative_gScore := gScore[current] + 1
					g, ok := gScore[next]
					if !ok || tentative_gScore < g {
						gScore[next] = tentative_gScore
						fScore[next] = tentative_gScore + h(next)
						queue.Enqueue(next)
					}
				}
				dx, dy = -dy, dx
			}
		}
		return minSteps
	}
	fmt.Println("Part 1:", findWay(k, false))
	for i := k; i <= len(inData); i++ {
		steps := findWay(i, true)
		if steps < 0 {
			fmt.Printf("Part 2: %d,%d\n", inData[i-1].x, inData[i-1].y)
			break
		}
	}
}
