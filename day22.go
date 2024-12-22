package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	mod := 16777216
	evolve := func(a int) int {
		a = ((a * 64) ^ a) % mod
		a = ((a / 32) ^ a) % mod
		a = ((a * 2048) ^ a) % mod
		return a
	}
	evolve2k := func(a int) int {
		for i := 0; i < 2000; i++ {
			a = evolve(a)
		}
		return a
	}
	nums := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, n)
	}
	s := 0
	for _, n := range nums {
		s += evolve2k(n)
	}
	fmt.Println("Part 1:", s)
	priceIterator := func(a int) func(func(int, [4]int) bool){
		return func(yield func(int, [4]int) bool) {
			last := [4]int{ 0, 0, 0, 0 }
			seen := map[[4]int]bool{}
			for i := 0; i < 2000; i++ {
				p1 := a % 10
				a = evolve(a)
				p2 := a % 10
				diff := p2 - p1
				if i < 4 {
					last[i] = diff
				} else {
					last[0], last[1], last[2], last[3] = last[1], last[2], last[3], diff
				}
				if i >= 3 {
					if _, ok := seen[last]; !ok {
						seen[last] = true
						if !yield(p2, last) {
							return
						}
					}
				}
			}
		}
	}
	uniq := map[[4]int]int{}
	for _, n := range nums {
		for price, last := range priceIterator(n) {
			uniq[last] += price
		}
	}
	maxp := 0
	for _, p := range uniq {
		if p > maxp {
			maxp = p
		}
	}
	fmt.Println("Part 2:", maxp)
}
