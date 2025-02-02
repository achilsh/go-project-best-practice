package easy_make_mistake

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroupCallWithMistakeAdd 演示错误使用 WaitGroup's Add()方法场景
func WaitGroupCallWithMistakeAdd() {
	var wg sync.WaitGroup
	for i := 0; i < 6; i++ {
		//wg.Add(1) //this is 正确的写法
		go func(ii int) {
			// 这是错误的写法。
			wg.Add(1) //如果在协程内部调用add(),这样会导致 外层 wg.Wait() 因为计数器没有加1，都是0 就不在等待.
			//正确的做法是在协程外层调用wg.Add(1) 这样 wg.Wait()因为计数叠加 有值才会被Wait()住。
			time.Sleep(100 * time.Millisecond)
			fmt.Println("call step: ", ii)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("call wait group done.")
}
