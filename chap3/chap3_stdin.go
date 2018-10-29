package main
import (
    "io"
    "fmt"
    "os"
)
// make([]int, 10, 100) 100個のintを持つ配列を割り当てる。そのあと、10こめまでの要素を示す 初期化はしてくれる
// 実行すると入力待ちになる。Enterが押されるたびに結果が変える
func main() {
    for {
        buffer := make([]byte, 5) // 5byte目まで入力
        size, err := os.Stdin.Read(buffer)
        if err == io.EOF {  // ctrl+Dがio.EOFなんだってさ
            fmt.Println("EOF")
            break
        }
        fmt.Printf("size=%d input='%s'\n", size, string(buffer))
    }
}
            
