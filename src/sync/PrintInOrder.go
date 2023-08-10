package main

import (
	"fmt"
	"sync"
)

//顺序打印123，有两种解法，
//1. 每个线程都创建一个channel，然后在完成的时候通过channel推送下一个线程
//2. 每个线程都抢锁，抢到之后如果轮到自己就输出数据，然后更新全局变量
func channel() {
	count := 5
	var channels = make([]chan int, count)
	for i := 0; i < count; i++ {
		channels[i] = make(chan int, 1)
	}
	channels[0] <- 0
	for i := 0; i < count; i++ {
		go func(index int) {
			for m := 0; m < 100; m++ {
				target := <-channels[index]
				next := (index + 1) % count
				fmt.Println("thread", index, " ", target, "next", next)
				channels[next] <- target + 1
			}
		}(i)
	}
	for {
	}
}

func main() {
	count := 5
	total := 0
	var mutex sync.Mutex
	for i := 0; i < count; i++ {
		go func(index int) {
			for m := 0; m < 100; m++ {
				mutex.Lock()
				if total%count == index {
					fmt.Println("thread", index, total)
					total++
				}
				mutex.Unlock()
			}
		}(i)
	}
	for {
	}
}
