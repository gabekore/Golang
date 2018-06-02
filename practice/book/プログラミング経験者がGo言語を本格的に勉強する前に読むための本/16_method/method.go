package main

import "fmt"

// $ go run method.go 

type Calc struct {
	a,b int
}

type MyInt int

// メソッド
// Calc構造体に紐付く関数
// 引数なし、戻り値int
// ※main()内の変数「ccc」に紐付くと思えばいい
// (p Calc)が紐付く先と思えばいい
// 紐付く先をレシーバと呼ぶらしいけど、よく分からん
// レシーバにできるのはtypeで定義したもののみ（ポインタも可能）
func (p Calc) Add() int {
	return p.a + p.b
}

// メソッド
// MyInt型に紐付く関数
// 引数int、戻り値MyInt
// ※main()内の変数「mmm」に紐付くと思えばいい
// (m MyInt)が紐付く先と思えばいい
// 紐付く先をレシーバと呼ぶらしいけど、よく分からん
// レシーバにできるのはtypeで定義したもののみ（ポインタも可能）
func (m MyInt) Add(n int) MyInt {
	return m + MyInt(n)
}


func main() {
	ccc := Calc{3,2}
	var mmm MyInt = 1
	
	fmt.Println(ccc.Add())	// 5
	fmt.Println(mmm.Add(2))	// 3
}