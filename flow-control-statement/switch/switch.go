package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linu.")
	default:
		fmt.Printf("%s. \n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow")
	default:
		fmt.Println("Too far away.")
	}

	// 呼び出し元の関数（今回だとmain関数）の終わりまで、関数の実行を遅らせるステートメント
	// なぜ使うのかはわかっていない。
	defer fmt.Println("world")
	fmt.Println("hello, ")

	fmt.Println("counting")

	// deferの呼び出し順はスタック構造で呼び出される。LIFO
	// 本を積み上げて、上から取っていくイメージ
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
