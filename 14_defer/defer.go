package main

import "fmt"

// ◆実行結果
// $ go run defer.go 
// main start1!
// main start2!
// main start3!
// main start4!
// sub start1!
// sub start2!
// sub start3!
// sub start4!
// sub finish!
// main finish!

func main() {
	// 最後に処理される
	// リソースの解放等、最後にやりたい処理を書いておく
	// newのすぐ下でdeferを書いておけば取得と解放が同時に書けて便利
	defer fmt.Println("main finish!")
	
	fmt.Println("main start1!")
	fmt.Println("main start2!")
	fmt.Println("main start3!")
	fmt.Println("main start4!")
	
	sub()
}

func sub() {
	defer fmt.Println("sub finish!")
	
	fmt.Println("sub start1!")
	fmt.Println("sub start2!")
	fmt.Println("sub start3!")
	fmt.Println("sub start4!")
}