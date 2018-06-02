package main

import "fmt"

// スライス型とは可変長な配列
// [...]ではなく[]
// あんまり気にする必要は無い

func main() {
	var a []int
	b := []string{"hoge", "fuga"}
	c := []int{1,2}
	
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}