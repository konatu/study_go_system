# io.Reader
- 外部からデータを読み込むための抽象化されたインターフェース。
```go
type Reader interface {
        Read(p []byte) (n int, err error)
}
```

## インターフェース
```go
type <型名> interface {
    メソッド名(引数の型, ...) (返り値の型, ...)
```

## Readメソッド
`Read(p []byte) (n int, err error)`
- 引数であるpは読み込んだ内容を一時的にバッファしておく場所となります。
- またその際にはあらかじめメモリを用意しておく必要があります。
- このメモリを用意しておくにはmake関数を使うと良い。

*例*
```go
//1024のバッファをmakeで作成する。
buffer := make([]byte, 1024)
//実際に読み込んだバイト数
size ,err := r.Read(buffer)
```

# 組み込みの補助関数
## ioutil.ReadAll()
- これは終端記号に当たるまで全てのデータを読み込みます。
```go
buffer, err := ioutil.ReadAll(reader)
```

## io.ReadFull
- これは決まったバイト数だけ確実に読み込みたい場合に使用する。
```go
buffer, err := make([]byte, 4)
size, err := io.ReadFull(reader, buffer)
```
- またサイズ分まで読み込めない場合はエラーとなる。


# コピー補助関数
## io.Copy
- 全て読み尽くして、書き込む
```go
// 全て読み込む
writesize, err := io.Copy(writer, reader)
// 指定したサイズのみを読み込む
writesize, err := io.CopyN(writer, reader, size)
```
