package main
import "fmt"

// define interface　Goだと~erのような名前をつけることが多い。
// interfaceは、メソッドの型だけを指定したもの
type Talker interface {
    Talk()
}

// 構造体を定義 複数の型のデータを一つにまとめられるもの
type Greeter struct {
    name string
}

// 構造体はTalkerインターフェースで定義されるメソッドを持つ
// メソッドは func レシーバ値 レシーバ型となる。レシーバは
// g Greetという感じで指定すると、funcの定義内でg.nameみたいな感じで使えるってことか？
// オーバーライドに近い？g = Greeterに、Talk()を入れてる？
// class Greeter:
//   def Talk(self):
// のselfがgに相当してるのかしら
func (g Greeter) Talk() {
    fmt.Printf("Hello, my name is %s\n", g.name)
}

func main() {
    // インターフェースの型を持つ変数を宣言
    var talker Talker
    // インターフェースを満たす構造体のポインタは代入可能
    talker = &Greeter{"wozozo"}
    talker.Talk()
}
