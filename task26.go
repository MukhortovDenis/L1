package main

import (
	"fmt"
	"os"
	"strings"
)

func check(str string) bool {
	str = strings.ToLower(str)
	newrune := []rune(str)
	for i := range newrune {
		for j := range newrune {
			if i == j{
				j++
			}else if newrune[i] == newrune[j]{
				return false
			}
		}
	}
	return true
}

func main() {
	var str string
	fmt.Fscan(os.Stdin, &str)
	fmt.Print(check(str))

}
