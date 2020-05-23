package main

import (
	"fmt"
	"sync"
)

type pipe struct {
	p chan string
}

type sender struct {
	sendPipe pipe
}

func (s *sender) run(wg *sync.WaitGroup) {
	str := "Hello,receiver ! I`m sender\n"
	s.sendPipe.p <- str
	wg.Done()

}

type receiver struct {
	receivePipe pipe
}

func (r *receiver) run(wg *sync.WaitGroup) {
	fmt.Println("the following is from sender:\n" + <-r.receivePipe.p)
	wg.Done()

}

func connect(s *sender, r *receiver, wg *sync.WaitGroup) {
	str := <-s.sendPipe.p
	r.receivePipe.p <- str

	wg.Done()
}

func main() {
	var (
		s  sender
		r  receiver
		wg sync.WaitGroup
	)

	s.sendPipe.p = make(chan string)
	defer close(s.sendPipe.p)
	r.receivePipe.p = make(chan string)
	defer close(r.receivePipe.p)

	wg.Add(3)
	go connect(&s, &r, &wg)
	go s.run(&wg)
	go r.run(&wg)

	wg.Wait()
}
