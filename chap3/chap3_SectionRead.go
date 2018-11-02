package main
import (
    "io"
    "os"
    "strings"
)

func main() {
    reader := strings.NewReader("Example of io.SectionReader\n")
    sectionReader := io.NewSectionReader(reader, 14, 7)  //文字列を一旦ラップしてから渡す
    io.Copy(os.Stdout, sectionReader)  // SectionReader(r, off, n) rを開始オフセットoffから、nバイト後のos.EOFまで読み込む
}
