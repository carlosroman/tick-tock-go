package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	out := os.Stdout
	ptr := NewPrinter(out)
	p := NewProcessor(ptr)

	for {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		go func() {
			time.Sleep(10 * time.Minute)
			p.Done()
		}()

		p.RunProcessor(ticker.C)
		fmt.Println("---------------------")
		fmt.Println("Change tick to:")
		fmt.Println("a) quack")
		fmt.Println("b) tick")
		fmt.Print("-> ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		switch text {
		case "a":
			ptr.UpdateSecondMsg("quack")
			break

		case "b":
			ptr.UpdateSecondMsg("tick")
			break

		default:
			break
		}

	}
}
