// This sample program demonstrates how to use an unbuffered
// channel to simulate a relay race between four goroutines.用goroutine模拟4个人的接力比赛
package main

import (
	"fmt"
	"sync"
	"time"
)

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

// main is the entry point for all Go programs.
func main() {
	// Create an unbuffered channel.创建一个无缓冲的通道，"接力棒"
	baton := make(chan int)

	// Add a count of one for the last runner.为最后一个跑步者将计数加1
	wg.Add(1)

	// First runner to his mark.第一个跑步者持有接力棒
	go Runner(baton)

	// Start the race.开始比赛
	baton <- 1

	// Wait for the race to finish.等待比赛结束。
	wg.Wait()
}

// Runner simulates a person running in the relay race.模拟接力比赛中的第一个跑者。
func Runner(baton chan int) {
	var newRunner int

	// Wait to receive the baton.等待接力棒
	runner := <-baton

	// Start running around the track.开始跑步
	fmt.Printf("Runner %d Running With Baton\n", runner)

	// New runner to the line.创建下一位跑者。
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		go Runner(baton)
	}

	// Running around the track.围着跑道跑。
	time.Sleep(100 * time.Millisecond)

	// Is the race over.比赛是否结束？
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		wg.Done()
		return
	}

	// Exchange the baton for the next runner.将接力棒递给下一位跑者。
	fmt.Printf("Runner %d Exchange With Runner %d\n",
		runner,
		newRunner)

	baton <- newRunner
}
