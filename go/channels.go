package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{}) // 使用空结构体作为信号
	go func() {
		time.Sleep(time.Second * 3)
		// 执行一些任务...
		close(done) // 任务完成，关闭channel作为信号
	}()
	<-done // 等待任务完成

}

// worker 是执行任务的goroutine
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		// 模拟任务执行时间
		time.Sleep(time.Second)
		// 将结果发送到results channel
		results <- j
		fmt.Printf("Worker %d finished job %d\n", id, j)
	}
}
