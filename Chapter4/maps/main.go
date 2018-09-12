package main

import (
	"fmt"
	"sort"
)

func main() {
	map1 := map[string]int{
		"Divya":  31,
		"Eshaan": 7,
	}

	map2 := make(map[string]bool)

	map2["Divya"] = true
	fmt.Println(map2["Divya"])

	fmt.Println(map1)
	delete(map1, "Eshaan")
	fmt.Printf("%q\n", map1)

	map1["Aroop"] = 39
	map1["Eshaan"] = 7

	name := []string{}
	for n := range map1 {
		name = append(name, n)

	}
	sort.Strings(name)

	fmt.Println(name)
	for _, n := range name {
		fmt.Printf("%s: %d\n", n, map1[n])
	}

	var map3 map[string]int
	fmt.Println(map3 == nil)
	fmt.Println(len(map3))
	//map3["Divya"] = 3 -- error

	//allocate map before storing data. Cant store to the nill
	map3 = map[string]int{
		"Divya": 34}

	fmt.Println(map3)

}
