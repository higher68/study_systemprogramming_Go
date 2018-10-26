package main
import (
    "io"
    "os"
)
func main() {
    file, err := os.Create("multiwriter.txt")
    if err != nil {
        panic(err)
    }
    writer := io.MultiWriter(file, os.Stdout)  // 複数のio.Writerに同時に書き込む
    // 今回はfileとstd
    io.WriteString(writer, "io.MultiWriter example\n")
}
