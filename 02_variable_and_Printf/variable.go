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
	"./pac1"
	some "./somepkg"
)

// 初期値を設定していない変数の初期値はゼロ値となる
// 関数の外での宣言はグローバル（パッケージプライベート）スコープ

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

// mainのグローバル変数を他パッケージから見ることはできない
var Abc = 49

// 定数
const c1 string = "hogehoge"
const (
	c2 = 999
	c3 = "fugafuga"
)

func main() {
	// ローカル変数の場合のみこの書き方（:=）ができる
	// varも型も書かなくていいのがメリット（逆に言えば型を書けない）
	x := 100
	y, z := 200, 300

	pac1.Abc = Abc


	// 定数
	const lc1 string = "barbar"
	const (
		lc2 = 888
		lc3 = "foofoo"
	)

	fmt.Println("-------------")
	fmt.Printf("c1=%s, ", c1)
	fmt.Printf("c2=%d, ", c2)
	fmt.Printf("c3=%s, ", c3)

	fmt.Printf("lc1=%s, ", lc1)
	fmt.Printf("lc2=%d, ", lc2)
	fmt.Printf("lc3=%s, ", lc3)

	fmt.Printf("Some_c1=%s, ", some.Some_c1)
	fmt.Printf("Some_c2=%d, ", some.Some_c2)
	fmt.Printf("Some_c3=%s\n", some.Some_c3)


	fmt.Println("-------------")
	
	some.Yamada = "やまだ"
	some.Tanaka = "たなか"
	
	fmt.Println("↓mainで確認")
	fmt.Printf("a=%d, ", a)
	fmt.Printf("b=%d, ", b)
	fmt.Printf("c=%d, ", c)
	fmt.Printf("d=%d, ", d)
	fmt.Printf("e=%d, ", e)
	fmt.Printf("f=%d, ", f)
	fmt.Printf("g=%d, ", g)
	fmt.Printf("h=%d, ", h)
	fmt.Printf("i=%d, ", i)
	fmt.Printf("j=%d, ", j)
	fmt.Printf("k=%d, ", k)
	fmt.Printf("l=%d, ", l)
	fmt.Printf("x=%d, ", x)
	fmt.Printf("y=%d, ", y)
	fmt.Printf("z=%d, ", z)
	fmt.Printf("some.Yamada=%s, ", some.Yamada)
	fmt.Printf("some.Tanaka=%s, ", some.Tanaka)
	fmt.Printf("Abc=%d\n", Abc)
	
	// 同一パッケージの別ファイルで、変数参照と設定
	fnc2()
	
	fmt.Println("-------------")
	fmt.Println("↓mainで再確認")
	fmt.Printf("a=%d, ", a)
	fmt.Printf("b=%d, ", b)
	fmt.Printf("c=%d, ", c)
	fmt.Printf("d=%d, ", d)
	fmt.Printf("e=%d, ", e)
	fmt.Printf("f=%d, ", f)
	fmt.Printf("g=%d, ", g)
	fmt.Printf("h=%d, ", h)
	fmt.Printf("i=%d, ", i)
	fmt.Printf("j=%d, ", j)
	fmt.Printf("k=%d, ", k)
	fmt.Printf("l=%d, ", l)
	fmt.Printf("some.Yamada=%s, ", some.Yamada)
	fmt.Printf("some.Tanaka=%s, ", some.Tanaka)
	fmt.Printf("Abc=%d\n", Abc)
	
	// 別パッケージの別ファイルで、変数参照
	pac1.Fnc1()
}
