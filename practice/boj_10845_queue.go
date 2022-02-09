// hallazzang이 굉장히 C스러운(객체가 없이 데이터와 메소드의 조합) 코드라고 했다.
// 스택은 Go스럽게(Receiver를 사용해서) 짜봐야겠다.


package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rw := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	defer rw.Flush()
	var n int
	fmt.Fscan(rw, &n)
	var queue []int
	for i := 0; i < n; i++ {
		var command string
		fmt.Fscan(rw, &command)
		switch command {
		case "push":
			var x int
			fmt.Fscan(rw, &x)
			push(&queue, x)
		case "pop":
			pop(&queue)
		case "size":
			size(queue)
		case "empty":
			empty(queue)
		case "front":
			front(queue)
		case "back":
			back(queue)
		}
	}
}

func push(queue *[]int, x int) {
	*queue = append(*queue, x)
}

func pop(queue *[]int) {
	if len(*queue) == 0 {
		fmt.Println(-1)
		return
	}
	fmt.Println((*queue)[0])
	if len(*queue) > 1 {
		*queue = (*queue)[1:]
		return
	}
	*queue = []int{}
}

func size(queue []int) {
	fmt.Println(len(queue))
}

func empty(queue []int) {
	if len(queue) == 0 {
		fmt.Println(1)
		return
	}
	fmt.Println(0)
}

func front(queue []int) {
	if len(queue) == 0 {
		fmt.Println(-1)
		return
	}
	fmt.Println(queue[0])
}

func back(queue []int) {
	if len(queue) == 0 {
		fmt.Println(-1)
		return
	}
	fmt.Println(queue[len(queue)-1])
}
