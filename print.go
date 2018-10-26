package main
import (
    "os"
)
// fmt.Printlnで呼び出す、OS.stdout.Writeを呼び出す
// なんかよくわからんが、"はよくて、'は囲う時に使っちゃダメぽい
func main() {
    os.Stdout.Write([]byte("os.Stdout extmple\n"))
}
    
    
