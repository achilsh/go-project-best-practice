package unit_test_demo

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func GetRedisHandle(ip string, port int, passwd string, dbIndex int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%d", ip, port),
		Password: passwd,  // no password set
		DB:       dbIndex, // use default DB
	})
	return rdb
}
func GetRedisHandleV2(host string, passwd string, dbIndex int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: passwd,  // no password set
		DB:       dbIndex, // use default DB
	})
	return rdb
}
func TestRedisMock(t *testing.T) {
	//获取一个本地redis的服务的句柄
	mRdsHandle := miniredis.RunT(t)

	//获取本地redis的ip和端口，其他redis 客户端可使用和连接该redis本地服务
	ipPortStr := mRdsHandle.Addr()
	//
	rdsCli := GetRedisHandleV2(ipPortStr, "", 10)

	testCase := []struct {
		inKey    string
		inValue  string
		outValue string
	}{
		{"aaa", "1111", "1111"},
	}
	for i := 0; i < len(testCase); i++ {
		rdsCli.Set(context.Background(), testCase[i].inKey, testCase[i].inValue, 10*time.Second)
		
		//手动把所有的ttl时间减少特定时间。这样就不需要等待实际过期时间。如果后续ttl时间<=0；过期key就会淘汰。
		mRdsHandle.FastForward(11 * time.Second)

		ret, err := rdsCli.Get(context.Background(), testCase[i].inKey).Result()
		assert.Nil(t, err)
		t.Logf("data: %v", ret)
		assert.Equal(t, ret, testCase[i].outValue)
	}

}
