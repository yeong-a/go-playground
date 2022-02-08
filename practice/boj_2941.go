package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	rw := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	defer rw.Flush()
	var str string
	fmt.Fscan(rw, &str)
	words := [8]string{"c=", "c-", "dz=", "d-", "lj", "nj", "s=", "z="}
	for _, word := range words {
		str = strings.ReplaceAll(str, word, "*")
	}
	fmt.Fprint(rw, len(str))
}
