package main

import (
	"bufio"
	"fmt"
	"os"
)

type Person struct {
	height int
	weight int
}

func main() {
	rw := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	defer rw.Flush()

	var n int
	fmt.Fscan(rw, &n)

	people := make([]Person, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(rw, &people[i].height, &people[i].weight)
	}
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < n; j++ {
			if people[i].height < people[j].height && people[i].weight < people[j].weight {
				count++
			}
		}
		rank[i] = count + 1
	}
	for _, r := range rank {
		fmt.Printf("%d ", r)
	}
}
