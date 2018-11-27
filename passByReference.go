package main

import (
	"fmt"
	"sync"
	"time"
)

type fi struct {
	a int
}

func main() {
	var wg sync.WaitGroup
	var sci fi
	for i := 0; i < 20; i++ {
		sci = fi{}
		time.Sleep(1 * time.Nanosecond)
		sci = fi{a: i}
		wg.Add(1)
		go printfi(&sci, &wg)
	}
	wg.Wait()
}

func printfi(hi *fi, wg *sync.WaitGroup) {
	fmt.Println(hi)
	time.Sleep(2 * time.Second)
	wg.Done()
}
