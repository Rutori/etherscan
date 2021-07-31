package client

import "time"

// throttler slows down api requests to meet limitation
type throttler struct {
	control chan struct{}
	timeout time.Duration
}

// newThrottler creates a new throttler
func newThrottler(requestLimit int, refresh time.Duration) *throttler {
	rpsControl := &throttler{
		control: make(chan struct{}, requestLimit),
		timeout: refresh,
	}

	go rpsControl.refresh()

	return rpsControl
}

// allow spends a request ticket from available pool
func (t *throttler) allow() {
	<-t.control
}

// refresh constantly fills tickets on timeout
func (t *throttler) refresh() {
	for {
		t.fill()

		time.Sleep(t.timeout)
	}
}

// fill adds new tickets to the pool
func (t *throttler) fill() {
	for i := 0; i < cap(t.control); i++ {
		select {
		case t.control <- struct{}{}:

		default:
			return
		}
	}
}
