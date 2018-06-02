package main

import "fmt"

// interfaceはメソッドを集めるもの
// これで「型」になると考えればいい
// このinterfaceをstructに当てはめて、これらのメソッド名を持っていることを保証できる
type Human interface {
	// var i int   ←変数はinterfaceに入れられない
	hello()
	walk()
}

// Taro構造体
type Taro struct {
	name string
	age  int
}
// Jiro構造体
type Jiro struct {
	english int
	math    int
}


// Taroに紐付く関数
func (m Taro) hello() {
	fmt.Printf("%s%d歳です\n", m.name, m.age)
}
func (m Taro) walk() {
	fmt.Println("トコトコ...")
}
func (m Taro) shout() {
	fmt.Println("わーー！")
}
// Jiroに紐付く関数
func (j Jiro) hello() {
	fmt.Printf("英語の点数=%v, 数学の点数=%v\n", j.english, j.math)
}
func (j Jiro) walk() {
	fmt.Println("歩くの嫌い")
}
func (j Jiro) shout() {
	fmt.Println("叫ばないよ")
}


func main() {
	// Taro構造体
	m := Taro{"たろう", 20}
	j := Jiro{89, 74}
	
	fmt.Println("--> 1")
	m.hello()
 	m.walk()
 	m.shout()

	fmt.Println("--> 2")
	j.hello()
 	j.walk()
 	j.shout()


	// Taro構造体をHuman構造体に入れる（型変換してるようなもの）
	// Humanはhello()とwalk()を持っていて、mも持っているので代入できる
	// Taroがもしもhello()もしくはwalk()を持っていなかった型が合わないので代入できない
	var h1 Human = m
	
	fmt.Println("--> 3")
	h1.hello()
	h1.walk()
// 	h1.shout()	←これはダメ、interfaceに無いから

	var h2 Human = j
	
	fmt.Println("--> 4")
	h2.hello()
	h2.walk()
// 	h2.shout()	←これはダメ、interfaceに無いから
}