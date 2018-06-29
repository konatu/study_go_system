package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//無限ループの開始
	for {
		//変数bufferを5byteのメモリ容量で作成
		buffer := make([]byte, 5)
		//変数sizeへの入力（5byte）をRead（入力待ち）から受け付ける
		size, err := os.Stdin.Read(buffer)
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		fmt.Printf("size:%d input= %s\n", size, string(buffer))
	}
}
