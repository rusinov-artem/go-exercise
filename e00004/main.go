package main

import "fmt"

func main() {
	k := 0
	go func() {
		k = 23
	}()
	k = 4
	fmt.Println(k)
}
