package main

import "fmt"
import "strconv"

func main() {
	// 関数型のポインタを定義
	// 引数はint,int、戻り値はstring
	var fnc  func(int,int) string
	
	// 関数ポインタの取得
	fnc = add
	fmt.Println(fnc(11,22))
	
	fnc = sub
	fmt.Println(fnc(15,29))
}

func add(a int, b int) string {
	return strconv.Itoa(a + b)
}
func sub(a int, b int) string {
	return strconv.Itoa(a - b)
}