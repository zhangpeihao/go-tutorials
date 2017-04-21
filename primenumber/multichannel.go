package primenumber

import (
	"fmt"
	"math"
	"sync"
)

const (
	MaxChannelSize = 10000
)

type MultiChannel struct {
	max   int
	first *PrimeChannel
	gate  sync.WaitGroup
}

type PrimeChannel struct {
	value int
	src   chan int
	next  *PrimeChannel
}

func NewMultiChannel(max int) *MultiChannel {
	return &MultiChannel{
		max: max,
		first: &PrimeChannel{
			value: 2,
			src:   make(chan int, int(math.Min(float64(max/2), float64(MaxChannelSize)))),
		},
	}
}

func (finder *MultiChannel) Name() string {
	return "MultiChannel"
}

func (finder *MultiChannel) generate() {
	finder.gate.Add(1)
	defer finder.gate.Done()

	ch := finder.first.src
	for i := finder.first.value + 1; i <= finder.max; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
	ch <- 0
}

func (finder *MultiChannel) filter(src <-chan int, dst chan<- int, prime int) {
}

func (finder *MultiChannel) Run() int {
	fmt.Print("")
	finder.gate.Add(1)
	go finder.first.run(&finder.gate)

	go finder.generate() // Start generate() as a subprocess.

	finder.gate.Wait()
	count := 0
	prime := finder.first
	for prime != nil {
		count++
		prime = prime.next
	}
	return count
}

func (prime *PrimeChannel) run(gate *sync.WaitGroup) {
	defer gate.Done()

	for i := range prime.src {
		//fmt.Printf("filter(%d) i: %d\n", prime.value, i)
		if i%prime.value != 0 {
			//fmt.Printf("filter(%d) i: got %d\n", prime.value, i)
			if prime.next == nil {
				//fmt.Printf("filter(%d) new prime %d\n", prime.value, i)
				size := MaxChannelSize / i
				if size < 256 {
					size = 256
				}
				prime.next = &PrimeChannel{
					value: i,
					src:   make(chan int, size),
				}
				gate.Add(1)
				go prime.next.run(gate)
			}
			prime.next.src <- i
		} else if i == 0 {
			if prime.next != nil {
				prime.next.src <- 0
			}
			//fmt.Printf("filter(%d) push 0\n", prime.value)
			break
		}
	}

}
