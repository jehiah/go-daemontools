package daemontools

import (
	"os"
	"testing"
)

func TestSvstat(t *testing.T) {
	_, err := os.Stat("/service/nsqd")
	if err != nil && os.IsNotExist(err) {
		t.SkipNow()
	}
	s, err := Svstat("/service/nsqd")
	if err != nil || s.PID == 0 {
		t.Errorf("%v %v", s, err)
	}
}
