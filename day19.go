package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ptree struct {
	next map[byte]*ptree
	end  bool
}

func newPTree() *ptree {
	m := map[byte]*ptree{}
	return &ptree{m, false}
}

func (tree *ptree) addPrefix(p string) {
	if len(p) == 0 {
		tree.end = true
		return
	}
	first := p[0]
	t, ok := tree.next[first]
	if !ok {
		t = newPTree()
		tree.next[first] = t
	}
	t.addPrefix(p[1:])
}

func (tree *ptree) findPrefix(p string) []string {
	found := []string{}
	var findNext func(int, *ptree)
	findNext = func(i int, t *ptree) {
		if t.end {
			found = append(found, p[:i])
		}
		if i < len(p) {
			ch := p[i]
			t, ok := t.next[ch]
			if ok {
				findNext(i+1, t)
			}
		}
	}
	findNext(0, tree)
	return found
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), ", ")
	scanner.Scan()
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	prefixes := newPTree()
	for _, part := range parts {
		prefixes.addPrefix(part)
	}
	cache := map[string]int{}
	var match func(string) int
	match = func(p string) int {
		if len(p) == 0 {
			return 1
		}
		if v, ok := cache[p]; ok {
			return v
		}
		n := 0
		pp := prefixes.findPrefix(p)
		for _, x := range pp {
			n += match(p[len(x):])
		}
		cache[p] = n
		return n
	}
	p1, p2 := 0, 0
	for _, line := range lines {
		n := match(line)
		if n > 0 {
			p1 += 1
			p2 += n
		}
	}
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
