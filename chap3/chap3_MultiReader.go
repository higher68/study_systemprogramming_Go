package main

import (
    "bytes"
    "io"
    "os"
)

func main() {
    header := bytes.NewBufferString("----- HEADER -----\n")
    content := bytes.NewBufferString("Example of io.MultiReader\n")
    footer := bytes.NewBufferString("----- FOOTER -----\n")

    reader := io.MultiReader(header, content, footer)
    // 上記、全てのreaderを繋げた出力が表示される
    io.Copy(os.Stdout, reader)
}
