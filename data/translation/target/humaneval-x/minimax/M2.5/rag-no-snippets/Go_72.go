package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func WillItFly(q []int, w int) bool {
	// Check if sum of elements is greater than w
	sum := 0
	for _, v := range q {
		sum += v
	}
	if sum > w {
		return false
	}

	// Check if it's a palindrome (balanced)
	i, j := 0, len(q)-1
	for i < j {
		if q[i] != q[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func main() {
	// Read input from stdin (format: first line contains q as space-separated integers, second line contains w)
	scanner := bufio.NewScanner(os.Stdin)

	var q []int
	w := 0

	if scanner.Scan() {
		line1 := scanner.Text()
		parts := strings.Split(line1, " ")
		for _, p := range parts {
			p = strings.TrimSpace(p)
			if p == "" {
				continue
			}
			n, err := strconv.Atoi(p)
			if err != nil {
				fmt.Println("Error parsing input")
				return
			}
			q = append(q, n)
		}
	}

	if scanner.Scan() {
		line2 := strings.TrimSpace(scanner.Text()
		)
		wVal, err := strconv.Atoi(line2)
		if err != nil {
			fmt.Println("Error parsing w")
			return
		}
		w = wVal
	}

	result := WillItFly(q, w)
	fmt.Println(result)
}
