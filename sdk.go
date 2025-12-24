package cryptomus

import (
	"github.com/imroc/req/v3"
)

type Cryptomus struct {
	HttpClient   *req.Client
	Merchant     string
	PaymentToken string
	PayoutToken  string
}

type Option func(*Cryptomus)

// New 创建一个新的 Cryptomus SDK 实例。
//
// 在使用 SDK 之前，请确保设置 Merchant 和 PaymentToken。参考示例：https://doc.cryptomus.com/business/general/getting-api-keys
//
// 同时查看请求格式：https://doc.cryptomus.com/business/general/request-format
//
// 您可以为每个 SDK 实例设置 PaymentToken 和 PayoutToken。
func New(options ...Option) *Cryptomus {
	cryptomus := DefaultCryptomus()

	for _, option := range options {
		option(cryptomus)
	}

	return cryptomus
}

func DefaultCryptomus() *Cryptomus {
	return &Cryptomus{
		HttpClient: DefaultHTTPClient(),
	}
}

func WithHttpClient(client *req.Client) Option {
	return func(c *Cryptomus) {
		c.HttpClient = client
	}
}

func WithMerchant(merchant string) Option {
	return func(c *Cryptomus) {
		c.Merchant = merchant
	}
}

func WithPaymentToken(token string) Option {
	return func(c *Cryptomus) {
		c.PaymentToken = token
	}
}

func WithPayoutToken(token string) Option {
	return func(c *Cryptomus) {
		c.PayoutToken = token
	}
}
