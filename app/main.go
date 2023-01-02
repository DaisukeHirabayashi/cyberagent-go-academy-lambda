package main

import (
	"net/http"

	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/service"
	"github.com/aws/aws-lambda-go/lambda"
)

func CreateOutpatientHistoires() error {
	outpatientHistories, err := service.GetLastDayOutpatientHistory()
	if err != nil {
		go http.Get("http://127.0.0.1:3000/notification?message=" + "error GetLastDayOutpatientHistory:" + err.Error())
		// err_json, _ := json.Marshal(payload{Text: "error GetLastDayOutpatientHistory:" + err.Error()})
		// http.PostForm("https://hooks.slack.com/services/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", url.Values{"payload": {string(err_json)}})
	}

	err = service.CreateOutpatientHistoires(outpatientHistories)

	if err != nil {
		go http.Get("http://127.0.0.1:3000/notification?message=" + "error GetLastDayOutpatientHistory:" + err.Error())
		// err_json, _ := json.Marshal(payload{Text: "error CreateOutpatientHistoires:" + err.Error()})
		// http.PostForm("https://hooks.slack.com/services/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", url.Values{"payload": {string(err_json)}})
	}
	return nil
}

func main() {
	lambda.Start(CreateOutpatientHistoires)
}
