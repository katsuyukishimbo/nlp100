package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func Reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	if sc.Scan() {
		fmt.Println(Reverse(sc.Text()))
	}
}
