package main

import (
	"fmt"
	"errors"
)

func main() {
	// test()がエラーを返してくる
	e := test()
	
	// 発生したエラーを確認する
	fmt.Println(e.Error())
}


func test() error {
	return errors.New("errors.Newでエラーを発生させる")
}