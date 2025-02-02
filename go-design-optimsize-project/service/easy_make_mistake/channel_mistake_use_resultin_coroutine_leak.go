package easy_make_mistake

import (
	"fmt"
	"runtime"
	"time"
)

// CoroutineLeakDemoImpl 错误使用 channel，导致协程泄露场景演示
func CoroutineLeakDemoImpl(ii int) {
	// 正确的写法是申明一个带缓存的channel,比如： var ch = make(chan int, 1);而不是 var ch = make(chan int);
	var ch = make(chan int) // 如果定义的channel是一个阻塞的，当使用该channel的协程是一个耗时的逻辑，逻辑结束后再向协程发通知写。
	go func() {
		time.Sleep(1 * time.Second) //模拟一个耗时的逻辑动作
		ch <- 1                     //高耗时后再发通知，往协程内写数据.
		fmt.Println("send notify to chan succ,ii: ", ii)
		close(ch)
	}()

	//
	select {
	case item := <-ch:
		fmt.Println("get data from chan: ", item)
	case <-time.After(10 * time.Millisecond): //当给channel读取操作设置一个超时时间，该时间比从上面读取到数据耗时要短，那么定时等待超时退出等待。
		// 这样从channel中读取数据的动作就跳出，永远也不会从channel读取数，这样业务完成后再往channel中写数据就会阻塞，那么上面的协程就会被阻塞住。
		fmt.Println("wait task timeout, cost time: ", 10*time.Millisecond)

	}
}

// CoroutineLeakDemo
func CoroutineLeakDemo() {
	for i := 0; i < 1000; i++ {
		CoroutineLeakDemoImpl(i)
		time.Sleep(1 * time.Millisecond)
		fmt.Println("run coroutine nums: ", runtime.NumGoroutine())
	}
	time.Sleep(5 * time.Second)
	fmt.Println("last run coroutine nums: ", runtime.NumGoroutine())

}
