package daemontools

import (
	"encoding/binary"
	"fmt"
	"io/fs"
	"os"
	"path"
	"syscall"
	"time"
)

type Status struct {
	Service    string
	PID        int
	Paused     bool
	Want       byte
	When       time.Time
	NormallyUp bool
}

func (s *Status) String() (o string) {
	hasPid := s.PID > 0
	o = fmt.Sprintf("%s: ", s.Service)
	if hasPid {
		o += fmt.Sprintf("up (pid %d) ", s.PID)
	} else {
		o += "down "
	}

	seconds := time.Now().Unix() - s.When.Unix()
	o += fmt.Sprintf("%d seconds ", seconds)
	if hasPid && !s.NormallyUp {
		o += " normally down"
	}
	if !hasPid && s.NormallyUp {
		o += " normally up"
	}
	if hasPid && s.Paused {
		o += " paused"
	}
	if !hasPid && s.Want == 'u' {
		o += " want up"
	}
	if hasPid && s.Want == 'd' {
		o += "want down"
	}
	return
}

type Error struct {
	Msg string
	Err error
}

func (e Error) Error() string {
	return e.Msg
}

func Svstat(service string) (s *Status, err error) {
	var f *os.File
	f, err = os.OpenFile(path.Join(service, "supervise", "ok"), os.O_WRONLY|syscall.O_NONBLOCK, 0)
	if err != nil {
		if err, ok := err.(*fs.PathError); ok && err.Err.Error() == "no such device or address" {
			return nil, Error{Msg: "supervise not running", Err: err}
		}
		return nil, err
	}
	f.Close()
	f, err = os.Open(path.Join(service, "supervise", "status"))
	if err != nil {
		return nil, err
	}
	b := make([]byte, 18)
	n, err := f.Read(b)
	if err != nil || n != 18 {
		return nil, err
	}
	f.Close()

	var normallyUp bool
	if _, err := os.Stat(path.Join(service, "down")); err != nil && os.IsNotExist(err) {
		normallyUp = true
	}

	s = &Status{
		Service:    service,
		PID:        int(binary.LittleEndian.Uint32(b[12:16])),
		Paused:     b[16] == '1',
		Want:       b[17],
		When:       taiUnpack(b[:8]),
		NormallyUp: normallyUp,
	}
	return s, nil
}
