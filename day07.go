package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"math/big"
)

func main() {
	eqs := [][]*big.Int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		p1, p2 := parts[0], strings.TrimSpace(parts[1])
		rest := strings.Split(p2, " ")
		eq := []*big.Int{}
		s := new(big.Int)
		s.SetString(p1, 10)
		eq = append(eq, s)
		for _, r := range rest {
			x := new(big.Int)
			x.SetString(r, 10)
			eq = append(eq, x)
		}
		eqs = append(eqs, eq)
	}
	var check func(*big.Int, []*big.Int) bool
	check = func(s *big.Int, parts []*big.Int) bool {
		if len(parts) == 1 {
			return s.Cmp(parts[0]) == 0
		}
		a, b, rest := parts[0], parts[1], parts[2:]
		if s.Cmp(a) < 0 || s.Cmp(b) < 0 {
			return false
		}
		c1 := new(big.Int)
		c1.Add(a, b)
		arr1 := []*big.Int{ c1 }
		arr1 = append(arr1, rest...)
		c2 := new(big.Int)
		c2.Mul(a, b)
		arr2 := []*big.Int{ c2 }
		arr2 = append(arr2, rest...)
		return check(s, arr1) || check(s, arr2)
	}
	r1 := new(big.Int)
	for _, eq := range eqs {
		if check(eq[0], eq[1:]) {
			r1.Add(r1, eq[0])
		}
	}
	fmt.Println("Part 1:", r1)

	var check2 func(*big.Int, []*big.Int) bool
	check2 = func(s *big.Int, parts []*big.Int) bool {
		if len(parts) == 1 {
			return s.Cmp(parts[0]) == 0
		}
		a, b, rest := parts[0], parts[1], parts[2:]
		if s.Cmp(a) < 0 || s.Cmp(b) < 0 {
			return false
		}
		c1 := new(big.Int)
		c1.Add(a, b)
		arr1 := []*big.Int{ c1 }
		arr1 = append(arr1, rest...)
		c2 := new(big.Int)
		c2.Mul(a, b)
		arr2 := []*big.Int{ c2 }
		arr2 = append(arr2, rest...)
		c3 := new(big.Int)
		c3.SetString(a.Text(10) + b.Text(10), 10)
		arr3 := []*big.Int{ c3 }
		arr3 = append(arr3, rest...)
		return check2(s, arr1) || check2(s, arr2) || check2(s, arr3)
	}
	r2 := new(big.Int)
	for _, eq := range eqs {
		if check2(eq[0], eq[1:]) {
			r2.Add(r2, eq[0])
		}
	}
	fmt.Println("Part 2:", r2)
}
