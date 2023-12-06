package snowflake

import (
	"sync"
	"time"
)

// Snowflake 结构体
type Snowflake struct {
	mu        sync.Mutex
	timestamp uint64
	machineID uint64
	counter   uint64
}

// NewSnowflake 创建 Snowflake 实例
func NewSnowflake(machineID uint64) *Snowflake {
	return &Snowflake{machineID: machineID}
}

// NextID 生成下一个唯一 ID
func (s *Snowflake) NextID() uint64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := uint64(time.Now().UnixNano() / 1000000) // 毫秒级别时间戳
	if now == s.timestamp {
		s.counter++
		if s.counter >= 1<<12 {
			for now <= s.timestamp {
				now = uint64(time.Now().UnixNano() / 1000000)
			}
		}
	} else {
		s.counter = 0
		s.timestamp = now
	}

	id := (now << 22) | (s.machineID << 12) | (s.counter)
	return id
}

func GetSnowflakeID() uint64 {
	// 创建 Snowflake 实例，指定机器 ID
	sf := NewSnowflake(1)

	// 生成1个唯一 ID
	return sf.NextID()
}
