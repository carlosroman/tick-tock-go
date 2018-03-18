package main

import (
	"fmt"
	"time"
)

type Processor interface {
	RunProcessor(c chan time.Time)
	Done()
}

type processor struct {
	done chan bool
	ptr  Printer
}

func NewProcessor(ptr Printer) Processor {
	return &processor{ptr: ptr, done: make(chan bool)}
}

func (p *processor) RunProcessor(c chan time.Time) {
	for {
		select {
		case <-p.done:
			fmt.Println("Done!")
			return
		case t := <-c:
			p.ptr.Print(t)
		}
	}
}

func (p *processor) Done() {
	p.done <- true
}
