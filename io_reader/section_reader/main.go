package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("Example of io.SectionReader\n")
	section := io.NewSectionReader(reader, 14, 7)
	io.Copy(os.Stdout, section)
}
