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
	url := "https://ytranking.net/ranking/"

	// getリクエスト
	res, _ := http.Get(url)
	defer res.Body.Close()
	// バッファ
	buf, _ := ioutil.ReadAll(res.Body)
	//文字検出
	det := chardet.NewTextDetector()
	detRes, _ := det.DetectBest(buf)
	//文字読み取り
	bytReader := bytes.NewReader(buf)
	reader, _ := charset.NewReaderLabel(detRes.Charset, bytReader)
	//パース
	docParse, _ := goquery.NewDocumentFromReader(reader)
	//トップ２０までのyoutuberのスクレイピング
	result := docParse.Find("div > article > div > div > section > ul > li").Text()
	fmt.Println(result)
}
