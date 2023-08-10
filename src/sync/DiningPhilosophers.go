package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//互斥条件：资源是独占的且排他使用，进程互斥使用资源，即任意时刻一个资源只能给一个进程使用，其他进程若申请一个资源，而该资源被另一进程占有时，则申请者等待直到资源被占有者释放。
//不可剥夺条件：进程所获得的资源在未使用完毕之前，不被其他进程强行剥夺，而只能由获得该资源的进程资源释放。
//请求和保持条件：进程每次申请它所需要的一部分资源，在申请新的资源的同时，继续占用已分配到的资源。
//循环等待条件：在发生死锁时必然存在一个进程等待队列{P1,P2,…,Pn},其中P1等待P2占有的资源，P2等待P3占有的资源，…，Pn等待P1占有的资源，形成一个进程等待环路，环路中每一个进程所占有的资源同时被另一个申请，也就是前一个进程占有后一个进程所深情地资源。

//在哲学家问题中，互斥条件不能被替换，不可剥夺条件也不能打破，因此如果想打破死锁，需要在其他两个条件上考虑
//请求和保持：设置获取锁的超时时间，这个在通过channel进行锁控制的时候可以使用
//循环等待：
//	偶数先右后左，奇数先左后右，这样就不会出现保持的情况
//  控制每次发起锁请求的哲学家人数，同时至多右4个人可以用餐，用餐完成的时候再归还一个令牌

func deadlock() {
	var forks []sync.Mutex = make([]sync.Mutex, 5)
	var waitGroup sync.WaitGroup
	waitGroup.Add(5)
	for i := 0; i < 5; i++ {
		go func(phi int) {
			fmt.Println("phi:", phi, "start")
			for count := 0; count < 100; count++ {
				//think
				rand.Seed(0)
				time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
				forks[phi].Lock()
				fmt.Println("phi:", phi, "take left fork")
				forks[(phi+1)%5].Lock()
				fmt.Println("phi:", phi, "take right fork")
				//eat
				fmt.Println("phi:", phi, "eating!")
				time.Sleep(time.Second)
				forks[phi].Unlock()
				fmt.Println("phi:", phi, "put left fork")
				forks[(phi+1)%5].Unlock()
				fmt.Println("phi:", phi, "put right fork")
			}
		}(i)
		waitGroup.Done()
	}
	waitGroup.Wait()
	for {
	}
}

//偶数先右后左，奇数先左后右，这样就不会出现保持的情况
func leftright() {
	var forks []sync.Mutex = make([]sync.Mutex, 5)
	var waitGroup sync.WaitGroup
	waitGroup.Add(5)
	for i := 0; i < 5; i++ {
		go func(phi int) {
			fmt.Println("phi:", phi, "start")
			for count := 0; count < 100; count++ {
				//think
				rand.Seed(0)
				time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
				var first, second int
				if phi%2 == 0 {
					first, second = phi, (phi+1)%5
				} else {
					first, second = (phi+1)%5, phi
				}
				forks[first].Lock()
				fmt.Println("phi:", phi, "take left fork")
				forks[second].Lock()
				fmt.Println("phi:", phi, "take right fork")
				//eat
				fmt.Println("phi:", phi, "eating!")
				time.Sleep(time.Second)
				forks[first].Unlock()
				fmt.Println("phi:", phi, "put left fork")
				forks[second].Unlock()
				fmt.Println("phi:", phi, "put right fork")
			}
		}(i)
		waitGroup.Done()
	}
	waitGroup.Wait()
	for {
	}
}

//控制每次发起锁请求的哲学家人数，同时至多右4个人可以用餐，用餐完成的时候再归还一个令牌
func maxCount() {
	tickets := make(chan interface{}, 4)
	for i := 0; i < 4; i++ {
		tickets <- i
	}
	var forks []sync.Mutex = make([]sync.Mutex, 5)
	var waitGroup sync.WaitGroup
	waitGroup.Add(5)
	for i := 0; i < 5; i++ {
		go func(phi int) {
			fmt.Println("phi:", phi, "start")
			for count := 0; count < 100; count++ {
				//think
				rand.Seed(0)
				time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
				ticket := <-tickets
				fmt.Println("phi:", phi, "get ticket", ticket)
				forks[phi].Lock()
				fmt.Println("phi:", phi, "take left fork")
				forks[(phi+1)%5].Lock()
				fmt.Println("phi:", phi, "take right fork")
				//eat
				fmt.Println("phi:", phi, "eating!")
				time.Sleep(time.Second)
				forks[phi].Unlock()
				fmt.Println("phi:", phi, "put left fork")
				forks[(phi+1)%5].Unlock()
				fmt.Println("phi:", phi, "put right fork")
				tickets <- ticket
			}
		}(i)
		waitGroup.Done()
	}
	waitGroup.Wait()
	for {
	}
}

//设置获取锁的超时时间，这个在通过channel进行锁控制的时候可以使用
func timeout() {
	var forks = make([]chan interface{}, 5)
	for i := 0; i < 5; i++ {
		forks[i] = make(chan interface{}, 1)
		forks[i] <- i
	}
	var waitGroup sync.WaitGroup
	waitGroup.Add(5)
	for i := 0; i < 5; i++ {
		go func(phi int) {
			fmt.Println("phi:", phi, "start")
			for count := 0; count < 100; count++ {
				//think
				select {
				case left := <-forks[phi]:
					{
						fmt.Println("phi:", phi, "take left fork")
						select {
						case right := <-forks[(phi+1)%5]:
							{
								fmt.Println("phi:", phi, "eating!")
								forks[(phi+1)%5] <- right
							}
						case <-time.After(time.Second):
							{
								fmt.Println("phi:", phi, "take right fork timeout")
							}
						}
						forks[phi] <- left
					}
				case <-time.After(time.Second):
					{
						fmt.Println("phi:", phi, "take left fork timeout")
					}
				}
			}
		}(i)
		waitGroup.Done()
	}

	waitGroup.Wait()
	for {
	}
}

func main() {
	timeout()
}
