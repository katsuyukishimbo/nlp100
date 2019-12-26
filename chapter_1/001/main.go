package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func appendEven(input string) string {
	var str []rune
	for i, v := range input {
		if i%2 == 0 {
			str = append(str, v)
		}
	}
	return string(str)
}

func bytesEven(input string) string {
	var buf bytes.Buffer
	r := []rune(input)
	for i := 0; i < len(r); i++ {
		if i%2 == 1 {
			continue
		}
		buf.WriteRune(r[i])
	}
	return buf.String()
}

func main() {
	if sc.Scan() {
		fmt.Println(bytesEven(sc.Text()))
	}
}
