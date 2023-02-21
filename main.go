package main

import (
	"fmt"
	"sort"
)

func main() {
	planet := make(map[int]string)
	planet[10460] = "Датомир"
	planet[10465] = "Татуин"
	planet[12120] = "Набу"
	planet[12240] = "Корусант"
	planet[12500] = "Альдераан"
	printSortedMap(planet)
}

func printSortedMap(m map[int]string) {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Printf("%d: %s\n", k, m[k])
	}
}
