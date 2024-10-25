package main

import "fmt"

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
	i := 2
	ui := uint16(2)
	fmt.Println("iの値：", i)
	fmt.Printf("i: %v %T\n", i, i)
	fmt.Printf("i: %[1]v %[1]T ui: %[2]v %[2]T\n", i, ui)

	// 複数の変数を同時に定義可能
	pi, title := 3.14, "Go"
	fmt.Printf("pi: %v title: %v\n", pi, title)

	x := 10
	y := 1.23
	z := float64(x) + y
	fmt.Println(z)

	fmt.Printf("Mac: %v, Windows: %v, Linux: %v\n", Mac, Windos, Linux)

	fmt.Printf("a: %v, b: %v, c: %v", a, b, c)
}
