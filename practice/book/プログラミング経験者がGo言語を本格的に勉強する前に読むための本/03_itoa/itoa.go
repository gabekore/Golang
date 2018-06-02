package main

import "fmt"

func main() {
	const (
		a = iota	// 0
		b = iota	// 1
		c = 99
		d = iota	// 3
		e = iota	// 4
		f = "abc"
		g
		h = iota	// 7
	)
	
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("c=", c)
	fmt.Println("d=", d)
	fmt.Println("e=", e)
	fmt.Println("f=", f)
	fmt.Println("g=", g)
	fmt.Println("h=", h)
	}