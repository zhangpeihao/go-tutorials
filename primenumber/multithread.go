package primenumber

import (
	"fmt"
	"math"
	"sync"
)

type Multithread struct {
	max     int
	threads int
}

func NewMultithread(max, threads int) *Multithread {
	return &Multithread{
		max:     max,
		threads: threads,
	}
}

func (finder *Multithread) Name() string {
	return "Multithread"
}

func filter(value int, result chan<- int) {
	//	fmt.Printf("filter(%d)\n", value)
	sqrt := int(math.Ceil(math.Sqrt(float64(value))))
	for j := 2; j <= sqrt; j++ {
		if value%j == 0 {
			//			fmt.Printf("filter(%d, %d) not prime\n", value, sqrt)
			return
		}
	}
	//	fmt.Printf("filter(%d, %d) prime\n", value, sqrt)
	result <- value
}

func (finder *Multithread) Run() int {
	fmt.Print("")
	count := 1
	source := make(chan int, 1000)
	result := make(chan int, 1000)
	gate := new(sync.WaitGroup)

	go func() {
		gate.Add(1)
		defer gate.Done()
		for i := 2; i < finder.max; i++ {
			source <- i
		}
		for i := 0; i < finder.threads; i++ {
			source <- 0
		}
	}()

	for i := 0; i < finder.threads; i++ {
		go func() {
			gate.Add(1)
			defer gate.Done()
			for {
				value := <-source
				if value == 0 {
					break
				}
				filter(value, result)
			}
		}()
	}

	done := make(chan struct{})
	go func() {
		gate.Wait()
		done <- struct{}{}
	}()

MAIN_LOOP:
	for {
		select {
		case prime := <-result:
			//		fmt.Printf("%d\n", prime)
			if prime > 1 {
				count++
			}
		case <-done:
			count += len(result)
			break MAIN_LOOP
		}
	}
	return count
}
