package main

import (
	"io"
	"sync"
	"time"
)

const (
	secDefault = "tick"
	minDefault = "tock"
	hrDefault  = "bong"
)

type Printer interface {
	Print(t time.Time)
	UpdateSecondMsg(msg string)
	UpdateMinuteMsg(msg string)
	UpdateHourMsg(msg string)
}

func NewPrinter(out io.Writer) Printer {
	return &printer{
		out: out,
		sec: []byte(secDefault),
		min: []byte(minDefault),
		hr:  []byte(hrDefault),
	}
}

type printer struct {
	sync.RWMutex
	out io.Writer
	sec []byte
	min []byte
	hr  []byte
}

func (p *printer) UpdateSecondMsg(msg string) {
	p.Lock()
	defer p.Unlock()
	p.sec = []byte(msg)
}

func (p *printer) UpdateMinuteMsg(msg string) {
	p.Lock()
	defer p.Unlock()
	p.min = []byte(msg)
}

func (p *printer) UpdateHourMsg(msg string) {
	p.Lock()
	defer p.Unlock()
	p.hr = []byte(msg)
}

func (p *printer) Print(t time.Time) {
	p.RLock()
	defer p.RUnlock()
	s := t.Second()
	m := t.Minute()

	if s > 0 {
		p.out.Write(p.sec)
		return
	}

	if m > 0 {
		p.out.Write(p.min)
		return
	}

	p.out.Write(p.hr)
}
