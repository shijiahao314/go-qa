package main

import (
	"fmt"
	"math/rand"
	"slices"
	"sort"
)

func main() {
	n := make([]int, 10, 10)
	for idx := range n {
		n[idx] = rand.Intn(999)
	}
	fmt.Println(n)
	sort.Ints(n)
	fmt.Println(n)
	slices.Reverse(n)
	fmt.Println(n)

}
