package main
import (
    "compress/gzip"
    "net/http"
    "encoding/json"
    "io"
    "os"
)
// goではfuncはfunction(引数 型, 引数 型, ) 返り値の型
func handler(w http.ResponseWriter,  r *http.Request) {
    // ResponseWriter はhttpレスポンスを作るのに使う
    w.Header().Set("Content-Encoding", "gzip")  // よくわからんから先進む
    w.Header().Set("Content-type", "application/json")
    // json化する前のデータ
    // mapは値をkey-valueで保存するデータ構造
    // map[keyの型]valueの型って感じで宣言
    source := map[string]string{"Hello": "World",
    }
    file1, err1 := os.Create("test2_3.json")
    file2, err2 := os.Create("test2_3.json.gz")
    if err1 != nil {
        panic(err1)
    }
    // NewEncoderは、新しいEncoder 構造体を作る。下では引数渡してる
    encoder := json.NewEncoder(io.MultiWriter(os.Stdout, file1))
    encoder.SetIndent("", "   ")  // prefixはよーわからんが、2ndargは使われるインデント
    encoder.Encode(source)
    // func (enc *Encoder) Encode(v interface{})
    // v(=入力したもの)をエンコードして書き込むっぽい
    if err2 != nil {
        panic(err2)
    }
    writer := gzip.NewWriter(io.MultiWriter(os.Stdout, file2))
    writer.Header.Name = "test2_3.json"
    // Header構造体は、圧縮されるファイルのメタを格納する
    writer.Close()
    
    
}
// Http系は、今の所無視。またやるらしいし
func main() {
    http.HandleFunc("/", handler)
    // HTTP、HTTPハンドラを指定。ハンドラは反応のこと
    http.ListenAndServe(":8080", nil)
    // listen：サーバからの待ち受け。
    // listenandserveは、指定したポートでhttpサーバに待受を開始させる、
    // 2ndargはハンドラらしい。よくわからん正直。基本nilらしい}
    
