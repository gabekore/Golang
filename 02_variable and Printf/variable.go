package main

// 変数の宣言色々とfmt.Printf


// ◆実行方法その１
// $ go run variable.go

// ◆実行方法その２
// $ go build variable.go 
// $ ./variable 

// ◆結果
// a=0
// b=10
// c=20
// d=0
// e=30
// f=40
// g=0
// h=0
// i=50
// j=60
// k=70
// l=80



import (
	"fmt"
)

var a int
var b int = 10
var c = 20

var (
	d int
	e int = 30
	f     = 40
)

var g, h int
var i, j int = 50, 60
var k, l = 70, 80

func main() {
	fmt.Printf("a=%d\n", a)
	fmt.Printf("b=%d\n", b)
	fmt.Printf("c=%d\n", c)
	fmt.Printf("d=%d\n", d)
	fmt.Printf("e=%d\n", e)
	fmt.Printf("f=%d\n", f)
	fmt.Printf("g=%d\n", g)
	fmt.Printf("h=%d\n", h)
	fmt.Printf("i=%d\n", i)
	fmt.Printf("j=%d\n", j)
	fmt.Printf("k=%d\n", k)
	fmt.Printf("l=%d\n", l)
}
