package main
import (
    "io"
    "fmt"
    "os"
)
func main() {
    buffer := make([bytes])
    os.Stdin.Read(buffer)
    file, err := os.Create("new.txt")
    defer file.close()
    io.Copy(os.Stdout, file)
}
