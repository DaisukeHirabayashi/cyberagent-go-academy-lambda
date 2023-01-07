package main

import (
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
)

func TestHandler(t *testing.T) {
	t.Run("Unable to get IP", func(t *testing.T) {
		time := time.Now().AddDate(0, 0, -1).Format("20060102")
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder("GET", "https://opendata.corona.go.jp/"+time,
			httpmock.NewStringResponder(200, "mocked"),
		)
		err := CreateOutpatientHistoires()
		if err == nil {
			t.Fatal("Error failed to trigger with an invalid request")
		}
	})
}
