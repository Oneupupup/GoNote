package test6

import (
	"fmt"
	"sync"
	"time"
)

func TestSyn() {
	// 1. 互斥锁
	var mutex sync.Mutex
	var c int
	// WaitGroup 用于等待一组线程的结束
	var wg sync.WaitGroup

	prime := func(n int) {
		defer wg.Done()
		for i := 2; i < n; i++ {
			if n%i == 0 {
				return
			}
		}
		mutex.Lock()
		c++
		mutex.Unlock()
	}
	for i := 2; i < 10000; i++ {
		wg.Add(1)
		go prime(i)
	}
	wg.Wait()
	fmt.Println(c)
}

// 2. 条件变量
func TestSyn2() {
	// cond : 条件变量
	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)
	for i := 0; i < 10; i++ {
		go func(i int) {
			cond.L.Lock()
			cond.Wait()
			fmt.Println("协程", i, "被唤醒")
			cond.L.Unlock()
		}(i)
	}
	for i:=0;i<15;i++{
		time.Sleep(time.Second)
		fmt.Print(".") // 用于观察程序运行状态
		if i == 4 {
			fmt.Println()
			cond.Signal()
		}
		if i == 9 {
			fmt.Println()
			cond.Broadcast()
		}

	}
}

// once
func TestSyn3() {
	var once sync.Once
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			once.Do(func() {
				fmt.Println("只执行一次")
			})
		wg.Done()
		}()
	}
	wg.Wait()
}

// Map
func TestSyn4() {
	// map 是sync.Map类型，不是map类型
	var m sync.Map
	m.Store(1,100)
	m.Store(2,200)
	m.Store(3,300)
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key,value)
		return true
	})
}

// 读写互斥锁
func TestSyn5() {
	
}