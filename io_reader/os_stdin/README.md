# os.Stdin
`Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")`
- 標準入力をするためのファイルを開く変数

# os.Read
`func (*File) Read`
`func (f *File) Read(b []byte) (n int, err error)`

- readはファイルから最大len（b）バイトを読み込みます。
- 読み取られたバイト数とエラーが発生した場合はエラーを返します。
- ファイルの終わりに、Readは0、io.EOFを返します。
- 何も入力してないと入力待ちででブロッキングされる。
