package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"
	. "github.com/carlosroman/tick-tock-go"
	"github.com/stretchr/testify/mock"
	"sync"
	"time"
)

var _ = Describe("Processor", func() {

	Describe("should call printer", func() {
		It("should call printer", func() {
			t := time.Now()
			prt := new(printMock)
			p := NewProcessor(prt)
			c := make(chan time.Time, 1)
			var wg sync.WaitGroup
			wg.Add(1)
			prt.On("Print", t).Once()
			go func() {
				go func() {
					fmt.Printf("Sending T: %s\n", t)
					c <- t
					fmt.Println("calling done")
					p.Done()
					wg.Done()
				}()
				p.RunProcessor(c)
			}()
			wg.Wait()
			Expect(prt.AssertNumberOfCalls(GinkgoT(), "Print", 1)).To(BeTrue())
		})
	})
})

type printMock struct {
	mock.Mock
}

func (p *printMock) Print(t time.Time) {
	p.Called(t)
}

func (p *printMock) UpdateSecondMsg(msg string) {
	p.Called(msg)
}

func (p *printMock) UpdateMinuteMsg(msg string) {
	p.Called(msg)
}

func (p *printMock) UpdateHourMsg(msg string) {
	p.Called(msg)
}

//
//type tickerMock struct {
//	mock.Mock
//}
//
//func (t *tickerMock) Stop() {
//	t.Called()
//}
//
//func (t *tickerMock) Tick(d time.Duration) <-chan time.Time {
//	args := t.Called(d)
//	mt := args.Get(0).(time.Time)
//	c := make(chan time.Time, 1)
//	defer func() {
//		c <- mt
//	}()
//	return c
//}
