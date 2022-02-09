package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack struct {
	items []int
}

func (s *Stack) Push(x int) {
	s.items = append(s.items, x)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1
	}
	var result int
	result, s.items = s.items[len(s.items)-1], s.items[:len(s.items)-1]
	return result
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Empty() bool {
	return len(s.items) == 0
}

func (s *Stack) Top() int {
	if len(s.items) == 0 {
		return -1
	}
	return s.items[len(s.items)-1]
}

func main() {
	rw := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	defer rw.Flush()
	var n int
	fmt.Fscan(rw, &n)
	stack := Stack{}
	
	for i := 0; i < n; i++ {
		var command string
		fmt.Fscan(rw, &command)
		switch command {
		case "push":
			var x int
			fmt.Fscan(rw, &x)
			stack.Push(x)
		case "pop":
			fmt.Println(stack.Pop())
		case "size":
			fmt.Println(stack.Size())
		case "empty":
			if stack.Empty() {
				fmt.Println(1)
			} else {
				fmt.Println(0)
			}
		case "top":
			fmt.Println(stack.Top())
		}
	}
}
