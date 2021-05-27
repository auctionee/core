package balance

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/auctionee/core/internal"
	"github.com/auctionee/core/pkg/models"
)

type RequestBalance struct {
	Login string `json:"Login"`
}
type RequestModify struct {
	Login  string `json:"Login"`
	Amount int    `json:"Amount"`
	Key    string `json:"Key"`
}
type Response struct {
	Balance int  `json:"Balance"`
	Status  bool `json:"Status"`
}

func GetBalance(user models.UserInfo) (balance int, err error) {
	requestData := RequestBalance{Login: user.Login}
	b, err := json.Marshal(requestData)
	if err != nil {
		return 0, err
	}
	reqBytes := bytes.NewBuffer(b)
	req, err := http.NewRequest(http.MethodGet, internal.BALANCE_URL, reqBytes)
	if err != nil {
		return 0, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}
	responseData := &Response{}
	if err = json.Unmarshal(resBytes, responseData); err != nil {
		return 0, err
	}

	return responseData.Balance, nil
}

func Charge(login string, amount int) error {
	requestData := RequestModify{
		Login:  login,
		Amount: amount,
		Key:    internal.SuperSecretKey}

	b, err := json.Marshal(requestData)
	if err != nil {
		return err
	}
	reqBytes := bytes.NewBuffer(b)
	req, err := http.NewRequest(http.MethodGet, internal.BALANCE_URL_CHARGE, reqBytes)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	responseData := &Response{}
	err = json.Unmarshal(resBytes, responseData)
	if err != nil || !responseData.Status {
		return err
	}

	return nil
}
func Refund(login string, amount int) error {
	requestData := RequestModify{
		Login:  login,
		Amount: amount,
		Key:    internal.SuperSecretKey}

	b, err := json.Marshal(requestData)
	if err != nil {
		return err
	}
	reqBytes := bytes.NewBuffer(b)
	req, err := http.NewRequest(http.MethodGet, internal.BALANCE_URL_REFUND, reqBytes)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	responseData := &Response{}
	err = json.Unmarshal(resBytes, responseData)
	if err != nil || !responseData.Status {
		return err
	}

	return nil
}
