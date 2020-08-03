package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func duplicateVowel(word []byte, i int) []byte {
	// 引数として可変長引数を受け取る場合、スライスに続けて...と記述すると、
	// スライス中の各項目を独立した引数として渡すことができる
	return append(word[:i+1], word[i:]...)
}

func removeVowel(word []byte, i int) []byte {
	return append(word[:i], word[i+1:]...)
}

func randBool() bool {
	// 0 か 1を返す乱数を0と比較
	return rand.Intn(2) == 0
}

func main() {
	// 乱数を作成
	rand.Seed(time.Now().UTC().UnixNano())

	// 標準入力を一行ずつ読み込み
	s := bufio.NewScanner(os.Stdin)

	// 入力の有無を判定してfor文を回す
	for s.Scan() {

		// sliceに入力された単語を格納
		word := []byte(s.Text())

		// 50%の確率でtrue/falseを生成
		if randBool() {
			// 母音のインデックスを初期化
			var vI int = -1
			for i, char := range word {
				switch char {
				// 母音かどうかを判定
				case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
					// 1/2の確率で、母音のインデックスを保存
					if randBool() {
						vI = i
					}
				}
			}
			if vI >= 0 {
				if randBool() {
					word = duplicateVowel(word, vI)
				} else {
					word = removeVowel(word, vI)
				}
			}
		}
		fmt.Println(string(word))
	}
}
