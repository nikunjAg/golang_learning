package main

import (
	"fmt"
	"time"
)

func greet(s string, done_channel chan bool) {
	fmt.Println(s)

	if done_channel != nil {
		done_channel <- true
	}
}

func slowGreet(s string, done_channel chan bool) {
	time.Sleep(time.Second * 3)
	fmt.Println(s)

	done_channel <- true
	close(done_channel)
}

func main() {
	var doneChannel = make(chan bool)

	go greet("Hello", doneChannel)
	go greet("How are you?", doneChannel)
	go slowGreet("How... are... you...?", doneChannel)
	go greet("I am fine", doneChannel)
	<-doneChannel
	go slowGreet("How... are... you... (2)?", doneChannel)
	go greet("Done with slow Greet?", doneChannel)
	go greet("Last!!", doneChannel)
	<-doneChannel
	<-doneChannel
	<-doneChannel
	<-doneChannel
	<-doneChannel
	<-doneChannel
}
