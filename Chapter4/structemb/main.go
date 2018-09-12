package main

import "fmt"

func main() {
	type point struct {
		x, y int
	}

	type circle struct {
		point
		radius int
	}

	type wheel struct {
		circle
		spokes int
	}
	var w wheel

	w = wheel{circle{point{8, 8}, 10}, 20}

	area := w.y * w.x
	fmt.Println(area)
}
