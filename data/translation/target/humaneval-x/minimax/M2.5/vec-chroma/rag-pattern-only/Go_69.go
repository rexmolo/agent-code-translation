package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Search(lst []int) int {
	if len(lst) == 0 {
		return -1
	}

	maxVal := slices.Max(lst)

	frq := make([]int, maxVal+1)
	for _, i := range lst {
		frq[i]++
	}

	ans := -1
	for i := 1; i < len(frq); i++ {
		if frq[i] >= i {
			ans = i
		}
	}

	return ans
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	// Parse input like "[4, 1, 2, 2, 3, 1]"
	line = strings.Trim(line, "[]")
	parts := strings.Split(line, ",")

	var lst []int
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		v, _ := strconv.Atoi(p)
		lst = append(lst, v)
	}

	result := Search(lst)
	fmt.Println(result)
}
