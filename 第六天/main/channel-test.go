package main

import "fmt"

func sum(s []int, c chan int) { //信道变量也需要命名

	sum := 0
	for _, v := range s {
		sum += v

	}
	c <- sum

}

/*
默认情况下，发送和接收操作在另一端准备好之前都会阻塞。这使得 Go 程可以在没有显式的锁或竞态变量的情况下进行同步。

以下示例对切片中的数进行求和，将任务分配给两个 Go 程。一旦两个 Go 程完成了它们的计算，它就能算出最终的结果。
*/
func main() {
	s := []int{2, 3, 4534, 34, 453, 52345, 2345325, 445}

	c := make(chan int)
	go sum(s[:len(s)/2], c) //为什么没用信道缓存也可以
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}
