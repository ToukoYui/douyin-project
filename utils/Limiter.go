package utils

import (
	"golang.org/x/time/rate"
	"sync"
	"time"
)

/**
* 接口限流工具
*
* @author: 张庭杰
* @date: 2023年02月17日 16:25
 *
*/

type Limiters struct {
	limiters *sync.Map
}

type Limiter struct {
	limiter *rate.Limiter
	lastGet time.Time //关键参数:上一次获取token的时间,是我们标记用户访问接口的关键参数
	key     string
}

var GlobalLimiters = &Limiters{
	limiters: &sync.Map{},
}
var once = sync.Once{}

// NewLimiter 通过传入{r:往桶里放Token的速率,b:令牌桶的大小,可以对某个id/ip做限制},新建一个自定义的限流器
func NewLimiter(r rate.Limit, b int, key string) *Limiter {
	once.Do(func() {
		go GlobalLimiters.clearLimiter()
	})
	keyLimiter := GlobalLimiters.getLimiter(r, b, key)
	return keyLimiter
}

func (l *Limiter) Allow() bool {
	l.lastGet = time.Now()
	return l.limiter.Allow()
}

// r:往桶里放Token的速率 b:令牌桶的大小 key:可对某id\ip做限制
func (ls *Limiters) getLimiter(r rate.Limit, b int, key string) *Limiter {
	limiter, ok := ls.limiters.Load(key)
	if ok {

		return limiter.(*Limiter)
	}
	l := &Limiter{
		limiter: rate.NewLimiter(r, b),
		lastGet: time.Now(),
		key:     key,
	}
	ls.limiters.Store(key, l)
	return l
}

// 清除过期的限流器
func (ls *Limiters) clearLimiter() {
	for {
		time.Sleep(1 * time.Minute)
		ls.limiters.Range(func(key, value interface{}) bool {
			//超过1分钟
			if time.Now().Unix()-value.(*Limiter).lastGet.Unix() > 60 {
				ls.limiters.Delete(key)
			}
			return true
		})
	}
}
