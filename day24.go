package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	re1 := regexp.MustCompile(`([a-z0-9]{3}): (1|0)`)
	re2 := regexp.MustCompile(`([a-z0-9]{3}) (AND|OR|XOR) ([a-z0-9]{3}) -> ([a-z0-9]{3})`)
	scanner := bufio.NewScanner(os.Stdin)
	signals := map[string]int{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		match := re1.FindStringSubmatch(scanner.Text())
		var val int
		if match[2] == "1" {
			val = 1
		} else if match[2] == "0" {
			val = 0
		} else {
			panic("wrong value")
		}
		signals[match[1]] = val
	}
	inOps := [][4]string{}
	for scanner.Scan() {
		match := re2.FindStringSubmatch(scanner.Text())
		inOps = append(inOps, [4]string(match[1:]))
	}
	ops := make([][4]string, len(inOps))
	copy(ops, inOps)
	for len(ops) > 0 {
		newOps := [][4]string{}
		for _, op := range ops {
			a, b := op[0], op[2]
			val1, ok1 := signals[a]
			val2, ok2 := signals[b]
			if ok1 && ok2 {
				c := op[3]
				inst := op[1]
				if inst == "AND" {
					signals[c] = val1 & val2
				} else if inst == "OR" {
					signals[c] = val1 | val2
				} else if inst == "XOR" {
					signals[c] = val1 ^ val2
				} else {
					panic("wrong inst")
				}
			} else {
				newOps = append(newOps, op)
			}
		}
		ops = newOps
	}
	p1 := 0
	for k, v := range signals {
		if strings.HasPrefix(k, "z") {
			bit, _ := strconv.Atoi(k[1:])
			p1 |= v << bit
		}
	}
	fmt.Println("Part 1:", p1)

	isNamed := func(a string) bool {
		return strings.HasPrefix(a, "x") || strings.HasPrefix(a, "y") || strings.HasPrefix(a, "z")
	}

	ops = make([][4]string, len(inOps))
	copy(ops, inOps)
	wrong := map[string]bool{}
	for _, op := range ops {
		a, inst, b, res := op[0], op[1], op[2], op[3]
		if strings.HasPrefix(res, "z") && inst != "XOR" && res != "z45" {
			wrong[res] = true
		}
		if inst == "XOR" {
			if !isNamed(res) && !isNamed(a) && !isNamed(b) {
				wrong[res] = true
			}
			for _, subop := range ops {
				a1, inst1, b1 := subop[0], subop[1], subop[2]
				if (res == a1 || res == b1) && inst1 == "OR" {
					wrong[res] = true
				}
			}
		}
		if inst == "AND" && a != "x00" && b != "x00" {
			for _, subop := range ops {
				a1, inst1, b1 := subop[0], subop[1], subop[2]
				if (res == a1 || res == b1) && inst1 != "OR" {
					wrong[res] = true
				}
			}
		}
	}
	wrongKeys := []string{}
	for k := range wrong {
		wrongKeys = append(wrongKeys, k)
	}
	sort.Strings(wrongKeys)
	p2 := strings.Join(wrongKeys, ",")
	fmt.Println("Part 2:", p2)
}
