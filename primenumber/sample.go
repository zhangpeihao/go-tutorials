package primenumber

import (
	"fmt"
	"math"
)

type Sample struct {
	max int
}

func NewSample(max int) *Sample {
	return &Sample{
		max: max,
	}
}

func (finder *Sample) Name() string {
	return "Sample"
}

func (finder *Sample) Run() int {
	fmt.Printf("")
	count := 1
MAIN_LOOP:
	for i := 2; i < finder.max; i++ {
		sqrt := int(math.Ceil(math.Sqrt(float64(i))))
		for j := 2; j <= sqrt; j++ {
			if i%j == 0 {
				continue MAIN_LOOP
			}
		}
		count++
		//		fmt.Printf("%d\n", i)
	}
	return count
}
