package main

import (
	"fmt"
	"sync"
)

var (
	msg string
	wg  sync.WaitGroup
)

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	msg = s
	m.Unlock()
}

func main() {
	msg = "How are you?"

	var mutex sync.Mutex

	wg.Add(2)
	go updateMessage("Not bad", &mutex)
	go updateMessage("Just fine", &mutex)
	wg.Wait()

	fmt.Println(msg)
}
