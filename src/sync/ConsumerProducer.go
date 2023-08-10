package main

import (
	"fmt"
	"math/rand"
)

func main() {
	channel := make(chan int, 10)
	go func() {
		rand.Seed(0)
		for {
			channel <- rand.Intn(10)
		}
	}()
	go func() {
		for {
			target := <-channel
			fmt.Println(target)
		}
	}()
	for {
	}
}
