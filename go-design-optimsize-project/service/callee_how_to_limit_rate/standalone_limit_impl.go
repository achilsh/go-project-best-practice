package callee_how_to_limit_rate

// 使用标准库中限流器来验证 令牌桶的单机限流方式。
import (
	"time"

	"golang.org/x/time/rate"
)

type Limiter struct {
	//其中 rate.Limiter 有三个方法：Allow, Reserve, and Wait.
	// 他们的区别是： Allow()如果没有token,调用就返回false.
	// Reserve() 如果没有token,调用返回 Reservation，包含等待使用token的时长。
	// Wait() 如果没有token, 调用将会阻塞直到有token为止。
	lim *rate.Limiter
	r   int //每秒产生的 token 数
	b   int //令牌桶的容量，限流器能够处理的突发流量的大小；该值高说明能处理较高的突发流量峰值。
}

func NewLimiter(r int, b int) *Limiter {
	item := &Limiter{
		r:   r, //
		b:   b, //
		lim: rate.NewLimiter(rate.Every(time.Second/time.Duration(r)), b),
	}
	return item
}

// Check 如果有token，则返回 true，否则返回false.
func (l *Limiter) Check() bool {
	return l.lim.Allow()
}
