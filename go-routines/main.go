package main

import (
	"fmt"
	"time"
)

func greet(s string) {
	fmt.Println(s)
}

func slowGreet(s string) {
	time.Sleep(time.Second * 3)
	fmt.Println(s)
}

func main() {
	go greet("Hello")
	go greet("How are you?")
	go slowGreet("How... are... you...?")
	go greet("I am fine")
}
