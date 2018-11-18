package main
import (
    "io"
    "io/ioutil"
    "os"
    "fmt"
    // "strings"
)
func main() {
    filename := ""
    for {buffer := make([]byte, 1) // 読み込むデータの格納場所を事前に用意
        size, err := os.Stdin.Read(buffer)  // バッファに格納
        if string(buffer) == "\n"{
            fmt.Println(err)
            break
        }
        filename += string(buffer)
        fmt.Println(size)
    }
    fmt.Println(filename)
    // fmt.Println("hoge")
    // fmt.Println(string(buffer))
    //fmt.Println("chap3_q1.go")
    // fmt.Println(strings.TrimRight(string(buffer), "\n"))
    // fmt.Println(len(strings.TrimRight(string(buffer), "\n")))
    // bufferは一時的にデータをためておく場所
    // filename := string(buffer)
    // file,err := os.Open(strings.TrimRight(filename, "\n"))
    file,err := os.Open(filename)
    if err != nil{
        panic(err)
    }
    buffer2, err :=  ioutil.ReadAll(file)  // ioutilはバッファを制御できないのでで、ちょい危険
    file_out, err := os.Create("new.txt")
    io.WriteString(file_out, string(buffer2))
    // defer file_out.close()
}
