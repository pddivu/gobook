package main

import "fmt"

func main() {
	slice := []string{"one", " ", "Three", " ", "Four"}
	fmt.Println(removespace(slice))
}
func removespace(rem []string) []string {
	i, count := 0, 0
	for _, r := range rem {
		if r == " " {
			copy(rem[i:], rem[i+1:])
			count++
		}
		i++
	}
	return rem[:(len(rem) - count)]
}
