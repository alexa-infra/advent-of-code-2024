package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	list1, list2 := []int{}, []int{}
	re := regexp.MustCompile(`([0-9]+)`)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		match := re.FindAllString(line, -1)
		a, _ := strconv.Atoi(match[0])
		b, _ := strconv.Atoi(match[1])
		list1 = append(list1, a)
		list2 = append(list2, b)
	}
	sort.Ints(list1)
	sort.Ints(list2)
	s1 := 0
	for i, p1 := range list1 {
		p2 := list2[i]
		s1 += abs(p1 - p2)
	}
	fmt.Println("Part 1:", s1)
	s2 := 0
	f := map[int]int{}
	for _, b := range list2 {
		f[b] += 1
	}
	for _, p1 := range list1 {
		s2 += f[p1] * p1
	}
	fmt.Println("Part 2:", s2)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
