package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Solve(N int) string {
	str := strconv.Itoa(N)
	sum := 0
	for _, c := range str {
		sum += int(c - '0')
	}
	return strconv.FormatInt(int64(sum), 2)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var N int
	fmt.Fscan(reader, &N)
	fmt.Println(Solve(N))
}
