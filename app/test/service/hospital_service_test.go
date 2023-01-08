package service_test

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/service"
	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/test/mock/mock_client"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	t.Run("success case for GetLastDayOutpatientHistory()", func(t *testing.T) {
		time := time.Now().AddDate(0, 0, -1).Format("20060102")
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mock_json, _ := os.ReadFile("../../../test_data/20221230.json")
		mock_client := mock_client.NewMockCoronaClientInterface(mockCtrl)
		mock_client.EXPECT().GetMedicalSystem(&time).Return(mock_json, nil)

		service := service.HospitalService{Client: mock_client}
		outpatientHistories, err := service.GetLastDayOutpatientHistory()

		if err != nil {
			t.Fatal("Error failed to trigger with an invalid request")
		}
		assert.EqualValues(t, 24606, len(outpatientHistories), "The two words should be the same.")
	})

	t.Run("failure case for GetLastDayOutpatientHistory()", func(t *testing.T) {
		time := time.Now().AddDate(0, 0, -1).Format("20060102")
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		response_error := errors.New("errror")
		mock_client := mock_client.NewMockCoronaClientInterface(mockCtrl)
		mock_client.EXPECT().GetMedicalSystem(&time).Return(nil, response_error)

		service := service.HospitalService{Client: mock_client}
		outpatientHistories, err := service.GetLastDayOutpatientHistory()

		if err == nil {
			t.Fatal("Error failed to trigger with an invalid request")
		}
		assert.EqualValues(t, 0, len(outpatientHistories), "The two words should be the same.")
	})
}
