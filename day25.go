package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	keys := [][][]byte{}
	key := [][]byte{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			keys = append(keys, key)
			key = [][]byte{}
			continue
		}
		key = append(key, []byte(scanner.Text()))
	}
	keys = append(keys, key)

	up := map[[5]int]bool{}
	down := map[[5]int]bool{}
	for _, k := range keys {
		kind := k[0][0]
		pins := [5]int{0, 0, 0, 0, 0}
		for i := 0; i < 5; i++ {
			for j := 1; j < 7 && k[j][i] == kind; j++ {
				pins[i] += 1
			}
		}
		if kind == '#' {
			up[pins] = true
		} else {
			down[pins] = true
		}
	}
	p1 := 0
	for k1 := range up {
		for k2 := range down {
			fit := true
			for i := 0; i < 5; i++ {
				if k1[i] > k2[i] {
					fit = false
					break
				}
			}
			if fit {
				p1 += 1
			}
		}
	}
	fmt.Println("Part 1:", p1)
}
