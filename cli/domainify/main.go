package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var tlds = []string{"com", "net"}

const allowedChars = "abcdefghijklmnopqrstuvwxyz0123456789_-"

func main() {
	rand.Seed(time.Now().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		text := strings.ToLower(s.Text())
		var newText []rune
		for _, r := range text {
			if unicode.IsSpace(r) {
				r = '-'
			}
			if !strings.ContainsRune(allowedChars, r) {
				continue
			}
			newText = append(newText, r)
		}
		fmt.Println(string(newText) + "." + tlds[rand.Intn(len(tlds))])
	}
}

// 文字列に対してrangeを実行すると、それぞれの文字の位置を表すインデックス値と、
// 文字を数値として表したrune型の値(int32のエイリアス)の2つが返される。
// インデックスは、文字列の中の何バイト目かを表している。
// GOでの文字列はUTF-8と解釈されるので、インデックスの値は、何文字目かを表すものではない。

// go build -o domainify
// go build コマンドだけだと、現在のフォルダ名と同名のコマンドがフォルダに作られる
// go install とすると、$GOPATH/binに作られる
