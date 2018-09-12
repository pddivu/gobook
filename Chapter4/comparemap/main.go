package main

import "fmt"

func main() {

	map1 := map[string]int{
		"Divya":  34,
		"Eshaan": 7,
	}
	map2 := map[string]int{
		"Divya":  34,
		"Eshaan": 7,
	}

	same := compare(map1, map2)
	fmt.Println(same)
}

/*
map1:= make(map[string]int)
func k(list []string) string  { return fmt.Sprintf("%q", list) }
func add(list []string)       { map1[k(list)]++ }
func Count(list []string) int { return map1[k(list)] }
*/
func compare(x, y map[string]int) bool {

	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		// checking if the key is not akey of y ,then value willbe 0
		if yv, ok := y[k]; !ok || xv != yv {
			return false
		}
	}
	return true
}
