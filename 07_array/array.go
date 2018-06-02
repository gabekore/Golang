package main

import "fmt"

func main() {
	var a [3]int
	b := [2]string{"hoge", "fuga"}
	c := [...]int{1,2}
	
	d := [...]int{1:1, 2:10}
	
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	
	fmt.Println(len(a))
	fmt.Println(len(b))
	fmt.Println(len(c))
	fmt.Println(len(d))
}