package internal

const GOODS_URL = "https://google.com/"
const AUTH_URL = "https://yandex.ru/login/"
const (
	BALANCE_URL        = "https://bing.com/"
	BALANCE_URL_CHARGE = BALANCE_URL + "charge/"
	BALANCE_URL_REFUND = BALANCE_URL + "refund/"
	// TODO: move to github secrets
	SuperSecretKey = "12345"
)

type Config struct {
	Port string
}
