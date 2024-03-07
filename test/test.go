package main

import "fmt"

func main() {
	s1 := "asd"
	s2 := "你好"
	fmt.Println(len(s1))
	fmt.Println(len(s2))
	for idx, c := range s2 {
		fmt.Println(idx, c)
	}

}
