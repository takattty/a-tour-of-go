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

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	fmt.Println(&a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(primes)
	fmt.Println(s)
}
