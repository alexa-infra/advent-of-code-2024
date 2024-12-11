package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	data := strings.Split(scanner.Text(), " ")
	c2024 := big.NewInt(int64(2024))
	mult := func(a string) string {
		b := new(big.Int)
		b.SetString(a, 10)
		r := new(big.Int)
		r.Mul(b, c2024)
		return r.Text(10)
	}
	convert := func(a string) string {
		b := new(big.Int)
		b.SetString(a, 10)
		return b.Text(10)
	}
	split := func(a string) (string, string) {
		n := len(a)
		part1, part2 := a[:n/2], a[n/2:]
		return convert(part1), convert(part2)
	}
	arr := make([]string, len(data))
	copy(arr, data)
	//fmt.Println(arr)
	for i := 0; i < 25; i++ {
		next := []string{}
		for _, v := range arr {
			if v == "0" {
				next = append(next, "1")
			} else if len(v)%2 == 0 {
				a, b := split(v)
				next = append(next, a, b)
			} else {
				a := mult(v)
				next = append(next, a)
			}
		}
		//fmt.Println(next)
		arr = next
	}
	fmt.Println("Part 1:", len(arr))
	h := map[string]int64{}
	for _, v := range data {
		h[v] += int64(1)
	}
	for i := 0; i < 75; i++ {
		hh := map[string]int64{}
		for k, v := range h {
			if k == "0" {
				hh["1"] += v
			} else if len(k)%2 == 0 {
				a, b := split(k)
				hh[a] += v
				hh[b] += v
			} else {
				a := mult(k)
				hh[a] += v
			}
		}
		h = hh
	}
	p2 := int64(0)
	for _, v := range h {
		p2 += v
	}
	fmt.Println("Part 2:", p2)
}
