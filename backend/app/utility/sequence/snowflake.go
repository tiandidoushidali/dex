package sequence

import (
	"sync"
	"time"
)

const (
	sequenceBits      = uint(12)
	workerBits        = uint(5)
	datacenterBits    = uint(5)
	sequenceMask      = uint64(-1 ^ -1<<sequenceBits) // sequenceBits 位全部是1， 其余是0 000000.....0011111.....
	workerIDShift     = sequenceBits
	datacenterIDShift = sequenceBits + workerBits
	timestampShift    = sequenceBits + workerBits + datacenterBits
)

type Snowflake struct {
	sync.Mutex
	epoch        uint64
	timestamp    uint64
	workerID     uint64
	datacenterID uint64
	sequence     uint64
}

func NewSnowflake(epoch uint64, workerID uint64, datacenterID uint64) *Snowflake {
	return &Snowflake{
		epoch:        epoch,
		timestamp:    0,
		workerID:     workerID,
		datacenterID: datacenterID,
		sequence:     0,
	}
}

func (s *Snowflake) Next() uint64 {
	s.Lock()
	defer s.Unlock()

	now := time.Now().UnixMilli()
	if s.timestamp == uint64(now) {
		s.sequence = (s.sequence + 1) & sequenceMask // 确保不超过每秒sequence数限制
		if s.sequence == 0 {
			// if ID has exceeded the uper limit
			// you need to wait for some millisecond before continuing to generate
			for now <= int64(s.timestamp) {
				now = time.Now().UnixMilli()
			}
		}
	} else {
		s.sequence = 0
	}
	var t uint64 = uint64(now - int64(s.epoch))
	s.timestamp = uint64(now)
	seq := (t << timestampShift) |
		(s.datacenterID << datacenterIDShift) |
		(s.workerID << workerIDShift) |
		s.sequence

	return seq
}
