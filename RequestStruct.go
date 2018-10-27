package main
import (
    "os"
    "net/http"
)
func main() {
    request, err := http.NewRequest("GET", "http://ascii.jp", nil)
    if err != nil {
        panic(nil) // 後続の処理を中止し、警告
    }
    request.Header.Set("X-TEST", "ヘッダーも追加できます")
    request.Write(os.Stdout)
}
