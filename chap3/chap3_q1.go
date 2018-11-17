package main
import (
    "io"
    "io/ioutil"
    "os"
)
func main() {
    buffer := make([]byte, 1024) // 読み込むデータの格納場所を事前に用意
    filename, err := os.Stdin.Read(buffer)  // バッファに格納
    if err != nil{
        panic(err)
    }
    // bufferは一時的にデータをためておく場所
    file,err := os.Open(string(filename))
    if err != nil{
        panic(err)
    }
    buffer2, err :=  ioutil.ReadAll(file)  // ioutilはバッファを制御できないのでで、ちょい危険
    file_out, err := os.Create("new.txt")
    io.Copy(file_out, string(buffer2))
    defer file.close()
}
