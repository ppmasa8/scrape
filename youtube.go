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
//出力結果

//1
//
//
//せんももあいしーCh Sen, Momo, Ai & Shii
//people1030万人
//play_circle_filled71億5437万8052回
//videocam3052本
//チャンネルの詳細
//2
//
//
//HikakinTV
//people992万人
//play_circle_filled87億6817万218回
//videocam2962本
//チャンネルの詳細
//3
//
//
//はじめしゃちょー（hajime）
//people952万人
//play_circle_filled82億6034万3370回
//videocam2379本
//チャンネルの詳細
//4
//
//
//Junya.じゅんや
//people945万人
//play_circle_filled42億9091万6652回
//videocam1645本
//チャンネルの詳細
//5
//
//
//タキロン Takilong Kids' Toys
//people737万人
//play_circle_filled28億29万1889回
//videocam1859本
//チャンネルの詳細
//6
//
//
//Travel Thirsty
//people694万人
//play_circle_filled29億3814万1795回
//videocam683本
//チャンネルの詳細
//7
//
//
//Fischer's-フィッシャーズ-
//people691万人
//play_circle_filled124億9169万4674回
//videocam2452本
//チャンネルの詳細
//8
//
//
//Nino's Home
//people624万人
//play_circle_filled5億5373万3545回
//videocam87本
//チャンネルの詳細
//9
//
//
//東海オンエア
//people607万人
//play_circle_filled89億6050万9210回
//videocam2017本
//チャンネルの詳細
//10
//
//
//米津玄師
//people602万人
//play_circle_filled37億8253万510回
//videocam68本
//チャンネルの詳細
//11
//
//
//SUSHI RAMEN【Riku】
//people586万人
//play_circle_filled14億2928万2068回
//videocam464本
//チャンネルの詳細
//12
//
//
//avex
//people568万人
//play_circle_filled92億2127万4600回
//videocam1万5038本
//チャンネルの詳細
//13
//
//
//Yuka Kinoshita木下ゆうか
//people548万人
//play_circle_filled21億3253万430回
//videocam2185本
//チャンネルの詳細
//14
//
//
//JunsKitchen
//people525万人
//play_circle_filled3億6244万4920回
//videocam35本
//チャンネルの詳細
//15
//
//
//HikakinGames
//people524万人
//play_circle_filled62億3133万3857回
//videocam1796本
//チャンネルの詳細
//16
//
//
//THE FIRST TAKE
//people492万人
//play_circle_filled13億6273万6860回
//videocam247本
//チャンネルの詳細
//17
//
//
//mania japansong
//people483万人
//play_circle_filled23億8257万1880回
//videocam2196本
//チャンネルの詳細
//18
//
//
//ヒカル（Hikaru）
//people453万人
//play_circle_filled37億283万9529回
//videocam1829本
//チャンネルの詳細
//18
//
//
//きまぐれクックKimagure Cook
//people453万人
//play_circle_filled17億9065万4097回
//videocam641本
//チャンネルの詳細
//20
//
//
//はねまりチャンネル・HANEMARI Channel
//people451万人
//play_circle_filled26億8868万301回
//videocam1472本
//チャンネルの詳細
