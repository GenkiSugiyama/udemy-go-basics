package main

import (
	"fmt"
)

// const宣言は定数となる
const secret = "abc"

// type宣言で独自の型を宣言可能
type Os int

const (
	Mac Os = iota + 1
	Windos
	Linux
)

// var() で変数を一括定義
var (
	a int
	b string
	c bool
)

func main() {
	// i := 2
	// ui := uint16(2)
	// fmt.Println("iの値：", i)
	// fmt.Printf("i: %v %T\n", i, i)
	// fmt.Printf("i: %[1]v %[1]T ui: %[2]v %[2]T\n", i, ui)

	// // 複数の変数を同時に定義可能
	// pi, title := 3.14, "Go"
	// fmt.Printf("pi: %v title: %v\n", pi, title)

	// x := 10
	// y := 1.23
	// z := float64(x) + y
	// fmt.Println(z)

	// fmt.Printf("Mac: %v, Windows: %v, Linux: %v\n", Mac, Windos, Linux)

	// fmt.Printf("a: %v, b: %v, c: %v", a, b, c)

	// pointer
	// var ui1 uint16                                 // 代入していないので値は「０」
	// fmt.Printf("memory address of ui: %p\n", &ui1) // 変数名の前に「&」をつけるとその変数の値が格納されているメモリの先頭アドレス（ポインタ 8バイト）を取得、「%p」で文字列内に出力
	// var ui2 uint16
	// fmt.Printf("memory address of ui: %p\n", &ui2)

	// var p1 *uint16                      // 型の前に「*」をつけることでポインタの宣言
	// fmt.Printf("value of p1: %v\n", p1) // メモリアドレスが入っていないので値は「nil」
	// p1 = &ui1
	// fmt.Printf("value of p1; %v\n", p1)
	// fmt.Printf("size of p1: %d[bytes]\n", unsafe.Sizeof(p1))
	// fmt.Printf("memory address of p1: %p\n", &p1)

	// fmt.Printf("value of ui1(dereference): %v\n", *p1) // ポインタ変数に「*」をつけることで「ポインタ変数が格納しているアドレス」に格納されている値を参照できる(dereference)
	// *p1 = 2
	// fmt.Printf("value of ui1: %v\n", ui1) // dereference で値を再代入するとui1の値も変わっている（そりゃそうだ）

	// var pp1 **uint16 = &p1 // ダブルポインタ ポインタのポインタを取得するには「**」をつける
	// fmt.Printf("value of pp1: %v\n", pp1)
	// fmt.Printf("value of p1(dereference): %v\n", *pp1)

	// shadowing
	// 「:=」で同じ変数名を定義すると別々のメモリに値が格納される
	// 共通の変数を使いまわしたい場合はコロンなしの「=」で値を代入する
	// ok, result := true, "A"
	// fmt.Printf("memory addresss of result: %p\n", &result)

	// if ok {
	// 	result = "B"
	// 	fmt.Println(result)
	// 	fmt.Printf("memory address of result: %p\n", &result)
	// } else {
	// 	result := "C"
	// 	fmt.Println(result)
	// }
	// fmt.Println(result)

	// slice
	//　配列の定義 [要素数]要素の型 で初期化 intの配列は要素数分の 0 が初期値に入る
	var a1 [3]int
	var a2 = [3]int{1, 2, 3}
	a3 := [...]int{1, 2} // [...] と記述すると代入時の要素数を自動カウントしてくれる

	fmt.Printf("%v, %v, %v\n", a1, a2, a3)
	fmt.Printf("%v, %v\n", len(a3), cap(a3)) // cap()は中身の要素の型によって出力結果が異なる、配列の場合は要素数でlen()と同じ
	fmt.Printf("%T, %T\n", a2, a3)           // 配列の型は中身の型と要素数によって決まる、同じ数値配列でも要素数が異なると別々の型となる

	// slice 配列は一度宣言したら要素数の追加等できない、あとから要素を追加したりする場合は slice型を使う
	var s1 []int // []の中身空で slice の宣言
	s2 := []int{}
	fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %[1]T %[1]v %v %v\n", s2, len(s2), cap(s2))
	fmt.Println(s1 == nil)   // var で宣言された slice は nil 扱いだが
	fmt.Println(s2 == nil)   // := で宣言された slice は nil ではない
	s1 = append(s1, 1, 2, 3) // append(slice変数, 要素1, 要素2, ~) で既存slice に要素を追加する
	fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1))

}
