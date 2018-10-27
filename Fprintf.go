package main
import (
    "os"
    "fmt"
    "time"
)
func main() {
    // 整形したデータをio.Writerへ書き出す汎用関数
    fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v", time.Now())
}

