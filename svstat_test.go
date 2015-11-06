package daemontools

import (
	"testing"
)

func TestSvstat(t *testing.T) {
	s, err := Svstat("/service/test")
	t.Errorf("%v %s", s, err)
}