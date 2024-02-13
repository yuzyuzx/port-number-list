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
	data, err := os.Open(dataFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer data.Close()

	// スキャナを用意
	// トークンは行
	scanner := bufio.NewScanner(data)

  // ファイルの各行を格納する配列
  rowText := []string{}

	// EOF（ファイル末尾）になるまでスキャンを繰り返す
	for scanner.Scan() {
		// スキャンした内容を文字列で取得
		// 返り値に行末尾の改行文字は含まれない
		s := scanner.Text()

		// タブをカンマに置換
		s = strings.ReplaceAll(s, "\t", ",")

    // 文字列を格納する
    rowText = append(rowText, s)

	}

	// ファイル終端に正常に到達したか
	// 正常にファイル末尾まで読み込まれたらerrはnilを返す
	if err := scanner.Err(); err != nil {
		fmt.Println("scanner.Err()")
    return
	}

  // 配列に格納している文字列を改行区切りでつなげる
  joinedText := strings.Join(rowText, "\n")
  fmt.Println(joinedText)

  // 保存用ファイルを新規作成
  newFile, err := os.Create("data/result.csv")
  if err != nil {
    fmt.Println("新規ファイル作成に失敗しました")
    return
  }

  // 作成したファイルに読み込んだ文字列を書き込む
  if _, err := newFile.WriteString(joinedText); err != nil {
    fmt.Println("ファイルへの書き込みに失敗しました")
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
