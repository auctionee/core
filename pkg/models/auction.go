package models

import (
	"encoding/json"
	"io"
)

type UserInfo struct {
	Login    string `json:"Login"`
	Password string `json:"Password"`
}
type Unit struct {
	Name  string `json:"Name"`
	Price int    `json:"Price"`
}
type BetInfo struct {
	Amount int    `json:"Amount"`
	AUID   string `json:"AUID"`
}
type StartInfo struct {
	Good            Unit   `json:"Unit"`
	DurationMinutes string `json:"Duration"`
}

type Request struct {
	UserInfo  UserInfo  `json:"UserInfo"`
	Start     bool      `json:"Start"`
	Bet       bool      `json:"Bet"`
	StartInfo StartInfo `json:"StartInfo"`
	BetInfo   BetInfo   `json:"BetInfo"`
}

func NewRequest() Request {
	return Request{}
}
func (s *Request) Unmarshall(body io.ReadCloser) error {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(bodyBytes, s); err != nil {
		return err
	}
	return nil
}
