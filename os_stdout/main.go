package main

import (
	"os"
)

func main() {
	os.Stdout.Write([]byte("os.stdout example\n"))
}
