package main

import "fmt"

// sub2()やsub3()をコールしている箇所とpanic("パニック2")やpanic("パニック3")の位置を
// 入れ替えて動かしてみる

// $ go run panic.go 
// main start
// sub1 start
// sub2 start
// panic: パニック2
// 
// goroutine 1 [running]:
// main.sub2()
// 	/Users/simauma/OneDrive/BootTech/GitHubRepo/Golang/15_panic/panic.go:28 +0x79
// main.sub1()
// 	/Users/simauma/OneDrive/BootTech/GitHubRepo/Golang/15_panic/panic.go:20 +0x62
// main.main()
// 	/Users/simauma/OneDrive/BootTech/GitHubRepo/Golang/15_panic/panic.go:11 +0x66
// exit status 2

// ↑こんな感じでエラーとなるので、エラー発生の行番号とかとソースコードを見比べてみて、
// 処理されない命令はどこか等を確認してみる


func main() {
	fmt.Println("main start")
	
	sub1()
	
	fmt.Println("main finish")
	
}


func sub1() {
	fmt.Println("sub1 start")
	sub2()
	panic("パニック1")
	fmt.Println("sub1 finish")
}


func sub2() {
	fmt.Println("sub2 start")
	panic("パニック2")
	sub3()
	fmt.Println("sub2 finish")
}


func sub3() {
	fmt.Println("sub3 start")
	
	panic("パニック3")
	fmt.Println("sub3 finish")
}