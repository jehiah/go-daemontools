package daemontools

import (
	"encoding/binary"
	"time"
)

const taiEpoch = 4611686018427387914

// http://cr.yp.to/libtai/tai64.html#tai64n
func taiUnpack(b []byte) time.Time {
	return time.Unix(int64(binary.BigEndian.Uint64(b[:8])-taiEpoch), 0)
}
