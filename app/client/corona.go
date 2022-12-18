package client

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/dao"
)

var host = "https://opendata.corona.go.jp/"

func GetMedicalSystem() {
	log.Println("Get")

	url := host + "api/covid19DailySurvey"

	client := new(http.Client)
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	q := req.URL.Query()
	q.Add("prefName", "石川県")

	resp, err := client.Do(req)
	log.Print(resp)

	if err != nil {
		log.Println("Error Request:", err)
		return
	}
	// resp.Bodyはクローズすること。クローズしないとTCPコネクションを開きっぱなしになる。
	defer resp.Body.Close()

	// 200 OK 以外の場合はエラーメッセージを表示して終了
	if resp.StatusCode != 200 {
		log.Println("Error Response:", resp.Status)
		return
	}

	// Response Body を読み取り
	body, _ := io.ReadAll(resp.Body)

	// JSONを構造体にエンコード
	var Hospitals []dao.Hospital
	json.Unmarshal(body, &Hospitals)
}
