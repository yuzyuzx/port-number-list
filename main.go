package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	// データ保存ディレクトリ
	dataDir = "data"
	// データファイル
	portNumberText = "port-number.txt"
)

func main() {
	dataFile, err := createFilePath()
	if err != nil {
		fmt.Println(err)
		return
	}

	// ファイルを用意
	data, _ := os.Open(dataFile)
	defer data.Close()

	// スキャナを用意
	// トークンは行
	scanner := bufio.NewScanner(data)

	// EOF（ファイル末尾）になるまでスキャンを繰り返す
	for scanner.Scan() {
		// スキャンした内容を文字列で取得
		// 返り値に行末尾の改行文字は含まれない
		s := scanner.Text()

		// タブをカンマに置換
		s = strings.ReplaceAll(s, "\t", ",")

		fmt.Println(s)
	}

	// ファイル終端に正常に到達したか
	// 正常にファイル末尾まで読み込まれたらerrはnilを返す
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "scanner.Err()")
	}

}

// ファイルパスを文字列で組み立てる
func createFilePath() (string, error) {
	// カレントディレクトリを取得
	currentDir, err := os.Getwd()
	if err != nil {
		return "", errors.New("createFilePath()")
	}

	return fmt.Sprintf("%s/%s/%s", currentDir, dataDir, portNumberText), nil
}
