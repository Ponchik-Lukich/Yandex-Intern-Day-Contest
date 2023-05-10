package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func contains(s []string, r string) bool {
	for _, a := range s {
		if a == r {
			return true
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	vowels := []string{"a", "e", "i", "o", "u", "y"}
	for i := 0; i < t; i++ {
		scanner.Scan()
		word := scanner.Text()
		lastTwo := word[len(word)-2:]
		last := word[len(word)-1:]
		if lastTwo == "sh" || lastTwo == "ch" || last == "s" || last == "x" || last == "z" {
			fmt.Println(scanner.Text() + "es")
		} else if last == "y" {
			if !contains(vowels, word[len(word)-2:len(word)-1]) {
				fmt.Println(scanner.Text()[:len(scanner.Text())-1] + "ies")
			} else {
				fmt.Println(scanner.Text() + "s")
			}
		} else {
			fmt.Println(scanner.Text() + "s")
		}
	}
}
