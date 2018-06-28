# os package
`func Create(name string) (*File, error)`
- 権限666で作成する。
- すでにあったらスキップ
- 成功時は返されたFileのメソッドをI/O に使用できる。(os.File)

# io package
`func MultiWriter(writers ...Writer) Writer`
- Unix teeコマンドのように全てを書き込む
- 書き込まれる際は一度に一行づつ書き込まれ
  全体の中で１つエラーが発生するとその処理も止まる。

`func WriteString(w Writer, s string) (n int, err error)`
- WriteStringは文字列sの内容をwに書き込みます。
- バイトスライスを受け入れます。
- wがWriteStringメソッドを実装している場合は、直接呼び出されます。
- それ以外の場合は、w.Writeは1回だけ呼び出されます。
