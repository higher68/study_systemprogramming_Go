package main
import (
    "bytes"
    "encoding/binary"  // 任意の実行環境のエンディアンを現在の実行環境のエンディアンに格納できるパッケージ
    "fmt"
)
// エンディアン変換
// CPUはリトルエンディアン
// リトルエンディアンでは、1000という数値があったとき、小さい桁からメモリに格納する
// ビッグエンディアン：大きい桁からメモリに格納される(ネットワークバイトオーダー)とも呼ぶ
func main() {
    // 32ビットのビッグエンディアンのデータ(10000)
    data := []byte{0x0, 0x0, 0x27, 0x10}
    var i int32
    // エンディアン変換
    binary.Read(bytes.NewReader(data), binary.BigEndian, &i)  // dataを読んでbigendianが変換されてiに格納される。
    fmt.Printf("data: %d\n", i)
}
    
