package main
import (
    "encoding/binary"
    "fmt"
    "os"
    "io"
    "hash/crc32"
)

func dumpChunk(chunk io.Reader) {
    var length int32
    binary.Read(chunk, binary.BigEndian, &length)
    buffer := make([]byte, 4)
    chunk.Read(buffer)
    fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)
}

func readChunks(file *os.File) []io.Reader {
    // チャンクを格納する配列
    var chunks []io.Reader

    // 最初の8バイトを飛ばす
    file.Seek(8, 0)
    var offset int64 = 8
    for {
        var length int32
        err := binary.Read(file, binary.BigEndian, &length)
        if err == io.EOF {
            break
        }
        chunks = append(chunks, io.NewSectionReader(file, offset, int64(length)+12))
        // 次のチャンクの先頭に移動
        // 現在位置は長さを読み終わった箇所なので
        // チャンク名(4バイト) + データ長 + CRC(4バイト)先に移動
        offset, _ = file.Seek(int64(length+8), 1)
    }
    return chunks
}

func textChunk(text string) io.Reader {
    byteData := []byte(text)
    var buffer bytes.Buffer
    binary.Write(&buffer, binary.BigEndian, int32(len(byteData)))
    buffer.WriteString("tEXt")
    buffer.Write(byteData)
    // CRCを計算して追加
    crc := crc32.NewIEEE()
    binary.Write(&buffer, binary.BigEndian, crc.Sum32())
    return &buffer
}

func main() {
    file, err != os.Open("Lenna.png")
    if err != nil{
        
