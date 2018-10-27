package main
import (
    "os"
    "fmt"
)
func main() {
    fmt.Fprintf(os.Stdout, "int:%d\n", 1234)
    fmt.Fprintf(os.Stdout, "string:%s\n", "hogeo")
    fmt.Fprintf(os.Stdout, "float:%f\n", 1.23456)
}
