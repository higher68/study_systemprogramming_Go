package main
import (
    "os"
)

func main() {
    file, err := os.Create("test.txt")
    // os.Createは新規ファイルのオープン
    if err != nil {
        // nilはNULLと一緒
        panic(err)
    }
    file.Write([]byte("os.File example\n"))
    // Writeはバイト列を受け取る
    file.Close()
}

        
