package id

import "sync"

// rangeSequencer is a Sequencer implementation that generates numbers within a specified range.
type rangeSequencer struct {
	mu    sync.Mutex
	min   int64
	max   int64
	value int64
}

// Generate produces the next number in the sequence, wrapping around to min if the max value is exceeded.
func (s *rangeSequencer) Generate() (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	value := s.value
	s.value++
	if s.value > s.max {
		s.value = s.min
	}
	return value, nil
}

// RangeSequencer creates a new Sequencer that generates numbers within the specified range [min, max].
// If an initial value is provided, it will start from that value, otherwise from min.
func RangeSequencer(min int64, max int64, value ...int64) Sequencer {
	if min > max {
		min = max
	}

	seq := rangeSequencer{
		min: min,
		max: max,
	}

	if len(value) > 0 && value[0] >= min && value[0] <= max {
		seq.value = value[0]
	} else {
		seq.value = min
	}

	return &seq
}
