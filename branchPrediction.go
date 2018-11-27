package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {

	var (
		n     = 100000
		pivot = 120
		a     = make([]int, n)
		max   = 1000
		count = 0
	)
	t := time.Now()
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(max)
	}

	sort.Ints(a)

	//fmt.Println(a)

	for i := 0; i < n; i++ {
		if a[i] > pivot {
			count++
		}
	}

	fmt.Println("Count:", count, "\nTime:", time.Since(t))

}
