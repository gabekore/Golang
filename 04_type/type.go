package main

import "fmt"


// int型をベースにした型宣言
type MyInt int

// 複数の宣言も可能
type (
	MyInt2 int
	MyInt3 int
)

// 宣言した型の利用
var n1 MyInt2 = 1
var n2 MyInt2 = 2
var n  int   = 10

func main() {
	fmt.Println(n1 + n2)
}
