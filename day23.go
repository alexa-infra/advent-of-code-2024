package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

func main() {
	graph := map[string][]string{}
	connect := func(a, b string) {
		arr, ok := graph[a]
		if !ok {
			arr = []string{}
		}
		graph[a] = append(arr, b)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "-")
		connect(parts[0], parts[1])
		connect(parts[1], parts[0])
	}
	uniq := map[[3]string]bool{}
	for k1, arr := range graph {
		for _, k2 := range arr {
			for _, k3 := range graph[k2] {
				for _, k4 := range graph[k3] {
					if k1 == k4 {
						set := [3]string{k1, k2, k3}
						sort.Strings(set[:])
						uniq[set] = true
					}
				}
			}
		}
	}
	p1 := 0
	for k := range uniq {
		for i := 0; i < 3; i++ {
			if strings.HasPrefix(k[i], "t") {
				p1 += 1
				break
			}
		}
	}
	fmt.Println("Part 1:", p1)

	names := []string{}
	for name := range graph {
		names = append(names, name)
	}
	largestGroup := []string{}
	for len(names) > 0 {
		group := []string{}
		for {
			updated := false
			for idx, current := range names {
				if len(group) == 0 {
					group = append(group, current)
					names[idx] = ""
					updated = true
				} else {
					connected := graph[current]
					groupConnected := true
					for _, node := range group {
						if !slices.Contains(connected, node) {
							groupConnected = false
							break
						}
					}
					if groupConnected {
						group = append(group, current)
						names[idx] = ""
						updated = true
					}
				}
			}
			if updated {
				unused := []string{}
				for _, current := range names {
					if current != "" {
						unused = append(unused, current)
					}
				}
				names = unused
			} else {
				break
			}
		}
		if len(group) > len(largestGroup) {
			largestGroup = group
		}
	}
	slices.Sort(largestGroup)
	p2 := strings.Join(largestGroup, ",")
	fmt.Println("Part 2:", p2)
}
