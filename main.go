package main

import (
	"fmt"
	"sync"
)

func printThis(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	letters := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i",
	}

	wg.Add(len(letters))
	for i, letter := range letters {
		go printThis(fmt.Sprintf("%d, %s", i, letter), &wg)
	}
	wg.Wait()
	wg.Add(1)
	printThis("this is the second", &wg)
}
