package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/munirwanis/sort"
)

func main() {
	runSorts(100)
	runSorts(1000)
	runSorts(10000)
	runSorts(100000)
	runSorts(1000000)
}

func runSorts(length int) {
	var array = rand.Perm(length)
	//fmt.Println(array)
	t0 := time.Now()
	sort.BubbleSort(array)
	t1 := time.Now()
	//fmt.Println(array)
	fmt.Printf("%d length took %v\n", length, t1.Sub(t0))
}
