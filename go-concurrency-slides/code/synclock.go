package main

import (
	"fmt"
	"sync"
	"time"
)

type printer struct {
	sync.Mutex
}

func (p *printer) Print(msg string) {
	p.Lock()
	fmt.Printf("Locked and loaded: %v\n", msg)
	time.Sleep(time.Second)
	p.Unlock()
}
func main() {
	prnt := new(printer)
	wg := new(sync.WaitGroup)
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(j int, w *sync.WaitGroup) {
			prnt.Print(fmt.Sprintf("Hello from Gopher #%d", j))
			w.Done()
		}(i, wg)
	}
	wg.Wait()
}
