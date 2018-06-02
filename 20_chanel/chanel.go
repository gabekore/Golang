package main

import "fmt"

//-----------------------------------------
// chan 型           送受信用のチャネル
// chan <- 型        送信専用のチャネル
// <- chan 型        受信専用のチャネル
//-----------------------------------------

func main() {
	// チャネルの生成
	ch := make(chan string)
	
	// ゴルーチンの生成
	go send(ch)
	
	// チャネルによる値の受信
	c := <-ch
	fmt.Printf("[値を受信：%s]\n", c)
}

func send(ch chan int) {
	fmt.Println("[値を送信：abc]")
	
	// チャネルによる値の送信
	ch <- "abc"
}