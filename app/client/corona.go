package client

import (
	"io"
	"log"
	"net/http"
)

var host = "https://opendata.corona.go.jp/"

func GetMedicalSystem(time *string) ([]byte, error) {
	log.Println("Get")

	url := host + "api/covid19DailySurvey/" + *time
	log.Print(url)

	client := new(http.Client)
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	resp, err := client.Do(req)
	log.Print(resp)

	if err != nil {
		log.Println("Error Request:", err)
		return nil, err
	}

	// resp.Bodyはクローズすること。クローズしないとTCPコネクションを開きっぱなしになる。
	defer resp.Body.Close()

	// 200 OK 以外の場合はエラーメッセージを表示して終了
	if resp.StatusCode != 200 {
		log.Println("Error Response:", resp.Status)
		return nil, err
	}

	// Response Body を読み取り
	body, _ := io.ReadAll(resp.Body)
	return body, nil
}
