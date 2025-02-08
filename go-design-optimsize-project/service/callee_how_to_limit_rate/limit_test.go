package callee_how_to_limit_rate

import (
	"testing"
	"time"
)

func TestLimitCheck(t *testing.T) {
	l := NewLimiter(3, 1)
	ch := make(chan int, 1000)
	signalStop := make(chan struct{}, 1)
	go func() {
		for {
			select {
			case <-signalStop:
				t.Logf("check routinue is close by caller.")
				return
			case data, ok := <-ch:
				if !ok {
					t.Logf("receive data chan.")
					return
				}
				if l.Check() == false {
					t.Logf("sender is to fast, data: %v", data)
				} else {

				}
			}
		}
	}()

	for i := 0; i < 20; i++ {
		ch <- i
		time.Sleep(1000 * time.Millisecond)
	}

	t.Logf("begin to close test.")
	close(signalStop)

	time.Sleep(10 * time.Second)

}
