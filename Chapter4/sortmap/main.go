package main

import (
	"fmt"
	"sort"
)

func main() {

	group := map[string]int{
		"Divya":  34,
		"Eshaan": 7,
		"Aroop":  39,
	}

	var names []string

	for name := range group {
		names = append(names, name)

	}
	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("%s :%d\n", name, group[name])
	}

}
