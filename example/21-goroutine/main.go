package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println("hello goroutine : " + fmt.Sprint(i))
}

func main() {
	for i := 0; i < 5; i++ {
		go func(j int) {
			hello(j)
		}(i)
	}
	time.Sleep(time.Second)

	// output:
	// hello goroutine : 4
	// hello goroutine : 0
	// hello goroutine : 1
	// hello goroutine : 2
	// hello goroutine : 3

	// tips:
	// goroutinue 调用匿名函数
	// go func( 参数列表 ){
	//	 函数体
	// }( 调用参数列表 )
}
