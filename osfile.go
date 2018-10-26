package main
import (
    "os"
    "compress/gzip"
    "io"
)
func main() {
    file, err := os.Create("test.txt.gz")
    if err != nil {
        panic(err)
    }
    writer := gzip.NewWriter(file)  //  writerに書き込まれたデータをgzip圧縮 os.fileに中継
    writer.Header.Name = "test.txt"  
    io.WriteString(writer, "gzip.Writer example\n")
    writer.Close()
}
