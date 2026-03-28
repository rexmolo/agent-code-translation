package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func FizzBuzz(n int) int {
	var ns []int
	for i := 0; i < n; i++ {
		if i%11 == 0 || i%13 == 0 {
			ns = append(ns, i)
		}
	}

	var sb strings.Builder
	for _, num := range ns {
		sb.WriteString(strconv.Itoa(num))
	}

	s := sb.String()
	ans := 0
	for _, c := range s {
		if c == '7' {
			ans++
		}
	}
	return ans
}

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	if !in.Scan() {
		return
	}
	n, _ := strconv.Atoi(in.Text())
	fmt.Println(FizzBuzz(n))
}