package main

import "fmt"
import "reflect"

type MyStruct struct {
	a string
	b int8
	c int16
}

type MyStruct2 struct {
	MyStruct  // メンバ名書いてないので継承みたいなイメージ
	d int32
}

type MyStruct3 struct {
	st MyStruct
	e int32
}

type MyStruct4 struct {
	f string `tag1:"value1" tag2:"value2"`
	g int `tag3:"value3"`
}


func main() {
	var st MyStruct
	st.a = "hoge"
	st.b = 10
	st.c = 20
	
	fmt.Println(st)		//=> {hoge 10 20}
	fmt.Println(st.a, st.b, st.c )		//=> hoge 10 20


	var st2 MyStruct2
	st2.a = "aaaaa"
	st2.b = 111
	st2.c = 222
	st2.d = 45
	fmt.Println(st2)
	
	var st3 MyStruct3
	st3.st.a = "bbb"
	st3.st.b = 123
	st3.st.c = 456
	st3.e    = 789
	fmt.Println(st3)

	var st4 MyStruct4
	field0 := reflect.TypeOf(st4).Field(0)
	field1 := reflect.TypeOf(st4).Field(1)
	fmt.Println(field0.Tag.Get("tag1"))
	fmt.Println(field1.Tag.Get("tag3"))
}