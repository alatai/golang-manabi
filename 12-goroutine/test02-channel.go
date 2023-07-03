package main

import (
	"fmt"
	"time"
)

/* 两个goroutine之间的通信——channel */

func testChannel() {
	// 定义一个channel
	channel := make(chan int)

	go func() {
		defer fmt.Println("goroutine结束")
		fmt.Println("goroutine正在运行...")

		// 将123发送个给channel
		channel <- 123

	}()

	// 从channel中接收数据，并赋值给num
	num := <-channel
	fmt.Println("num = ", num)
	fmt.Println("main goroutine 结束")
}

// 带有缓冲的channel
func testChannel02() {
	// 当channel已经满，再向里面写数据就会阻塞
	// 当channel为空，从里面取数据也会阻塞
	c := make(chan int, 3)
	fmt.Println("len(c) = ", len(c), ", cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("子goroutine结束")

		for i := 0; i < 5; i++ {
			c <- i
			fmt.Println("goroutine正在运行，发送的元素 = ", i, ", len(c) = ", len(c), ", cap(c) = ", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 5; i++ {
		// 从c中接收数据，并赋值给num
		num := <-c
		fmt.Println("num = ", num)
	}

	fmt.Println("主goroutine结束")
}

// channel关闭
// channel不像文件一样需要经常去关闭，只有当确实没有任何发送数据了，或者想显示的结束range循环之类的，采取关闭channel
// 关闭channel后，无法向channel在发送数据（引发panic错误后导致接收立刻返回零值）
// 关闭channel后，可以继续从channel接收数据
// 对于nil channel，无论收发都会被阻塞
func testChannel03() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		// close可以关闭一个channel
		close(c)
	}()

	/*
		for {
			// ok如果为true表示channel没有关闭，如果为false表示channel已经关闭
			if data, ok := <-c; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
	*/

	// 利用range将上述for循环简写，使用range来迭代channel
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("Main Finished...")
}

// channel与select
// 单流程下一个goroutine只能监控一个channel的状态，select可以完成监控多个channel的状态
func testChannel04() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}

		quit <- 0
	}()

	// main goroutine
	fibonacci(c, quit)
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1

	for {
		select {
		// 如果c可写，则进入该case执行
		case c <- x:
			x = y
			y = x + y
		// 如果quit可读，则进入该case执行
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
