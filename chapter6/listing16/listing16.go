// This sample program demonstrates how to use a mutex
// to define critical sections of code that need synchronous
// access.用mutex定义一段需要同步访问代码临界区资源同步问题。
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// counter is a variable incremented by all goroutines.
	counter int

	// wg is used to wait for the program to finish.
	wg sync.WaitGroup

	// mutex is used to define a critical section of code.mutex用来定义一段代码的临界区。
	mutex sync.Mutex
)

// main is the entry point for all Go programs.
func main() {
	// Add a count of two, one for each goroutine.
	wg.Add(2)

	// Create two goroutines.
	go incCounter(1)
	go incCounter(2)

	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Printf("Final Counter: %d\n", counter)
}

// incCounter increments the package level Counter variable
// using the Mutex to synchronize and provide safe access.用互斥锁同步保证安全访问。
func incCounter(id int) {
	// Schedule the call to Done to tell main we are done.函数退出时调用Done（）通知main函数工作已经完成
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// Only allow one goroutine through this
		// critical section at a time.同一时间只允许一个goroutine进入这个临界区。
		mutex.Lock()
		{
			// Capture the value of counter.补货当前counter的值
			value := counter

			// Yield the thread and be placed back in queue.当前goroutine从线程退出，并放回到队列
			runtime.Gosched()

			// Increment our local value of counter.增减本地value值
			value++

			// Store the value back into counter.将value1保存回counter
			counter = value
		}
		mutex.Unlock()
		// Release the lock and allow any释放锁，
		// waiting goroutine through.允许正在等待的goroutine进入临界区
	}
}
