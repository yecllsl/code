// This sample program demonstrates how to create goroutines and
// how the scheduler behaves.如何创建goroutine，以及调度器的行为
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// main is the entry point for all Go programs.
func main() {
	// Allocate 1 logical processor for the scheduler to use.分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(1)

	// wg is used to wait for the program to finish.wg用来等待程序完成
	// Add a count of two, one for each goroutine.计数器加2，表示要等待2个goroutine
	var wg sync.WaitGroup//计数信号量
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// Declare an anonymous function and create a goroutine.声明一个匿名函数，并创建goroutine
	go func() {
		// Schedule the call to Done to tell main we are done.在函数退出时调用wg.Done，来通知main函数工作已经完成。
		defer wg.Done()

		// Display the alphabet three times
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()//最后（）说明匿名函数执行

	// Declare an anonymous function and create a goroutine.
	go func() {
		// Schedule the call to Done to tell main we are done.
		defer wg.Done()

		// Display the alphabet three times
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
