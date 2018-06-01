package main


import (
	"fmt"
	some "./somepkg"
)

func fnc2() {
	fmt.Println("-------------")
	
	fmt.Println("↓fnc2")
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
	
	a = 1
	b = 2
	c = 3
	d = 4
	e = 5
	f = 6
	g = 7
	h = 8
	i = 9
	j = 10
	k = 11
	l = 12
	some.Yamada = "ヤマダ"
	some.Tanaka = "タナカ"
	
	fmt.Println("-------------")
	fmt.Println("↓fnc2 変数の値変更")
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
	
	// 定数参照
	fmt.Printf("c1=%s, ", c1)
	fmt.Printf("c2=%d, ", c2)
	fmt.Printf("c3=%s\n", c3)

// 	fmt.Printf("lc1=%s, ", lc1)
// 	fmt.Printf("lc2=%d, ", lc2)
// 	fmt.Printf("lc3=%s\n", lc3)

}