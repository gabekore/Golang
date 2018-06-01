package main

// ◆実行方法その１
// $ go run *.go

// ◆実行方法その２
// $ go build *.go
// $ ./p1

// ◆結果
// Hello, p0 World!
// Hello, p1 World!
// Hello, p2 World!
// Hello, p3 World!

import "fmt"
import "./pac3"

// import "pac4"	←異なるpackage同士を同じフォルダ内に置けない

func fnc1() {
	fmt.Println("Hello, p1 World!")
}

func main() {
	fmt.Println("Hello, p0 World!")
	fnc1()
	fnc2()
	pac3.Fnc3()
	// 	pac4.Fnc4()
}
