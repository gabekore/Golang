package main

import "fmt"


func main() {
	var p *int
	
	n := 10
	
	p = &n
	
	fmt.Println("ポインタ自体のアドレス：", &p)
	fmt.Println("ポインタが保持する値  ：",  p)
	fmt.Println("ポインタ経由の変数の値：", *p)
	fmt.Println("変数のアドレス        ：", &n)
	fmt.Println("変数の値              ：",  n)
	
	fmt.Println("------------------")
	var p1 *int
	fmt.Println("p :", p1)
	p1 = new(int)
	fmt.Println("p :", p1)
	fmt.Println("*p:", *p1)
	*p1 = 5
	fmt.Println("p :", p1)
	fmt.Println("*p:", *p1)
	
// 	fmt.Println("------------------")
// 	var p2 *int
// 	fmt.Println("p :", p2)
// 	p2 = &(34)    ←この書き方はできないよ
// 	fmt.Println("p :", p2)
// 	fmt.Println("*p:", *p2)
// 	*p2 = 87
// 	fmt.Println("p :", p2)
// 	fmt.Println("*p:", *p2)
}