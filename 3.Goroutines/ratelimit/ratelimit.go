package ratelimit

import (
	"context"
	"errors"
	"time"
)

// Limiter is precise rate limiter with context support.
type Limiter struct {
	cnt      chan int
	d        time.Duration
	stopped  chan bool
	maxCount int
	timer    *time.Timer
}

var ErrStopped = errors.New("limiter stopped")

// NewLimiter returns limiter that throttles rate of successful Acquire() calls
// to maxSize events at any given interval.
func NewLimiter(maxCount int, interval time.Duration) *Limiter {
	ch1 := make(chan int, 1)
	ch1 <- 0
	ch2 := make(chan bool, 1)
	ch2 <- false
	var timer *time.Timer = time.NewTimer(interval)
	return &Limiter{ch1, interval, ch2, maxCount, timer}
}

func (l *Limiter) Acquire(ctx context.Context) error {
	flag := <-l.stopped
	l.stopped <- flag
	if flag {
		return ErrStopped
	}
	select {
	case cur := <-l.cnt:
		if cur == 0 {
			l.timer.Reset(l.d)
		}
		cur++
		if cur != l.maxCount {
			l.cnt <- cur
		}
	case <-ctx.Done():
		return ctx.Err()
	case <-l.timer.C:
		l.cnt = make(chan int, 1)
		l.cnt <- 0
	}
	return nil
}

func (l *Limiter) Stop() {
	<-l.stopped
	l.timer.Stop()
	l.stopped <- true
}
