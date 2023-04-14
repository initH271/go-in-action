package test_goroutine

import (
	"runtime"
	"sync"
	"time"
)

func StartTwoGoroutine() {
	start := time.Now()
	numCPU := 8
	runtime.GOMAXPROCS(numCPU) // 设置可同时执行的最大CPU数，默认为每一个协程分配一个逻辑处理器。 这里设置仅允许同时执行一个逻辑处理器
	var wait sync.WaitGroup    // 信号量，用于阻塞主线程
	wait.Add(8)

	go func() {

		defer wait.Done() // 协程完成后，调用Done()方法，阻塞计数器减-1
		for i := 0; i < 1000; i++ {
			println("goroutine A: ", i)
		}
	}()

	go func() {

		defer wait.Done()
		for i := 0; i < 1000; i++ {
			println("goroutine B: ", i)
		}
	}()

	go func() {

		defer wait.Done() // 协程完成后，调用Done()方法，阻塞计数器减-1
		for i := 0; i < 1000; i++ {
			println("goroutine C: ", i)
		}
	}()

	go func() {

		defer wait.Done()
		for i := 0; i < 1000; i++ {
			println("goroutine D: ", i)
		}
	}()

	go func() {

		defer wait.Done() // 协程完成后，调用Done()方法，阻塞计数器减-1
		for i := 0; i < 1000; i++ {
			println("goroutine E: ", i)
		}
	}()

	go func() {

		defer wait.Done()
		for i := 0; i < 1000; i++ {
			println("goroutine F: ", i)
		}
	}()

	go func() {

		defer wait.Done() // 协程完成后，调用Done()方法，阻塞计数器减-1
		for i := 0; i < 1000; i++ {
			println("goroutine G: ", i)
		}
	}()

	go func() {

		defer wait.Done()
		for i := 0; i < 1000; i++ {
			println("goroutine H: ", i)
		}
	}()
	wait.Wait() // 阻塞直到计数器为0
	println("计算机逻辑处理器数:", runtime.NumCPU(), ", 当前允许同时运行的逻辑处理器数:", numCPU, ", 当前存在的协程数:", runtime.NumGoroutine())
	println("消耗时间： ", time.Since(start).String())
}
