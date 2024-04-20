package main

import "fmt"

func evenOdd(num int) string {
	if num%2 == 0 {
		return "even"
	}
	return "odd"
}

func main() {
	fmt.Println(evenOdd(4))
}
