package client

import "time"

type throttler struct {
	control chan struct{}
	timeout time.Duration
}

func newThrottler(requestLimit int, refresh time.Duration) *throttler {
	rpsControl := &throttler{
		control: make(chan struct{}, requestLimit),
		timeout: refresh,
	}

	go rpsControl.refresh()

	return rpsControl
}

func (t *throttler) allow() {
	<-t.control
}

func (t *throttler) refresh() {
	for {
		t.fill()

		time.Sleep(t.timeout)
	}
}

func (t *throttler) fill() {
	for i := 0; i < cap(t.control); i++ {
		select {
		case t.control <- struct{}{}:

		default:
			return
		}
	}
}
