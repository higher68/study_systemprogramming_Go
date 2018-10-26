package main
import (
    "bytes"
    "fmt"
)
func main() {
    var buffer bytes.Buffer  //interfaceの型を持つvariable定義
    // Writeメソッドで書き込まれる内容を貯めて、あとでまとめて結果を受け取れる
    buffer.Write([]byte("bytes.Buffer example\n"))
    fmt.Println(buffer.String())
}
