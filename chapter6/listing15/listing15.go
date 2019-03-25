// This sample program demonstrates how to use the atomic
// package functions Store and Load to provide safe access
// to numeric types.用Store和Load保证数值类型的安全访问
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	// shutdown is a flag to alert running goroutines to shutdown.
	//shutdown 是通知正在执行的goroutine停止工作的标志。
	shutdown int64

	// wg is used to wait for the program to finish.wg用来等待程序结束
	wg sync.WaitGroup
)

// main is the entry point for all Go programs.
func main() {
	// Add a count of two, one for each goroutine.
	wg.Add(2)

	// Create two goroutines.
	go doWork("A")
	go doWork("B")

	// Give the goroutines time to run.
	time.Sleep(1 * time.Second)

	// Safely flag it is time to shutdown.安全的设置shutdown标志
	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)

	// Wait for the goroutines to finish.
	wg.Wait()
}

// doWork simulates a goroutine performing work and
// checking the Shutdown flag to terminate early.
func doWork(name string) {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()

	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)

		// Do we need to shutdown.程序该结束了，用Load安全的读取数值。
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}
