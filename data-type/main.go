package main

import "fmt"

// 構造体。型の異なるデータ型の変数を集めたデータ構造。
// フィールド = 各変数
type Vertex struct {
	X, Y int
}

func main() {
	// ポインタは値のメモリアドレスを指す
	// つまり、値がメモリのどの位置にいるかの情報を指している。
	i, j := 42, 2701

	// 変数iのポインタをpに代入。&オペレータが、ポインタを引き出している。
	// iの値がどこにあるかの住所(つまりメモリのどこにいるか)のイメージ
	p := &i
	// *オペレータでポインタの指す先の変数を表示
	// *オペレータは変数じゃなくてポインタの指す先の変数を表示する
	// ↓invalid operation: cannot indirect i (variable of type int)compilerInvalidIndirection
	// fmt.Println(*i)
	fmt.Println(p)
	fmt.Println(&p)
	fmt.Println(*p)
	// ポインタの指す先の変数を変更している。（同じポインタ）
	*p = 21
	// *オペレータ無しでのポインタへの代入は出来ない
	// p = 30
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

	v := Vertex{1, 2}
	// フィールドの参照や代入は.を用いる
	v.X = 4
	vp := &v
	fmt.Println(v.X)
	fmt.Println(vp)
	fmt.Println(&vp.X)
	fmt.Println(&vp.Y)

	// 配列の長さを変える事は出来ない。固定長
	// が変えたい時もある。その場合はスライスという可変長のデータ構造を使う。
	// 全て同じ型の要素を含んでいる。
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	fmt.Println(&a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	// []T がスライスの型を表す。のでsはスライスってこと。
	// もちろん配列のインデックスを指す。
	// 最初の要素は含むが、最後の要素は除いた半開区間で選択。
	// なので↓は1番目から3番目までの値をスライスとして切り出している
	var s []int = primes[1:4]
	fmt.Println(primes)
	fmt.Println(s)

	// スライスは配列への参照のようなものです。
	// スライスはどんなデータも格納しておらず、単に元の配列の部分列を指し示しています。
	// スライスの要素を変更すると、その元となる配列の対応する要素が変更されます。
	// 同じ元となる配列を共有している他のスライスは、それらの変更が反映されます

	// スライスは長さと容量を持っている
	// 長さは、スライスに含まれる要素の数
	// 容量は、スライスの最初の値から元の配列の要素数を合わせた数
}
