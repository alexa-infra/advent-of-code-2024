package main

import (
	"fmt"
	"os"
	"bufio"
	"regexp"
	"strconv"
	"math/big"
)

func main() {
	re1 := regexp.MustCompile(`Button (A|B): X\+(\d+), Y\+(\d+)`)
	re2 := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)
	scanner := bufio.NewScanner(os.Stdin)
	type pair struct {
		x, y int
	}
	type cond struct {
		a, b, c pair
	}
	games := []cond{}
	for scanner.Scan() {
		line1 := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()
		scanner.Scan()
		line3 := scanner.Text()
		scanner.Scan()

		m1 := re1.FindStringSubmatch(line1)
		m2 := re1.FindStringSubmatch(line2)
		m3 := re2.FindStringSubmatch(line3)
		ax, _ := strconv.Atoi(m1[2])
		ay, _ := strconv.Atoi(m1[3])
		bx, _ := strconv.Atoi(m2[2])
		by, _ := strconv.Atoi(m2[3])
		cx, _ := strconv.Atoi(m3[1])
		cy, _ := strconv.Atoi(m3[2])
		a, b, c := pair{ax, ay}, pair{bx, by}, pair{cx, cy}
		games = append(games, cond{a, b, c})
	}
	p1 := 0
	for _, game := range games {
		a, b, c := game.a, game.b, game.c
		y1 := a.x * c.y - a.y * c.x
		y2 := a.x * b.y - a.y * b.x
		if y2 == 0 || y1 % y2 != 0 {
			continue
		}
		y := y1 / y2
		x1 := c.x - b.x * y
		x2 := a.x
		if x2 == 0 || x1 % x2 != 0 {
			continue
		}
		x := x1 / x2
		if x < 0 || y < 0 || x > 100 || y > 100 {
			continue
		}
		p1 += 3 * x + y
	}
	fmt.Println("Part 1:", p1)
	shift := new(big.Int)
	shift.SetString("10000000000000", 10)
	
	type pairBig struct {
		x, y *big.Int
	}
	conv := func(p pair) pairBig {
		a := new(big.Int)
		a.SetInt64(int64(p.x))
		b := new(big.Int)
		b.SetInt64(int64(p.y))
		return pairBig{ a, b }
	}
	zero := new(big.Int)
	zero.SetString("0", 10)
	p2 := new(big.Int)
	t1 := new(big.Int)
	t2 := new(big.Int)
	y1 := new(big.Int)
	y2 := new(big.Int)
	y := new(big.Int)
	x1 := new(big.Int)
	x2 := new(big.Int)
	x := new(big.Int)
	for _, game := range games {
		a, b, c := conv(game.a), conv(game.b), conv(game.c)
		c.x.Add(c.x, shift)
		c.y.Add(c.y, shift)
		t1.Mul(a.x, c.y)
		t2.Mul(a.y, c.x)
		y1.Sub(t1, t2)
		//y1 := a.x * c.y - a.y * c.x
		t1.Mul(a.x, b.y)
		t2.Mul(a.y, b.x)
		y2.Sub(t1, t2)
		//y2 := a.x * b.y - a.y * b.x
		if y2.Cmp(zero) == 0 {
			continue
		}
		mod := new(big.Int)
		mod.Mod(y1, y2)
		if mod.Cmp(zero) != 0 {
			continue
		}
		y.Div(y1, y2)
		t1.Mul(b.x, y)
		x1.Sub(c.x, t1)
		//x1 := c.x - b.x * y
		x2.Set(a.x)
		// x2 := a.x
		if x2.Cmp(zero) == 0 {
			continue
		}
		mod.Mod(x1, x2)
		if mod.Cmp(zero) != 0 {
			continue
		}
		x.Div(x1, x2)
		//x := x1 / x2

		if x.Cmp(zero) < 0 || y.Cmp(zero) < 0 {
			continue
		}
		t1.SetInt64(int64(3))
		x.Mul(x, t1)
		x.Add(x, y)
		p2.Add(p2, x)
	}
	fmt.Println("Part 2:", p2.Text(10))
}
