package main
import (
    "io"
    "os"
    "net"
)
func main() {
    // 関数内のみ、:=を使って変数宣言時にvarを省略可能
    conn, err := net.Dial("tcp", "ascii.jp:80")
    if err != nil {
        panic(err)
    }
    io.WriteString(conn, "GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n")
    // net.Connがio.Readerインターフェースでもあることを利用、io.Copyで画面に出力
    io.Copy(os.Stdout, conn)
}
