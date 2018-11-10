package main
import (
    "bufio"
    "fmt"
    "io"
    "strings"
)

var source = `1行め
2行目
3行目 `

func main() {
    reader := bufio.NewReader(strings.NewReader(source))
    for {
        line, err := reader.ReadString('\n')   // \nで分割する
        fmt.Printf("%#v\n", line)  // %#vだと"で囲った文字列が表示される
        if err == io.EOF {
            break
        }
    }
}
