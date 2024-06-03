package main

import (
	"fmt"
	"time"
)

type Limiter struct {
	// 用来存储信号量
	Signals chan time.Time
	// 每秒里最多的并发数
	MaxPerPeriod int
}

func NewLimiter(maxPerPeriod int) *Limiter {
	limiter := &Limiter{
		Signals:      make(chan time.Time, maxPerPeriod),
		MaxPerPeriod: maxPerPeriod,
	}
	go func() {
		for t := range time.Tick(time.Second) {
			for i := 0; i < maxPerPeriod; i++ {
				limiter.Signals <- t
			}
		}
	}()
	return limiter
}

func (l *Limiter) Allow() {
	<-l.Signals
}

func main() {
	limiter := NewLimiter(2)
	requests := make(chan int, 10)
	for i := 0; i < 10; i++ {
		requests <- i
	}
	for req := range requests {
		limiter.Allow()
		fmt.Println("burstyRequests", req, time.Now())
	}

}
