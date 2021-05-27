package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/auctionee/core/internal"
	"github.com/auctionee/core/pkg/models"
)

type Response struct {
	Success bool `json:"success"`
}

type Request struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (r *Request) ToJSON() *bytes.Buffer {
	data, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return bytes.NewBuffer(data)
}

func Permissions(user models.UserInfo) error {
	reqData := Request{
		Login:    user.Login,
		Password: user.Password,
	}
	reqBytes := reqData.ToJSON()
	if reqBytes == nil {
		return fmt.Errorf("bad user data")
	}
	req, err := http.NewRequest(http.MethodGet, internal.AUTH_URL, reqBytes)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	rawBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	r := &Response{}
	err = json.Unmarshal(rawBody, r)

	if !r.Success || err != nil {
		return err
	}

	return nil
}
