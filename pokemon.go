package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://zukan.pokemon.co.jp/detail/018"
	fmt.Println("以下のurlからスクレイピングします。")
	fmt.Println("=>", url)

	res, _ := http.Get(url)
	defer res.Body.Close()

	buffer, _ := ioutil.ReadAll(res.Body)

	detect    := chardet.NewTextDetector()
	detRes, _ := detect.DetectBest(buffer)
	fmt.Println("文字コードを判定します。")
	fmt.Println("=>", detRes.Charset)

	byteReader := bytes.NewReader(buffer)
	reader, _  := charset.NewReaderLabel(detRes.Charset, byteReader)

	docParse, _ := goquery.NewDocumentFromReader(reader)

	result := docParse.Find("title").Text()
	fmt.Println("スクレイプしたテキストがでてきます。")
	fmt.Println("=>", result)
}