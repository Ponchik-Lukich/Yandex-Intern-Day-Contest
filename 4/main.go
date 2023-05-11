package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n int64
	var i int64
	fmt.Fscan(in, &n)
	hash := make(map[int64]int64)
	for i = 0; i < n; i++ {
		var a, b int64
		fmt.Fscan(in, &a, &b)
		hash[a] = b
	}
	var minLen int64 = -1
	set := make(map[int64]int64)
	for i = 0; i < n; i++ {
		var a int64
		fmt.Fscan(in, &a)
		if _, ok := set[hash[a]]; ok {
			length := i - set[hash[a]]
			if minLen == -1 || length < minLen {
				minLen = length
			}
			set[hash[a]] = i
		} else {
			set[hash[a]] = i
		}
		if len(set) == int(n) {
			minLen = i + 1
			break
		}
	}
	fmt.Fprint(out, minLen)
}
