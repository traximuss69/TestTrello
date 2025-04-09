package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if sign(a) == 1 && sign(b) == 1 {
		fmt.Println(1)
	} else {
		fmt.Println(-1)
	}
}
func sign(n int) int {
	n1 := n % 10
	n2 := (n / 10) % 10
	n3 := (n / 100) % 10
	n4 := (n / 1000) % 10
	n5 := (n / 10000) % 10
	n6 := (n / 100000) % 10
	if n1+n2+n3 == n4+n5+n6 {
		return 1
	} else {
		return -1
	}
}
