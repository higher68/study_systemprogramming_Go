package main
import (
    "os"
    "io"
    "encoding/csv"
)
func main() {
    file, err := os.Create("q2_2_sample.csv")
    if err != nil {
        panic(err)
    }
    writer := csv.NewWriter(io.MultiWriter(os.Stdout, file))  // NewWriterは、io.Writerを引数に取れば何でも良い。なので、その一つであるMultiWriterで二つに書く
    writer.Write([]string{"hoge", "1"})
    writer.Write([]string{"boge", "2"})
    writer.Write([]string{"foge", "3"})
    writer.Flush() // これを呼び出して初めて書き込まれる
}
