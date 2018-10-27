package main
import (
    "os"
    "encoding/json"
)
func main() {
    encoder := json.NewEncoder(os.Stdout)  // json形式の構造体呼び出し？out=stdout
    encoder.SetIndent("", "   ")
    encoder.Encode(map[string]string{
        "example": "encoding/json",
        "hello": "world",  // jsonのときはここにも,がないとダメなのか。
    })
}
