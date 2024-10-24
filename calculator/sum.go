package calculator

import "fmt"

// 小文字の変数、関数は同一package内で共有可能
var offset float64 = 1

// 大文字の変数、関数はpackage間を超えて呼び出し可能
var Offset float64 = 1

func Sum(a float64, b float64) float64 {
	fmt.Println("同一パッケージ内の小文字メソッド呼び出し：", multiply(1, 2))
	return a + b + offset
}
