package daemontools

import (
	"testing"
)

func TestSvstat(t *testing.T) {
	s, err := Svstat("/service/v3_api")
	if err != nil || s.PID == 0 {
		t.Errorf("%v %v", s, err)
	}
}
