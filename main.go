package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
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

// struct 構造体
type Task struct {
	Title    string
	Estimate int
}

// センチネルエラー
// 特定のエラーを事前に定義しておくことで、エラー検知時に定義したエラーを細かく特定、キャッチすることができる
var ErrCustom = errors.New("not found")

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

	// // slice
	// //　配列の定義 [要素数]要素の型 で初期化 intの配列は要素数分の 0 が初期値に入る
	// var a1 [3]int
	// var a2 = [3]int{1, 2, 3}
	// a3 := [...]int{1, 2} // [...] と記述すると代入時の要素数を自動カウントしてくれる

	// fmt.Printf("%v, %v, %v\n", a1, a2, a3)
	// fmt.Printf("%v, %v\n", len(a3), cap(a3)) // cap()は中身の要素の型によって出力結果が異なる、配列の場合は要素数でlen()と同じ
	// fmt.Printf("%T, %T\n", a2, a3)           // 配列の型は中身の型と要素数によって決まる、同じ数値配列でも要素数が異なると別々の型となる

	// // slice 配列は一度宣言したら要素数の追加等できない、あとから要素を追加したりする場合は slice型を使う
	// var s1 []int // []の中身空で slice の宣言
	// s2 := []int{}
	// fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1))
	// fmt.Printf("s2: %[1]T %[1]v %v %v\n", s2, len(s2), cap(s2))
	// fmt.Println(s1 == nil)   // var で宣言された slice は nil 扱いだが
	// fmt.Println(s2 == nil)   // := で宣言された slice は nil ではない
	// s1 = append(s1, 1, 2, 3) // append(slice変数, 要素1, 要素2, ~) で既存slice に要素を追加する
	// fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1))
	// s3 := []int{4, 5, 6}
	// s1 = append(s1, s3...) // sliceに別のsliceを追加するときは append(追加元slice, 取り込まれるslice...) のように第二引数の slice に ...（スリードット）を繋げる
	// fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1))

	// s4 := make([]int, 0, 2) // makeメソッドを使用してsliceを宣言する make(型情報, 要素数, 容量)
	// fmt.Printf("s4: %[1]T %[1]v %v %v\n", s4, len(s4), cap(s4))
	// s4 = append(s4, 1, 2, 3, 4)
	// fmt.Printf("s4: %[1]T %[1]v %v %v\n", s4, len(s4), cap(s4))

	// s5 := make([]int, 4, 6)
	// fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))
	// fmt.Printf("%v\n", s5[1:3])
	// s6 := s5[1:3] // slice[index:index] の指定で 左のindex<= 対象 < 右のindex の範囲の値を抽出できる 今回だとindex:1番目以上3番未満なので、要素の2番目、3番目の2つを抽出
	// fmt.Printf("s6: %[1]T %[1]v %v %v\n", s6, len(s6), cap(s6))
	// s6[1] = 10
	// fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))
	// fmt.Printf("s6: %[1]T %[1]v %v %v\n", s6, len(s6), cap(s6))
	// // sliceの一部を別のsliceに代入するとそのslice同士はメモリを共有することになる
	// // 新しく宣言したsliceに値を追加すると、元のsliceの中身も書き換えられる
	// s6 = append(s6, 20)
	// fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5)) // s6に要素を追加するとs5の4番目（index 4）の値も書き換わる
	// fmt.Printf("s6: %[1]T %[1]v %v %v\n", s6, len(s6), cap(s6))

	// // map: key-valueの組み合わせ、C#のDictionaryみたいな
	// var m1 map[string]int  // map[keyの型]valueの型 でmapの宣言
	// m2 := map[string]int{} // := での宣言時は valueの型の後ろに空の {} をつける
	// fmt.Println(m1 == nil) // slice 同様 var で宣言された map は nil だが
	// fmt.Println(m2 == nil) // := で宣言されたものは nil とならない

	// m2["A"] = 10 // key とそれに対応するvalueのセット
	// m2["B"] = 20
	// m2["C"] = 30
	// fmt.Printf("m2: %[1]T %[1]v %v %v\n", m2, len(m2), m2["A"])
	// delete(m2, "B")
	// fmt.Printf("m2: %[1]T %[1]v %v %v\n", m2, len(m2), m2["B"]) // 存在しないkeyを呼び出すと0を返すが存在するkeyの値か、keyが存在せず返した値かは見分けられる

	// value, exist := m2["A"] // keyを指定して値を取り出す時の第二戻り値にそのkeyが存在するかどうかを判定するbool値がある
	// fmt.Printf("%v, %v\n", value, exist)
	// value, exist = m2["B"] // bool値の値がtrue なら存在する値を取り出した、falseならkeyが存在せず0を返した、と判定できる
	// fmt.Printf("%v, %v\n", value, exist)

	// for k, v := range m2 {
	// 	fmt.Printf("%v, %v\n", k, v)
	// } // for ～ range map {} で map 内の要素をすべて取り出す、map はハッシュマップとなっているので順番は保障されていない

	// struct型の実体化
	// task1 := Task{
	// 	Title:    "Learn Golang",
	// 	Estimate: 3,
	// }
	// task1.Title = "Learning Go"
	// fmt.Printf("%[1]T %+[1]v, %v\n", task1, task1.Title)

	// var task2 Task = task1
	// task2.Title = "new"
	// fmt.Printf("task1: %v task2: %v\n", task1.Title, task2.Title)

	// task1p := &Task{ // 初期化時の型の先頭に & をつけてポインタを取得
	// 	Title:    "Learn concurrency",
	// 	Estimate: 2,
	// }
	// fmt.Printf("task1p: %T %+v %v\n", task1p, *task1p, unsafe.Sizeof(task1p)) // %Tで型取得、変数の頭に * をつけてポインタが示すアドレスに格納されている値を取得（この場合は直前で初期化された構造体）

	// (*task1p).Title = "Changed" // デリファレンスで構造体にアクセスして値を書き換えることが可能
	// fmt.Printf("task1p: %+v\n", *task1p)

	// task1.extendEstimate()
	// fmt.Printf("task1 value receiver: %+v\n", task1.Estimate)

	// task1.extendEstimatePointer()
	// fmt.Printf("task1 pointer receiver: %+v\n", task1.Estimate) // ポインタレシーバーを受け取る関数の処理では実体の値が変わっている

	// funcDefer()

	// files := []string{"file1.csv", "file2.csv", "file3.csv"}
	// fmt.Println(trimExtension(files...))

	// name, err := fileChecker("file.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(name)

	// // 無名関数
	// sample := 2
	// func(i int) { // 関数の名前を定義していない
	// 	fmt.Println(i)
	// }(sample) // 定義直後に()と引数を渡すことで即時実行される

	// noNameFunc := func(i int) int { // 無名関数を変数に渡している
	// 	return i + 10
	// } // 定義直後に()を使用していないので即時実行していない
	// fmt.Println(noNameFunc(sample)) // 無名関数を格納した変数に()と引数を渡して実行される

	// // 引数として使われる無名関数の定義
	// fn := func(file string) string {
	// 	return file + ".csv"
	// }
	// // 無名関数をaddExtensionに渡してaddExtensionを実行
	// addExtension(fn, "sample")

	// fn2 := multiply()   // 無名関数を返り値とする関数を呼び出すと、返り値の無名関数が返ってくる（小泉構文）
	// fmt.Println(fn2(2)) // 格納された変数に()と引数を渡すと無名関数が実行される

	// f3 := countUp() // f3 にcountUp()の返り値である無名関数が格納されている
	// for i := 1; i <= 5; i++ {
	// 	v := f3(2)
	// 	fmt.Printf("%v\n", v)
	// }

	// var stringp *string
	// string := "hello world"
	// stringp = &string            // ポインタ取得
	// fmt.Printf("%p\n", stringp)  // ポインタ出力
	// fmt.Printf("%v\n", *stringp) // デリファレンスにより値取得

	// // vehicle構造体の実体化
	// // ポインタ変数を使用して定義されたメソッドを呼び出すので変数にポインタを格納している
	// v := &vehicle{0, 5}

	// // speedUpAndDown() は controllerインターフェースを引数にとる関数
	// // vehicle構造体は controllerインターフェースに定義された関数を実装しており
	// // controllerインターフェース を実装しているとみなされる
	// // そのためspeedUpAndDown()の引数としてvehicle構造体の実体が使用可能となる
	// speedUpAndDown(v)

	// b := &bycycle{0, 5}
	// speedUpAndDown(b)

	// // String() を定義していない状態で変数に格納された値を出力すると値がそのままでる？
	// // Stringerインターフェースの String()を実装した状態で出力すると
	// // Println()内でStringerインターフェースの実装有無を判断し、実装している場合はその実装内容が出力される
	// // 以下の出力結果は、String()　を実装したいない状態だと「&{0,5}」だが
	// // String()を実装している状態だとその返り値が出力される
	// fmt.Println(v)

	// if
	// 	a := -1
	// 	if a == 0 {
	// 		fmt.Println("zero")
	// 	} else if a > 0 {
	// 		fmt.Println("positive")
	// 	} else if a < 0 {
	// 		fmt.Println("negative")
	// 	}

	// 	// for
	// 	for i := 0; i < 5; i++ {
	// 		fmt.Println(i)
	// 	}

	// 	// var i int
	// 	// for {
	// 	// 	if i > 3 {
	// 	// 		break
	// 	// 	}
	// 	// 	fmt.Println(i)
	// 	// 	i++
	// 	// 	time.Sleep(2 * time.Second)
	// 	// }

	// loop: // 直下のfor文に名前を付けている
	// 	for i := 0; i < 10; i++ {
	// 		switch i {
	// 		case 2:
	// 			continue //2,3のときは次の処理へスキップする
	// 		case 3:
	// 			continue
	// 		case 9:
	// 			break loop // 9 のときに 単純な break だけを書いてしまうと switch文を抜けることになるが名前を指定することで一番外のfor文から抜けることを明示している
	// 		default:
	// 			fmt.Printf("%v\n", i)
	// 		}
	// 	}

	// 	// for ~ range
	// 	items := []item{
	// 		{price: 10},
	// 		{price: 20},
	// 		{price: 30},
	// 	}

	// 	for _, i := range items { // for ~ range で slice型の中身をひとつづつ取り出してコピーを渡している(1つ目の変数はインデックス、2つ目の変数はそのインデックスの位置にある要素の「コピー」)
	// 		i.price *= 1.1
	// 	}
	// 	fmt.Printf("%+v\n", items) // for ~ range で要素を受け取る変数には実体のコピーが渡されているため i の値が変わっても実体に影響はない（値は定義時のまま）

	// 	for i := range items {
	// 		items[i].price *= 1.1 // 実体の値を変えたい場合は インデックスを使って実体に直接アクセスしたうえで処理を行う
	// 	}
	// 	fmt.Printf("%+v\n", items)

	// error
	error01 := errors.New("something wrong") // errors.New で errors.errorString構造体のアドレスを渡している（ポインタ）
	error02 := errors.New("something wrong")
	fmt.Printf("%[1]p, %[1]T, %[1]v\n", error01)
	fmt.Println(error01.Error())
	fmt.Println(error01 == error02)

	// errorsのラップ
	err0 := fmt.Errorf("add info: %w", error01) // %W を使って *errors.errorString をラップしている ラップ後の型は *fmt.wrapError となる
	fmt.Printf("%[1]p, %[1]T, %[1]v\n", err0)
	fmt.Println(errors.Unwrap(err0)) // *fmt.wrapError 型に対しては errors.Unwrap() を使えるため Unwrap により元の errors.errorString を取り出すことができる
	fmt.Printf("%T\n", errors.Unwrap(err0))

	err1 := fmt.Errorf("add info: %v", errors.New("original error"))
	fmt.Printf("%[1]v, %[1]T\n", err1) // %v で wrap した *errors.errorString の型はそのまま
	fmt.Println(errors.Unwrap(err1))   // *errors.errorString に対する Unwrap() はない、このように Unwrapすると nil を返す

	// センチネルエラー を何段階かラップする
	err2 := fmt.Errorf("in repository layer: %w", ErrCustom)
	fmt.Println(err2)
	err2 = fmt.Errorf("in service layer: %w", err2)
	fmt.Println(err2)

	// ラップされたセンチネルエラーとセンチネルエラーを直接比較することはできない
	// errors.Is()を使用することでラップを自動的にアンラップして判定してくれる
	if errors.Is(err2, ErrCustom) {
		fmt.Println("matched")
	}

	file := "dummy.txt"
	err3 := fileChecker(file)
	if err3 != nil {
		if errors.Is(err3, os.ErrNotExist) {
			fmt.Println(err3)
			fmt.Printf("%v is not found\n", file)
		} else {
			fmt.Println("unknown error")
		}
	}
}

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in checker: %w", err)
	}
	defer f.Close()
	return nil
}

