package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//lines := make(map[string]bool)
	lines := map[string]bool{"Divya": true}
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		data := input.Text()
		if !lines[data] {
			lines[data] = true
			fmt.Println(data)
			fmt.Println(lines)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "deleteduplicate: %v\n", err)
		os.Exit(1)
	}
}
