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
	arr := make([]int64, n)
	for i = 0; i < n; i++ {
		var a int64
		fmt.Fscan(in, &a)
		arr[i] = hash[a]
	}
	set := make(map[int64]int64)
	var minLen int64 = -1
	for i = 0; i < n; i++ {
		if _, ok := set[arr[i]]; ok {
			length := i - set[arr[i]]
			if minLen == -1 || length < minLen {
				minLen = length
			}
			set[arr[i]] = i
		} else {
			set[arr[i]] = i
		}
		if len(set) == int(n) {
			minLen = i + 1
			break
		}
	}
	fmt.Fprint(out, minLen)
}
