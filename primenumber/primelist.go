package primenumber

import (
	"fmt"
	"math"
)

type PrimeList struct {
	max   int
	first *PrimeItem
}

type PrimeItem struct {
	value int
	next  *PrimeItem
}

func NewPrimeList(max int) *PrimeList {
	return &PrimeList{
		max: max,
		first: &PrimeItem{
			value: 2,
		},
	}
}

func (finder *PrimeList) Name() string {
	return "PrimeList"
}

func (finder *PrimeList) Run() int {
	fmt.Print("")
	count := 1
	tail := finder.first
MAIN_LOOP:
	for i := 3; i < finder.max; i++ {
		item := finder.first
		sqrt := int(math.Ceil(math.Sqrt(float64(i))))
		for {
			if item.value > sqrt {
				break
			}
			if i%item.value == 0 {
				continue MAIN_LOOP
			}
			if item.next != nil {
				item = item.next
			} else {
				break
			}
		}
		// New prime
		tail.next = &PrimeItem{
			value: i,
		}
		tail = tail.next
		count++
	}
	return count
}
