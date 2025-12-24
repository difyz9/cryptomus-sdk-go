package cryptomus

import "context"

type BalanceResponse struct {
	*HTTPResponse
	Result []BalanceResult `json:"result,omitempty"`
}

// Merchant 代表商户账户。
type Merchant struct {
	// 钱包 UUID
	UUID string `json:"uuid"`
	// 商户钱包余额
	Balance string `json:"balance"`
	// 钱包货币代码（例如 BTC、ETH、USDT 等）
	CurrencyCode string `json:"currency_code"`
}

// User 代表个人账户。
type User struct {
	// 钱包 UUID
	UUID string `json:"uuid"`
	// 个人钱包余额
	Balance string `json:"balance"`
	// 钱包货币代码（例如 BTC、ETH、USDT 等）
	CurrencyCode string `json:"currency_code"`
}

// Balance 代表余额列表。
//
// 余额列表包括商户（merchant）和个人（user）钱包。
type Balance struct {
	Merchant []Merchant `json:"merchant"`
	User     []User     `json:"user"`
}

type BalanceResult struct {
	Balance Balance `json:"balance"`
}

// Balance 返回余额列表。
//
// 详情：https://doc.cryptomus.com/business/balance
//
// 示例：
//
//	result, err := sdk.Balance()
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) Balance() (*BalanceResponse, error) {
	return sdk.BalanceWithContext(context.Background())
}

// BalanceWithContext 返回余额列表。
//
// 详情：https://doc.cryptomus.com/business/balance
//
// 示例：
//
//	result, err := sdk.BalanceWithContext(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) BalanceWithContext(ctx context.Context) (*BalanceResponse, error) {
	var result BalanceResponse

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, nil)).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(BalanceEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}