type item struct {
	price float32
}

// receiver
// 値レシーバーを持つ関数
// Goではある型の実体に対する関数を定義できる
// そのメソッドが受け取る型の実体をレシーバーと呼ぶ
// 値レシーバーを受け取る関数で実体の値を変えても、実際は受け取ったレシーバーのコピーに対して変更を加えるため、実体には影響がない
func (task Task) extendEstimate() {
	task.Estimate += 10
}

// 実体に変更を加えたい場合はポインタレシーバーを使用する
func (taskp *Task) extendEstimatePointer() {
	taskp.Estimate += 10 // デリファレンスで実体にアクセスして、Estimateの値を変更している (*taskp).Estimate
}

// function defer
// defer をつけた処理は関数が終了する直前に呼ばれる
// 複数deferがある場合は下→上の順番で処理が実行される
func funcDefer() {
	defer fmt.Printf("main func final-finish\n")
	defer fmt.Printf("main func semi-finish\n")
	fmt.Printf("hello world\n")
}

func trimExtension(files ...string) []string { // 引数の型指定の頭に「...」をつけることでその型の任意の数の引数を受け取ることができる
	out := make([]string, 0, len(files))
	for _, f := range files {
		out = append(out, strings.TrimSuffix(f, ".csv"))
	}
	return out
}

