package main

import "fmt"

func main() {
	// １．普通の関数
	fmt.Println(normal_func(1,2))
	
	// ２．引数が可変長
	variablearguments_func("あ", "い", "う", "え", "お")
	variablearguments_func("か")
	
	// ３．戻り値が複数
	var a,b = multi_func()
	fmt.Println(a,b)

	// ４．戻り値に名前を付ける
	// var c,d = namedret_func()  戻り値の変数の名前が合致しないのでダメ
	var ret1,ret2 = namedret_func()
	fmt.Println(ret1, ret2)

	// ５．戻り値に名前を付ける（returnにも戻り値を書いてしまった場合）
	var ret3,ret4 = namedret2_func()
	fmt.Println(ret3, ret4)


}


// １．普通の関数
// 戻り値や引数はあってもなくてもどっちでもいい
// 最後に付いている型（int）が戻り値の型
func normal_func(a int, b int) int {
	return a + b
}

// ２．引数が可変長
func variablearguments_func(a string, params ...string) {
	// １個目の引数
	fmt.Printf("%s : ", a)
	
	// ２個目以降の引数
	// _（アンスコ）：ブランク識別子：省略してると
	for _, p:=range params {
		fmt.Printf("%s", p)
	}
	
	fmt.Printf("\n")
}


// ３．戻り値が複数
func multi_func() (int, string) {
	return 100, "hoge"
}

// ４．戻り値に名前を付ける
func namedret_func() (ret1 int, ret2 int) {
	ret1 = 10		// １個目の戻り値
	ret2 = 20		// ２個目の戻り値
	return
}


// ５．戻り値に名前を付ける（returnにも戻り値を書いてしまった場合）
func namedret2_func() (ret3 int, ret4 int) {
	ret3 = 10		// returnに値を書いているので結局これは無視される
	ret4 = 20		//     〃
	return 30, 40
}



