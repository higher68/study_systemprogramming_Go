package main
import (
    "os"
    "io"
)
// os.File構造体を使ってファイルからの入力をする
// os.Create(), os.Open()はos.OpenFileという関数のエイリアスで、同じシステムコールが呼ばれている
func main() {
    file, err := os.Open("chap3_WriteFile.go")
    if err != nil {
        panic(err)
    }
    defer file.Close()  // ファイルを一度開いたらcloseする必要ある
    // deferは現在のスコープが終了したらその後ろに書いてある行の処理を実行する
    io.Copy(os.Stdout, file)  // 標準出力にファイルの内容を全て書き出す
    file, err = os.Open("WriteFile.go")
    if err != nil {
        panic(err)
    }
    defer file.Close()  // ファイルを一度開いたらcloseする必要ある
    io.Copy(os.Stdout, file)  // 標準出力にファイルの内容を全て書き出す
}
