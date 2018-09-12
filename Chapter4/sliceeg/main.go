package main

import "fmt"

func main() {
	var slice []string
	slice = []string{"one", " ", "three"}
	fmt.Println(emptydata(slice))
	fmt.Println(nonemptydata(slice))
}

func emptydata(z []string) []string {

	i := 0
	for _, r := range z {
		if r != " " {
			z[i] = r
			i++
		}
	}

	return z[:i]
}

func nonemptydata(z []string) []string {
	s := z[:0]
	for _, r := range z {
		if r != " " {
			s = append(s, r)
		}
	}
	return s
}
