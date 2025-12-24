package cryptomus

import (
	"context"
)

type Discount struct {
	// 货币代码（例如 BTC、ETH、USDT 等）
	Currency string `json:"currency"`
	// 区块链网络（例如 ARBITRUM、AVALANCHE、BCH 等）
	Network string `json:"network"`
	// 折扣百分比（例如 5、10、-15 等）
	Discount int `json:"discount"`
}

type ListOfDiscount []Discount

type ListOfDiscountResponse struct {
	*HTTPResponse
	Result ListOfDiscount `json:"result"`
}

// ListOfDiscount 返回折扣列表。
//
// 详情：https://doc.cryptomus.com/business/discount/list
//
// 示例：
//
//	result, err := sdk.ListOfDiscount()
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) ListOfDiscount() (*ListOfDiscountResponse, error) {
	return sdk.ListOfDiscountWithContext(context.Background())
}

// ListOfDiscountWithContext 返回折扣列表。
//
// 详情：https://doc.cryptomus.com/business/discount/list
//
// 示例：
//
//	result, err := sdk.ListOfDiscountWithContext(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) ListOfDiscountWithContext(ctx context.Context) (*ListOfDiscountResponse, error) {

	var result ListOfDiscountResponse

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, nil)).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(ListOfDiscountsEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}

type SetDiscountToPaymentMethodRequest struct {
	// 货币代码（例如 BTC、ETH、USDT 等）
	Currency string `json:"currency"`
	// 区块链网络（例如 ARBITRUM、AVALANCHE、BCH 等）
	Network string `json:"network"`
	// 折扣百分比（例如 5、10、-15 等）
	Discount int `json:"discount_percent"`
}

type SetDiscountToPaymentMethodResponse struct {
	*HTTPResponse
	Result Discount `json:"result,omitempty"`
}

// SetDiscountToPaymentMethod 为支付方式设置折扣。
//
// 详情：https://doc.cryptomus.com/business/discount/set
//
// 折扣百分比：
//   - 正数 (>0)：为使用该加密货币支付的买家提供折扣。
//     适用于促销活动或鼓励使用特定加密货币支付。
//   - 负数 (<0)：为使用该加密货币支付增加一个百分比（填充）。
//     这可以帮助覆盖加密货币/法币转换成本或调整当地汇率差异。
//
// 示例：
//
//	err := sdk.SetDiscountToPaymentMethod(payload)
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) SetDiscountToPaymentMethod(payload SetDiscountToPaymentMethodRequest) (*SetDiscountToPaymentMethodResponse, error) {
	return sdk.SetDiscountToPaymentMethodWithContext(context.Background(), payload)
}

// SetDiscountToPaymentMethodWithContext 为支付方式设置折扣。
//
// 详情：https://doc.cryptomus.com/business/discount/set
//
// 折扣百分比：
//   - 正数 (>0)：为使用该加密货币支付的买家提供折扣。
//     适用于促销活动或鼓励使用特定加密货币支付。
//   - 负数 (<0)：为使用该加密货币支付增加一个百分比（填充）。
//     这可以帮助覆盖加密货币/法币转换成本或调整当地汇率差异。
//
// 示例：
//
//	err := sdk.SetDiscountToPaymentMethodWithContext(ctx, payload)
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) SetDiscountToPaymentMethodWithContext(ctx context.Context, payload SetDiscountToPaymentMethodRequest) (*SetDiscountToPaymentMethodResponse, error) {

	var result SetDiscountToPaymentMethodResponse

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, payload)).
		SetBody(payload).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(SetDiscountToPaymentMethodEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}
