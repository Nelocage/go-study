package main

import (
	"fmt"
	"time"
)

func channeltest() {
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}

//通道缓冲

func channelbuffer() {
	message := make(chan string, 2)

	message <- "buffed "
	message <- "channel"

	fmt.Println(<-message)
	fmt.Println(<-message)
}

func channel_syschronization() {
	done := make(chan bool, 1)
	go worker(done)

	<-done
	time.Now()
}
func worker(done chan bool) {
	fmt.Println("working... ")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
	time.Now()
}
func main() {
	channel_syschronization()
}
