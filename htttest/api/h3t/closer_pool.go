package h3t

import "io"

type CloserPool struct {
	clist []io.Closer
}

func (inst *CloserPool) Add(c io.Closer) {
	if c == nil {
		return
	}
	inst.clist = append(inst.clist, c)
}

func (inst *CloserPool) Close() error {

	list := inst.clist
	count := len(list)
	var err error

	for i := count - 1; i >= 0; i-- {
		cl := list[i]
		if cl == nil {
			continue
		}
		e := cl.Close()
		if e != nil {
			err = e
		}
	}

	return err
}
