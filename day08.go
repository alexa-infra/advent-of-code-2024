package main

import (
	"bufio"
	"fmt"
	"os"
)

type coord2d struct {
	x, y int
}

func (a coord2d) Add(b coord2d) coord2d {
	return coord2d{a.x + b.x, a.y + b.y}
}
func (a coord2d) Sub(b coord2d) coord2d {
	return coord2d{a.x - b.x, a.y - b.y}
}

func main() {
	data := [][]byte{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data = append(data, scanner.Bytes())
	}
	m, n := len(data), len(data[0])
	inBound := func(p coord2d) bool {
		return p.x >= 0 && p.x < m && p.y >= 0 && p.y < n
	}
	ant := map[byte][]coord2d{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ch := data[i][j]
			if ch != '.' {
				p := coord2d{i, j}
				ant[ch] = append(ant[ch], p)
			}
		}
	}
	p1 := map[coord2d]bool{}
	for _, arr := range ant {
		k := len(arr)
		for i := 0; i < k; i++ {
			a := arr[i]
			for j := i + 1; j < k; j++ {
				b := arr[j]
				dd := coord2d{b.x - a.x, b.y - a.y}
				if r1 := b.Add(dd); inBound(r1) {
					p1[r1] = true
				}
				if r2 := a.Sub(dd); inBound(r2) {
					p1[r2] = true
				}
			}
		}
	}
	fmt.Println("Print 1:", len(p1))
	p2 := map[coord2d]bool{}
	for _, arr := range ant {
		k := len(arr)
		if k == 1 {
			p2[arr[0]] = true
			continue
		}
		for i := 0; i < k; i++ {
			a := arr[i]
			for j := i + 1; j < k; j++ {
				b := arr[j]
				dd := coord2d{b.x - a.x, b.y - a.y}
				for c := b; inBound(c); c = c.Add(dd) {
					p2[c] = true
				}
				for c := a; inBound(c); c = c.Sub(dd) {
					p2[c] = true
				}
			}
		}
	}
	fmt.Println("Print 2:", len(p2))
}
