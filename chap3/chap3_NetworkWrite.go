package main
import (
    "io"
    "os"
    "net"
)

func main() {
    conn, err := net.Dial("tcp", "ascii.jp:80")  // net.dialでnet.conn型を返す
    // httpの通信そのものが読み込まれる。
    if err != nil {
        panic(err)
    }
    conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
    io.Copy(os.Stdout, conn)  // 標準出力にコピーしてデータを一括でコピー
}

        
    
