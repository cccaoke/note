package main

import (
	"fmt"
	"time"
)

func printOdd(ch chan struct{}) {
	for i := 1; i <= 100; i++ {
		if i&1 == 1 {
			fmt.Println("printer1 print: ", i)
		}
		ch <- struct{}{}
	}
}

func printEven(ch chan struct{}) {
	for i := 1; i <= 100; i++ {
		<-ch
		if i&1 == 0 {
			fmt.Println("printer2 print: ", i)
		}
	}
}

func main() {
	ch := make(chan struct{})
	go printOdd(ch)
	go printEven(ch)
	time.Sleep(3 * time.Second)
}
