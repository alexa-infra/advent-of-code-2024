package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	text := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		text = append(text, line)
	}
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	r := int64(0)
	for _, line := range text {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			r += int64(a * b)
		}
	}
	fmt.Println("Part 1:", r)
	re1 := regexp.MustCompile(`(do\(\))|(don't\(\))|(mul\((\d{1,3}),(\d{1,3})\))`)
	r = int64(0)
	enabled := true
	for _, line := range text {
		matches := re1.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if match[0] == "do()" {
				enabled = true
			} else if match[0] == "don't()" {
				enabled = false
			} else if enabled {
				a, _ := strconv.Atoi(match[4])
				b, _ := strconv.Atoi(match[5])
				r += int64(a * b)
			}
		}
	}
	fmt.Println("Part 2:", r)
}
