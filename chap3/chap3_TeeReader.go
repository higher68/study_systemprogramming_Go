package main
import (
    "bytes"
    "fmt"
    "io"
    "io/ioutil"
)
func main() {
    var buffer bytes.Buffer
    reader := bytes.NewBufferString("Example of io.TeeReader\n")
    teeReader := io.TeeReader(reader, &buffer)  // 読み込んだ内容を別のio.Writerに書き出す。MultiWriterは書き込んだ内容を書き出す今回だとreaderを読んで、bufferにも書き込んでる
    // データを読み捨てる
    _, _ = ioutil.ReadAll(teeReader)
    // けどバッファに残ってる
    fmt.Println(buffer.String())
}
