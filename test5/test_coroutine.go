package test5

import (
	"fmt"
	"sync"
)

var count int = 0
var lock sync.Mutex

func PrimeNum(n int) {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return
		}
	}
	fmt.Printf("%d is prime num\n", n)
	lock.Lock() // 加锁
	count++
	lock.Unlock() // 解锁
}

// 协程
func TestCoroutine1() {
	for i := 2; i < 100001; i++ {
		go PrimeNum(i)
	}
	// 等待输入，否则主线程退出，协程也会退出
	var input string
	fmt.Scanln(&input)
	fmt.Printf("count = %d\n", count)
}

// 协程 + channel
func PushNum(n int, ch chan int) {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return
		}
	}
	ch <- n
}

func TestCoroutine2() {
	var ch_ chan int = make(chan int, 1000)
	for i := 2; i < 100001; i++ {
		go PushNum(i, ch_)
	}
print:
	for {
		select {
		case num := <-ch_:
			fmt.Printf("%d is prime num\n", num)
		default:
			fmt.Printf("default\n")
			break print
		}
	}
}
