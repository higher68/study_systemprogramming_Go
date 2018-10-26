package main
import (
    "bufio"
    "os"
)
func main() {
    buffer := bufio.NewWriter(os.Stdout)  // bufio.Writerは呼び出したら前に呼び出したところから現在まで書き出す。サイズ指定して、ある一定のサイズ以上書いたら書き出すことも可能
    buffer.WriteString("bufio.Writer ")
    buffer.Flush()  // Flush()を呼ぶと、後続のio.Writerに書き出す
    buffer.WriteString("example\n")
    buffer.Flush()
}
