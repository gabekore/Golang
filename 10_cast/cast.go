package main

import "fmt"

func main() {
	var a int = 1
	var b int16 = 2
	var c int = a + int(b)
// 	var c int = a + b    これはダメ
	
	fmt.Println(c)
}