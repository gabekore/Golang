package pac1

import "fmt"
import "../somepkg"

// mainのグローバル変数を他パッケージから見ることはできない
var Abc int

func Fnc1() {
	fmt.Println("-------------")
	
	fmt.Println("Hello, p1 World!")
  	fmt.Printf("Abc=%d, ", Abc)
 	fmt.Printf("somepkg.Yamada=%s, ", somepkg.Yamada)
 	fmt.Printf("somepkg.Tanaga=%s\n", somepkg.Tanaka)

	fmt.Printf("Some_c1=%s, ", somepkg.Some_c1)
	fmt.Printf("Some_c2=%d, ", somepkg.Some_c2)
	fmt.Printf("Some_c3=%s\n", somepkg.Some_c3)

}