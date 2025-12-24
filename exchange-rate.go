package cryptomus

import (
	"context"
	"fmt"
)

type ExchangeRate struct {
	// 源货币代码（例如 BTC、ETH、USDT 等）
	From string `json:"from,omitempty"`
	// 目标货币代码（例如 BTC、ETH、USDT 等）
	To string `json:"to,omitempty"`
	// 汇率（例如 0.0001、0.0002、0.0003 等）
	Course string `json:"course,omitempty"`
}

type ExchangeRateList []ExchangeRate

type ExchangeRateResponse struct {
	*HTTPResponse
	Result ExchangeRateList `json:"result,omitempty"`
}

// ExchangeRateList 返回汇率列表。
//
// 详情：https://doc.cryptomus.com/business/exchange-rates/list
//
// 示例：
//
//	result, err := sdk.ExchangeRateList("ETH")
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) ExchangeRateList(currency string) (*ExchangeRateResponse, error) {
	return sdk.ExchangeRateListWithContext(context.Background(), currency)
}

// ExchangeRateListWithContext 返回汇率列表。
//
// 详情：https://doc.cryptomus.com/business/exchange-rates/list
//
// 示例：
//
//	result, err := sdk.ExchangeRateListWithContext(ctx, "ETH")
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) ExchangeRateListWithContext(ctx context.Context, currency string) (*ExchangeRateResponse, error) {
	url := fmt.Sprintf(ExchangeRateListEndpoint.URL(), currency)

	var result ExchangeRateResponse

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Get(url); err != nil {
		return nil, err
	}

	return &result, nil
}
