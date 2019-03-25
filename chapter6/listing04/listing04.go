// This sample program demonstrates how the goroutine scheduler
// will time slice goroutines on a single thread.展示goroutine调度器如何在单个线程上切分时间片的
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

// main is the entry point for all Go programs.
func main() {
	// Allocate 1 logical processors for the scheduler to use.
	runtime.GOMAXPROCS(1)

	// Add a count of two, one for each goroutine.
	wg.Add(2)

	// Create two goroutines.创建两个goroutine
	fmt.Println("Create Goroutines")
	go printPrime("A")
	go printPrime("B")

	// Wait for the goroutines to finish.等待goroutine结束
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}

// printPrime displays prime numbers for the first 5000 numbers.显示500以内的素数值
func printPrime(prefix string) {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}
