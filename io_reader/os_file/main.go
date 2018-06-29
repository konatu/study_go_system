package main

import (
	"io"
	"os"
)

func main() {
	// os.Openで既存のファイルをfile変数に格納
	file, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, file)
	// defer=終了時の挙動
	defer file.Close()
}
