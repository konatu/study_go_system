# os.File
- ファイルからの入力にはos.File構造体を使います。
- 新規作成はos.Create()
- 既存のファイルを開くのはos.Open()
- 内部的には上記2つはos.OpenFile()関数のフラグを違いのエイリアス
```
func Open(name string) (*File, error) {
　　　return OpenFile(name, O_RDONLY, 0)
}
　
func Create(name string) (*File, error) {
　　　return OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
}
```

# os.Stdout
- 標準出力をするためのファイルを開く変数
`Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")`
