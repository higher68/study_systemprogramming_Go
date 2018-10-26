package main
import (
    "net/http"
    "io"
)
// http.ResponseWriterは、Webサーバからブラウザにメッセージを書き込むのに使う
func handler(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "http.ResponseWriter sample")
    // io.Writeを使うと、相手がどんなものでも、書き込める
}
func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)  // :8080ポートでWebサーバ起動
}
