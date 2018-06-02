package main

import (
	"fmt"
	"time"
)

// グローバル変数
var n = 0

// main()はsub()の終了を待たないため、
// 以下を動かすと、sub()は最後まで処理されずすぐに終わってしまう
func main() {
	fmt.Println("main start")
	
	go sub()
	
	// スリープタイムを10ではなく2000とかにするとsub()の終了を待てる
 	time.Sleep(time.Millisecond * 10)
	
	fmt.Println("main finish")
}

func sub() {
	fmt.Println("sub start")
	
	time.Sleep(time.Millisecond * 1000)
	
	fmt.Println("sub finish")
}