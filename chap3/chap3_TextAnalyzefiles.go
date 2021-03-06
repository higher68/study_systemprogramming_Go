package main
import (
    "encoding/csv"
    "fmt"
    "io"
    "strings"
)
// `で囲まれた部分はGoの文字列リテラルル・・・改行やスペースが意味をもつ
var csvSource =
    `13101,"102  ","1020072","ﾄｳｷﾖｳﾄ","ﾁﾖﾀﾞｸ","ｲｲﾀﾞﾊﾞｼ","東京都","千代田区","飯田橋",0,0,1,0,0,0
13101,"102  ","1020082","ﾄｳｷﾖｳﾄ","ﾁﾖﾀﾞｸ","ｲﾁﾊﾞﾝﾁﾖｳ","東京都","千代田区","一番町",0,0,0,0,0,0
13101,"101  ","1010032","ﾄｳｷﾖｳﾄ","ﾁﾖﾀﾞｸ","ｲﾜﾓﾄﾁﾖｳ","東京都","千代田区","岩本町",0,0,1,0,0,0
13101,"101  ","1010047","ﾄｳｷﾖｳﾄ","ﾁﾖﾀﾞｸ","ｳﾁｶﾝﾀﾞ","東京都","千代田区","内神田",0,0,1,0,0,0
13101,"100  ","1000011","ﾄｳｷﾖｳﾄ","ﾁﾖﾀﾞｸ","ｳﾁｻｲﾜｲﾁﾖｳ","東京都","千代田区","内幸町",0,0,1,0,0,0`

func main() {
    reader := strings.NewReader(csvSource)
    csvReader := csv.NewReader(reader)
    for {
        // fmt.Println("hoge1")
        line, err := csvReader.Read()  // csv.NewReaderはReadを呼ぶと、行の情報を返す
        if err == io.EOF {
            break
        }
        // fmt.Println("hoge")
        fmt.Println(line[2], line[6:9])
    }
}

    
