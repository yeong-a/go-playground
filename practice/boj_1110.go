package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	rw := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	defer rw.Flush()
	var num, count int
	fmt.Fscan(rw, &num)
	n := num
	for {
		a := strconv.Itoa(n % 10)
		b := strconv.Itoa((n/10 + n%10) % 10)
		sum, _ := strconv.Atoi(a + b)
		count++
		if num == sum {
			break
		}
		n = sum
	}
	fmt.Fprintln(rw, count)
}
