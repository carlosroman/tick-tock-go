package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/carlosroman/tick-tock-go"
	"github.com/stretchr/testify/mock"
	"time"
)

var _ = Describe("Printer", func() {

	Context("When printer gets time", func() {

		var w *mockWriter
		var p Printer
		BeforeEach(func() {
			w = new(mockWriter)
			p = NewPrinter(w)
		})

		Describe("with default values", func() {
			It("Should print correctly for every seconds", func() {
				t, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
				w.On("Write", mock.Anything).Return(10, nil)
				p.Print(t)
				Expect(
					w.AssertCalled(GinkgoT(), "Write", []byte("tick"))).To(BeTrue(), "Expected tick")

			})

			It("Should print correctly for every minute", func() {
				t, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:00Z")
				w.On("Write", mock.Anything).Return(10, nil)
				p.Print(t)
				Expect(
					w.AssertCalled(GinkgoT(), "Write", []byte("tock"))).To(BeTrue(), "Expected tock")

			})

			It("Should print correctly for every hour", func() {
				t, _ := time.Parse(time.RFC3339, "2006-01-02T15:00:00Z")
				w.On("Write", mock.Anything).Return(10, nil)
				p.Print(t)
				Expect(
					w.AssertCalled(GinkgoT(), "Write", []byte("bong"))).To(BeTrue(), "Expected bong")

			})
		})

		Describe("with updated messages", func() {
			It("Should print correctly for every seconds", func() {
				p.UpdateSecondMsg("quack")
				t, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
				w.On("Write", mock.Anything).Return(10, nil)
				p.Print(t)
				Expect(
					w.AssertCalled(GinkgoT(), "Write", []byte("quack"))).To(BeTrue(), "Expected quack")

			})

			It("Should print correctly for every minute", func() {
				p.UpdateMinuteMsg("woof")
				t, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:00Z")
				w.On("Write", mock.Anything).Return(10, nil)
				p.Print(t)
				Expect(
					w.AssertCalled(GinkgoT(), "Write", []byte("woof"))).To(BeTrue(), "Expected woof")

			})

			It("Should print correctly for every hour", func() {
				p.UpdateHourMsg("meow")
				t, _ := time.Parse(time.RFC3339, "2006-01-02T15:00:00Z")
				w.On("Write", mock.Anything).Return(10, nil)
				p.Print(t)
				Expect(
					w.AssertCalled(GinkgoT(), "Write", []byte("meow"))).To(BeTrue(), "Expected meow")

			})
		})
	})
})

type mockWriter struct {
	mock.Mock
}

func (w *mockWriter) Write(p []byte) (n int, err error) {
	args := w.Called(p)
	return args.Int(0), args.Error(1)
}
