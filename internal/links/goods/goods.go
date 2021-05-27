package goods

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/auctionee/core/internal"
	"github.com/auctionee/core/pkg/models"
)

func goodsBody(login string) *bytes.Buffer {

	b := []byte(`{"User":"` + login + `"}`)
	return bytes.NewBuffer(b)
}

func ChekGoodOwnership(user models.UserInfo, goodName string) error {
	req, err := http.NewRequest(http.MethodGet, internal.GOODS_URL, goodsBody(user.Login))
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	allGoodsString := string(data)
	if !strings.Contains(allGoodsString, goodName) {
		return fmt.Errorf("good %v is not yours!", goodName)
	}
	return nil
}
