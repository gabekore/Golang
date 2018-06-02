package main

import "fmt"

func main() {
	// mapは連想配列やハッシュや辞書みたいなもん
	m := map[string]int{"a":1, "b":2}
	fmt.Println(m)
	
	m["b"] = 39
	fmt.Println("a =", m["a"], ", b =", m["b"])
}