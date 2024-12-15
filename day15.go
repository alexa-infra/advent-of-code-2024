package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	initialData := [][]byte{}
	moves := []byte{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		initialData = append(initialData, []byte(scanner.Text()))
	}
	for scanner.Scan() {
		line := []byte(scanner.Text())
		moves = append(moves, line...)
	}
	data := [][]byte{}
	for _, row := range initialData {
		data = append(data, append([]byte{}, row...))
	}
	m, n := len(data), len(data[0])
	type pair struct {
		x, y int
	}
	var pos pair
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if data[i][j] == '@' {
				pos = pair{i, j}
			}
		}
	}
	printMap := func() {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				fmt.Print(string([]byte{data[i][j]}))
			}
			fmt.Println()
		}
		fmt.Println()
	}
	mapMoves := map[byte]pair{
		'<': pair{0, -1},
		'>': pair{0, 1},
		'^': pair{-1, 0},
		'v': pair{1, 0},
	}
	for _, move := range moves {
		diff, ok := mapMoves[move]
		if !ok {
			continue
		}
		if data[pos.x][pos.y] != '@' {
			printMap()
			panic("wrong pos")
		}
		nextPos := pair{pos.x + diff.x, pos.y + diff.y}
		ch := data[nextPos.x][nextPos.y]
		if ch == '#' {
			continue
		}
		if ch == '.' {
			data[pos.x][pos.y] = '.'
			data[nextPos.x][nextPos.y] = '@'
			pos = nextPos
			continue
		}
		if ch == 'O' {
			nextPos2 := nextPos
			for data[nextPos2.x][nextPos2.y] == 'O' {
				nextPos2 = pair{nextPos2.x + diff.x, nextPos2.y + diff.y}
			}
			ch2 := data[nextPos2.x][nextPos2.y]
			if ch2 == '#' {
				continue
			}
			if ch2 == '.' {
				data[pos.x][pos.y] = '.'
				data[nextPos.x][nextPos.y] = '@'
				data[nextPos2.x][nextPos2.y] = 'O'
				pos = nextPos
			}
		}
	}
	p1 := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if data[i][j] == 'O' {
				p1 += i*100 + j
			}
		}
	}
	fmt.Println("Part 1:", p1)

	data = [][]byte{}
	for _, row := range initialData {
		nrow := []byte{}
		for _, ch := range row {
			if ch == '#' {
				nrow = append(nrow, '#', '#')
			}
			if ch == '.' {
				nrow = append(nrow, '.', '.')
			}
			if ch == 'O' {
				nrow = append(nrow, '[', ']')
			}
			if ch == '@' {
				nrow = append(nrow, '@', '.')
			}
		}
		data = append(data, nrow)
	}
	m, n = len(data), len(data[0])
	var canMove func(pair, pair) bool
	canMove = func(p, dir pair) bool {
		if data[p.x][p.y] == '.' {
			return true
		}
		if data[p.x][p.y] == '#' {
			return false
		}
		if dir.x != 0 {
			if data[p.x][p.y] == '[' {
				return canMove(pair{p.x + dir.x, p.y}, dir) && canMove(pair{p.x + dir.x, p.y + 1}, dir)
			}
			if data[p.x][p.y] == ']' {
				return canMove(pair{p.x + dir.x, p.y}, dir) && canMove(pair{p.x + dir.x, p.y - 1}, dir)
			}
			return true
		} else {
			return canMove(pair{p.x, p.y + dir.y}, dir)
		}
	}
	var doMove func(pair, pair, pair)
	doMove = func(p1, p2, dir pair) {
		a, b := data[p1.x][p1.y], data[p2.x][p2.y]
		if a != '[' || b != ']' {
			printMap()
			fmt.Println(string([]byte{a, b}), dir)
			panic("wrong doMove")
		}
		if dir.x != 0 {
			n1 := pair{p1.x + dir.x, p1.y}
			n2 := pair{p2.x + dir.x, p2.y}
			if data[n1.x][n1.y] == '[' && data[n2.x][n2.y] == ']' {
				doMove(n1, n2, dir)
			}
			if data[n1.x][n1.y] == ']' {
				doMove(pair{n1.x, n1.y - 1}, n1, dir)
			}
			if data[n2.x][n2.y] == '[' {
				doMove(n2, pair{n2.x, n2.y + 1}, dir)
			}
			data[n1.x][n1.y] = '['
			data[n2.x][n2.y] = ']'
			data[p1.x][p1.y] = '.'
			data[p2.x][p2.y] = '.'
		} else if dir.y > 0 {
			n1 := pair{p2.x, p2.y + dir.y}
			if data[n1.x][n1.y] != '.' {
				n2 := pair{n1.x, n1.y + dir.y}
				doMove(n1, n2, dir)
			}
			data[p1.x][p1.y] = '.'
			data[p2.x][p2.y] = '['
			data[n1.x][n1.y] = ']'
		} else {
			n1 := pair{p1.x, p1.y + dir.y}
			if data[n1.x][n1.y] != '.' {
				n2 := pair{n1.x, n1.y + dir.y}
				doMove(n2, n1, dir)
			}
			data[p2.x][p2.y] = '.'
			data[n1.x][n1.y] = '['
			data[p1.x][p1.y] = ']'
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if data[i][j] == '@' {
				pos = pair{i, j}
			}
		}
	}
	for _, move := range moves {
		diff, ok := mapMoves[move]
		if !ok {
			continue
		}
		if data[pos.x][pos.y] != '@' {
			printMap()
			panic("wrong pos")
		}
		nextPos := pair{pos.x + diff.x, pos.y + diff.y}
		ch := data[nextPos.x][nextPos.y]
		if ch == '#' {
			continue
		}
		if ch == '.' {
			data[pos.x][pos.y] = '.'
			data[nextPos.x][nextPos.y] = '@'
			pos = nextPos
			continue
		}
		if ch == '[' || ch == ']' {
			if canMove(nextPos, diff) {
				if ch == '[' {
					doMove(nextPos, pair{nextPos.x, nextPos.y + 1}, diff)
				} else {
					doMove(pair{nextPos.x, nextPos.y - 1}, nextPos, diff)
				}
				data[pos.x][pos.y] = '.'
				data[nextPos.x][nextPos.y] = '@'
				pos = nextPos
			}
		}
	}
	p2 := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if data[i][j] == '[' {
				p2 += i*100 + j
			}
		}
	}
	fmt.Println("Part 2:", p2)
}
