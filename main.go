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
	// 読み込むファイル
	portNumberText = "port-number.txt"
	// 新規作成して書き込むファイル
	portNumberCsv = "port-number.csv"
)

// ファイルパスを組み立てる構造体
type FilePath struct {
	dir      string
	filename string
	filepath string
}

// ファイルパスを組み立てる
func (fp *FilePath) CreateFilePath() error {
	// カレントディレクトリを取得
	currentDir, err := os.Getwd()
	if err != nil {
		return errors.New("CreateFilePath->os.Getwd()")
	}

	fp.filepath = fmt.Sprintf(
		"%s/%s/%s",
		currentDir, fp.dir, fp.filename)

	return nil
}

// テキスト処理を行う構造体
type TextFileProcessor struct {
	OriginalFilePath string
	NewFilePath      string
	AfterText        string
	Target           string
	Replace          string
}

// テキストを置換して置換後の文字列を返す
func (t *TextFileProcessor) ReplaceText() error {
	originalFileData, err := os.Open(t.OriginalFilePath)
	if err != nil {
		return fmt.Errorf("ファイルを開けませんでした")
	}
	defer originalFileData.Close()

	// スキャナを用意
	// トークンは行
	scanner := bufio.NewScanner(originalFileData)

	// ファイルの各行を格納する配列
	rowText := []string{}

	// EOF（ファイル末尾）になるまでスキャンを繰り返す
	for scanner.Scan() {
		// スキャンした内容を文字列で取得
		// 返り値に行末尾の改行文字は含まれない
		s := scanner.Text()

		// 置換処理
		s = strings.ReplaceAll(s, t.Target, t.Replace)

		// 文字列を格納する
		rowText = append(rowText, s)
	}

	// ファイル終端に正常に到達したか
	// 正常にファイル末尾まで読み込まれたらerrはnilを返す
	if err := scanner.Err(); err != nil {
		return fmt.Errorf(
			"error: TextFileProcessor()->ReplaceText()->scanner.Err()")
	}

	// 配列に格納している文字列を改行区切りでつなげる
	joinedText := strings.Join(rowText, "\n")
	fmt.Println(joinedText)

	t.AfterText = joinedText

	return nil
}

func (t *TextFileProcessor) Save() error {
	// 保存用ファイルを生成する
	f, err := os.Create(t.NewFilePath)
	if err != nil {
		return fmt.Errorf("新規ファイル作成に失敗しました")
	}

	// 作成したファイルに読み込んだ文字列を書き込む
	if _, err := f.WriteString(t.AfterText); err != nil {
		return fmt.Errorf("ファイルへの書き込みに失敗しました")
	}

	return nil
}

func main() {
	originalFilePath := FilePath{
		dir:      dataDir,
		filename: portNumberText,
	}
	if err := originalFilePath.CreateFilePath(); err != nil {
		fmt.Println(err)
		return
	}

	newFilePath := FilePath{
		dir:      dataDir,
		filename: portNumberCsv,
	}
	if err := newFilePath.CreateFilePath(); err != nil {
		fmt.Println(err)
		return
	}

	tfp := TextFileProcessor{
		OriginalFilePath: originalFilePath.filepath,
		NewFilePath:      newFilePath.filepath,
		Target:           "\t",
		Replace:          ",",
	}

  tfp.ReplaceText()
	tfp.Save()
}