// func fileChecker(name string) (string, error) { // 引数を複数返す関数を定義する場合は丸カッコとカンマで引数の型を指定する
// 	f, err := os.Open(name) // os.Open()で指定ファイルを開く、第一戻り値には開いた際の参照、第二戻り値には開けなかった場合のエラーが格納される
// 	if err != nil {
// 		return "", errors.New("file not found") // 何らかのエラーが発生した場合は空文字とエラーメッセージを返している
// 	}
// 	defer f.Close()

// 	return name, nil
// }

// 無名関数を引数として受け取る関数の定義
// 以下の関数 addExtensionは 文字列を受け取り文字列を返す無名関数と 任意の文字列を引数にとる関数として定義
func addExtension(f func(file string) string, name string) {
	fmt.Println(f(name)) // 第一引数の無名関数に第二引数で受け取った文字列を渡して呼び出している
}

// 無名関数を返り値とする関数の定義
func multiply() func(int) int {
	return func(i int) int {
		return i * 1000
	}
}

// closure
// ある変数を関数内に格納することで、その関数以外の影響を受けないようにする
func countUp() func(int) int {
	count := 0
	return func(n int) int {
		count += n
		return count
	}
}

// interface
// 型の一種 メソッドを定義できる
type controller interface {
	speedUp() int
	speedDown() int
}

