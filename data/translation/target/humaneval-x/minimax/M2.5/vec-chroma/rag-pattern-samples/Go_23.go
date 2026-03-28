package main

import "fmt"

func Strlen(str string) int {
	return len(str)
}

func main() {
	fmt.Println(Strlen(""))
	fmt.Println(Strlen("abc"))
}
