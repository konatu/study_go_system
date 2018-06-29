package main

import (
	"io"
	"os"
)

func main() {
	// filein変数に`os.Open`で開いたmain.goを格納
	filein, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}
	// fileout変数にos.Createで新しいfile.goを格納
	fileout, err := os.Create("file.go")
	if err != nil {
		panic(err)
	}
	// io.Copyを使用してfileinをfileoutにコピー
	io.Copy(fileout, filein)
	defer fileout.Close()
	defer filein.Close()
}
