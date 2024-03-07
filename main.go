package main

import "fmt"

func main() {
	var a = 10
	defer func() {
		fmt.Println(a)
	}()
	a = 15
}
