package main

import (
	"bufio"
	"fmt"
	"os"
)

var cables []int64
var k int

func main() {
	rw := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	defer rw.Flush()

	var n int
	fmt.Fscan(rw, &k)
	fmt.Fscan(rw, &n)

	sum := int64(0)
	cables = make([]int64, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(rw, &cables[i])
		sum += cables[i]
	}
	min := int64(1)
	max := sum / int64(n)
	result := int64(0)
	for min <= max {
		mid := (max + min) / 2
		if isPossible(n, mid) {
			result = mid
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	fmt.Println(result)
}

func isPossible(n int, length int64) bool {
	count := 0
	for i := 0; i < k; i++ {
		cable := cables[i]
		for cable >= length {
			cable -= length
			count++
		}
	}
	return count >= n
}
