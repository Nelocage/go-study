package main

import (
	"fmt"
	"time"
)

func test01() {
	for a := 0; a < 10; a++ {
		fmt.Println("另一个协程")
	}

}

func test02() {
	go test01()
	time.Sleep(time.Second * 5)
	fmt.Println("主协程")

}

func main() {
	test02()
}
