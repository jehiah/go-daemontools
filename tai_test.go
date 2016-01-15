package daemontools

import (
	"encoding/binary"
	"testing"
	"time"
)

func TestTaiUnpack(t *testing.T) {
	type testCase struct {
		have   uint64
		expect time.Time
	}
	tests := []testCase{
		{0x400000002a2b2c2d, time.Date(1992, 6, 2, 8, 6, 59, 0, time.UTC)},
	}
	for _, tc := range tests {
		b := make([]byte, 8)
		binary.BigEndian.PutUint64(b, tc.have)
		got := taiUnpack(b)
		if !got.Equal(tc.expect) {
			t.Errorf("got %s expected %s for %v", got, tc.expect, tc.have)
		}
	}
}
