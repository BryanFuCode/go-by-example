package main

import "fmt"

func CalSquare() {
	src := make(chan int)
	dest := make(chan int, 2)
	go func() {
		defer close(src)
		for i := 0; i < 10; i++ {
			src <- i
		}
	}()
	go func() {
		defer close(dest)
		for i := range src {
			dest <- i * i
		}
	}()
	for i := range dest {
		fmt.Println(i)
	}

}

func main() {
	CalSquare()
	// output:
	// 0
	// 1
	// 4
	// 9
	// 16
	// 25
	// 36
	// 49
	// 64
	// 81
}
