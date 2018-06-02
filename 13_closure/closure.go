package main

import "fmt"

func main() {
	fmt.Println("====> 1")
	var fnc = f()	// ここでf()が実行される
	
	// fnc変数にはクロージャーが入っているので、
	// 以降、f()は動かず、クロージャーのみが動く
	
	fmt.Println("====> 2")
	fmt.Println(fnc())
	
	fmt.Println("====> 3")
	fmt.Println(fnc())
	
	fmt.Println("====> 4")
	fmt.Println(fnc())
}

// 「func() int」がf()の戻り値
// func()はクロージャーと言うか無名関数で、名前はfunc固定
// func以外の名前だとビルド通らない
func  f()  func() int {
	fmt.Println("--> f()")
	a := 0
	
	// クロージャー
	return func() int {
		fmt.Println("--> sub()")
		a++
		return a
	}
}