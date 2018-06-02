package main

import "fmt"

// $ go run method2.go 

// 親クラスみたいなイメージ
type Calc struct {
	a,b int
}

// 子クラスみたいなイメージ
type MyStruct struct {
	Calc  // メンバ名書いてないので継承みたいなイメージ
}

// 親クラスに紐付けていることに注意
func (p Calc) Add() int {
	return p.a + p.b
}


func main() {
	// 子クラスを使っている
	var s MyStruct
	s.a = 5
	s.b = 4

	// Add()メソッドは親クラスのメソッドだが、
	// 継承しているので子クラスでも使える
	fmt.Println(s.Add())	// 5
}