package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/zhangpeihao/go-tutorials/primenumber"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)

	max := 100000
	var finder primenumber.PrimeFinder
	finder = primenumber.NewSample(max)
	run(finder)
	finder = primenumber.NewMultithread(max, 8)
	run(finder)
	finder = primenumber.NewMultiChannel(max)
	run(finder)
	finder = primenumber.NewPrimeList(max)
	run(finder)
}

func run(finder primenumber.PrimeFinder) {
	startAt := time.Now()
	count := finder.Run()
	fmt.Printf("%s result is %d, cost %.3fms\n", finder.Name(), count, float64(time.Now().Sub(startAt))/1000000.0)
}
