// This sample program demonstrates how to use an unbuffered如何用无缓冲通道来模拟2个goroutine间的网球比赛。
// channel to simulate a game of tennis between two goroutines.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

// main is the entry point for all Go programs.
func main() {
	// Create an unbuffered channel.创建一个无缓冲通道
	court := make(chan int)

	// Add a count of two, one for each goroutine.技术器2，表示等待两个goroutine
	wg.Add(2)

	// Launch two players.启动两个选手
	go player("Nadal", court)
	go player("Djokovic", court)

	// Start the set.发球
	court <- 1

	// Wait for the game to finish.等待游戏结束。
	wg.Wait()
}

// player simulates a person playing the game of tennis.模拟一个选手在打网球。
func player(name string, court chan int) {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()

	for {
		// Wait for the ball to be hit back to us.等待球被击打过来
		ball, ok := <-court//会锁住goroutine，直到有数据发送到通道里。
		if !ok {
			// If the channel was closed we won.如果通道关了，我们就赢了。
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// Pick a random number and see if we miss the ball.选随机数，然后用这个随机数判断我们是否丢球。
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			// Close the channel to signal we lost.关闭通道，表示我们输了
			close(court)
			return
		}

		// Display and then increment the hit count by one.显示击球数，并将击球数加一
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// Hit the ball back to the opposing player.将球打向对手
		court <- ball//ball重新放入通道，这个时候两个goroutine都会被锁住，直到交换完成。
	}
}