type vehicle struct {
	speed       int
	enginePower int
}

type bycycle struct {
	speed      int
	humanPower int
}

// ポインタレシーバーを使用して構造体の実体に対するメソッドを定義
// vehicle構造体の実体のメソッドとしてspeedUp()とspeedDown()を実装
// あるinterfaceに定義されているメソッドをすべて実装した型は
// 自動的にそのinterfaceを実装したことになる
// vehicle や bycycle は controllerインターフェースを実装したことになる
func (v *vehicle) speedUp() int {
	v.speed += 10 * v.enginePower
	return v.speed
}

func (v *vehicle) speedDown() int {
	v.speed -= 5 * v.enginePower
	return v.speed
}

func (v vehicle) String() string {
	return fmt.Sprintf("このvehicleのスピードは: %v (エンジンパワーは %v)", v.speed, v.enginePower)
}

func (b *bycycle) speedUp() int {
	b.speed += 3 * b.humanPower
	return b.speed
}

func (b *bycycle) speedDown() int {
	b.speed -= 1 * b.humanPower
	return b.speed
}

// controllerインターフェースを引数にとる関数を定義
// 関数内で speedUp() と speedDown() を実行している
func speedUpAndDown(c controller) {
	fmt.Printf("current speed: %v\n", c.speedUp())
	fmt.Printf("current speed: %v\n", c.speedDown())
}
