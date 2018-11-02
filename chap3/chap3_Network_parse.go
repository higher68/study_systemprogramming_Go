package main
import (
    "bufio"
    "fmt"
    "io"
    "net"
    "net/http"
    "os"
)
func main() {
    conn, err := net.Dial("tcp", "ascii.jp:80")
    if err != nil {
        panic(err)
    }
    conn.Write([]byte("GET / HTTP/1.0\r\nHos: ascii.jp\r\n\r\n"))
    res, err := http.ReadResponse(bufio.NewReader(conn), nil)  // bufio.Readerでnet.connをラップすると、http.Response構造体のオブジェクトが返ってくる。
    // HTTPヘッダなどに分解されているため、返ってくる構造体はすごく扱いやすいよ
    // ヘッダーを表示してみる
    fmt.Println(res.Header)
    // ボディーを表示してみる。最後にはclose()すること
    defer res.Body.Close()
    io.Copy(os.Stdout, res.Body)
}
