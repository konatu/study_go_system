package main

import (
  "io"
  "os"
  "fmt"
  "encoding/binary"
)

func dumpChunk(chunk io.Reader){
  var length int32
  binary.Read(chunk, binary.BigEndian, &length)
  buffer := make([]byte 4)
  chunk.Read(buffer)
  fmt.Println("chunk '%v'(%d bytes)\n", string(buffer), length)
}

func readChunk(file *os.File)[]io.Reader{
  //チャンクを格納する配列
  var chunks io.Reader

  file.Seek(8, 0)
  var offset int64 = 8

  for {
    var length int32
    err := binary.Read(file, order, data)
  }
}
