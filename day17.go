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
	re1 := regexp.MustCompile(`Register (A|B|C): (\d+)`)
	re2 := regexp.MustCompile(`Program: ([0-9,]+)`)
	inRegisters := map[string]int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		match := re1.FindStringSubmatch(scanner.Text())
		if match == nil {
			panic("Wrong format re1")
		}
		regName := match[1]
		regValue, err := strconv.Atoi(match[2])
		if err != nil {
			panic("Wrong number format")
		}
		inRegisters[regName] = regValue
	}
	scanner.Scan()
	match := re2.FindStringSubmatch(scanner.Text())
	if match == nil {
		panic("Wrong format re2")
	}
	parts := strings.Split(match[1], ",")
	nums := make([]int, len(parts))
	for i, part := range parts {
		value, err := strconv.Atoi(part)
		if err != nil {
			panic("Wrong number format")
		}
		nums[i] = value
	}
	run := func(a, b, c int64, n int) []int {
		registers := map[string]int64{
			"A": a,
			"B": b,
			"C": c,
		}
		output := []int{}
		pos := 0
		getLiteral := func(operand int) int64 {
			return int64(operand)
		}
		getCombo := func(operand int) int64 {
			if operand < 4 {
				return getLiteral(operand)
			}
			if operand == 4 {
				return registers["A"]
			}
			if operand == 5 {
				return registers["B"]
			}
			if operand == 6 {
				return registers["C"]
			}
			panic("Wrong combo operand")
		}
		for pos < len(nums) {
			opcode := nums[pos]
			operand := nums[pos+1]
			pos += 2
			if opcode == 0 {
				registers["A"] = registers["A"] / (1 << getCombo(operand))
			}
			if opcode == 1 {
				registers["B"] ^= getLiteral(operand)
			}
			if opcode == 2 {
				registers["B"] = getCombo(operand) % 8
			}
			if opcode == 3 {
				if registers["A"] != 0 {
					pos = int(getLiteral(operand))
				}
			}
			if opcode == 4 {
				registers["B"] ^= registers["C"]
			}
			if opcode == 5 {
				output = append(output, int(getCombo(operand)%8))
			}
			if opcode == 6 {
				registers["B"] = registers["A"] / (1 << getCombo(operand))
			}
			if opcode == 7 {
				registers["C"] = registers["A"] / (1 << getCombo(operand))
			}
			if n > 0 && len(output) >= n {
				break
			}
		}
		return output
	}
	p1 := run(int64(inRegisters["A"]), int64(inRegisters["B"]), int64(inRegisters["C"]), 0)
	p1s := make([]string, len(p1))
	for i, p := range p1 {
		p1s[i] = strconv.Itoa(p)
	}
	fmt.Println("Part 1:", strings.Join(p1s, ","))

	// B = A % 8
	// B = B ^ 3
	// C = A / (1 << B)
	// B = B ^ 5
	// A = A / 8
	// B = B ^ C
	// output <- B % 8
	// if A != 0 {
	//   goto begin
	// }

	//run2 := func(xx int64, n int) []int {
	//	a, b, c := xx, int64(0), int64(0)
	//	out := []int{}
	//	for {
	//		b = a % int64(8)
	//		b = b ^ int64(3)
	//		//c = a / (1 << b)
	//		c = a >> b
	//		b = b ^ int64(5)
	//		//a = a / 8
	//		a = a >> int64(3)
	//		b = b ^ c
	//		out = append(out, int(b%8))
	//		if n > 0 && len(out) >= n {
	//			break
	//		}
	//		if a == 0 {
	//			break
	//		}
	//	}
	//	return out
	//}
	test := func(masks map[int64]bool, shift int, a, b, c, d int) map[int64]bool {
		u := map[int64]bool{}
		for i := int64(0); i < int64(1 << 19); i++ {
			r1 := run(i, 0, 0, 4)
			//r1 := run2(i, 4)
			if len(r1) >= 4 && r1[0] == a && r1[1] == b && r1[2] == c && r1[3] == d {
				if len(masks) == 0 {
					u[i] = true
				}
				for k := range masks {
					ss := k >> shift
					if i&ss == ss {
						u[(i << shift) | k] = true
					}
				}
			}
		}
		return u
	}
	current := map[int64]bool{}
	for t := 0; t < len(nums)-4; t++ {
		current = test(current, 3 * t, nums[t], nums[t+1], nums[t+2], nums[t+3])
	}
	eq := func(a, b []int) bool {
		if len(a) != len(b) {
			return false
		}
		for i, x := range a {
			if b[i] != x {
				return false
			}
		}
		return true
	}
	res := []int64{}
	for k := range current {
		arr := run(k, int64(0), int64(0), 0)
		//arr := run2(k, 0)
		if eq(nums, arr) {
			res = append(res, k)
		}
	}
	sort.Slice(res, func(a, b int) bool {
		return res[a] < res[b]
	})
	fmt.Println("Part 2:", res[0])
}
