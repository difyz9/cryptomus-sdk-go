package cryptomus

import (
	"context"
	"time"
)

type RecurringPaymentPeriod string

const (
	RecurringPaymentPeriodWeekly     RecurringPaymentPeriod = "weekly"
	RecurringPaymentPeriodMonthly    RecurringPaymentPeriod = "monthly"
	RecurringPaymentPeriodThreeMonth RecurringPaymentPeriod = "three_month"
)

type CreateRecurringPaymentRequest struct {
	Amount         string                 `json:"amount"`
	Currency       string                 `json:"currency"`
	Name           string                 `json:"name"`
	Period         RecurringPaymentPeriod `json:"period"`
	ToCurrency     string                 `json:"to_currency,omitempty"`
	OrderID        string                 `json:"order_id,omitempty"`
	URLCallback    string                 `json:"url_callback,omitempty"`
	DiscountDays   int                    `json:"discount_days,omitempty"`
	DiscountAmount string                 `json:"discount_amount,omitempty"`
	AdditionalData string                 `json:"additional_data,omitempty"`
}

type CreateRecurringPaymentResponse struct {
	*HTTPResponse
	Result RecurringPaymentData `json:"result,omitempty"`
}

type RecurringPaymentData struct {
	UUID           string     `json:"uuid"`
	Name           string     `json:"name"`
	OrderID        string     `json:"order_id,omitempty"`
	Amount         string     `json:"amount"`
	Currency       string     `json:"currency"`
	PayerCurrency  string     `json:"payer_currency,omitempty"`
	PayerAmountUSD string     `json:"payer_amount_usd"`
	PayerAmount    string     `json:"payer_amount,omitempty"`
	URLCallback    string     `json:"url_callback,omitempty"`
	Period         string     `json:"period"`
	Status         string     `json:"status"`
	URL            string     `json:"url"`
	LastPayOff     *time.Time `json:"last_pay_off,omitempty"`
}

// CreateRecurringPayment 创建循环支付。
//
// 详情：https://doc.cryptomus.com/business/recurring/creating
//
// 示例：
//
//	result, err := sdk.CreateRecurringPayment(CreateRecurringPaymentRequest{
//		Amount:   "0.0001",
//		Currency: "BTC",
//		Name:     "Test recurring payment",
//		Period:   RecurringPaymentPeriodWeekly,
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) CreateRecurringPayment(request CreateRecurringPaymentRequest) (*CreateRecurringPaymentResponse, error) {
	return sdk.CreateRecurringPaymentWithContext(context.Background(), request)
}

// CreateRecurringPaymentWithContext 创建循环支付。
//
// 详情：https://doc.cryptomus.com/business/recurring/creating
//
// 加密货币循环支付是使用数字资产自动化定期交易的一种方式。它们可用于订阅服务、捐赠、会员资格和其他定期支付。
//
// 要使用循环支付，您需要创建一个支付，指定金额、货币和支付频率，然后将其分享给您的付款人。付款人将被重定向到 Cryptomus 网站，在那里他需要登录以确认支付计划并进行首次支付。之后，支付将按照计划自动进行。
//
// 示例：
//
//	result, err := sdk.CreateRecurringPaymentWithContext(ctx, CreateRecurringPaymentRequest{
//		Amount:   "0.0001",
//		Currency: "BTC",
//		Name:     "Test recurring payment",
//		Period:   RecurringPaymentPeriodWeekly,
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) CreateRecurringPaymentWithContext(ctx context.Context, payload CreateRecurringPaymentRequest) (*CreateRecurringPaymentResponse, error) {
	var result CreateRecurringPaymentResponse

	reqByte, err := ToJSON(payload)
	if err != nil {
		return nil, err
	}

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, payload)).
		SetBody(reqByte).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(CreateRecurringPaymentEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}

type RecurringPaymentInformationRequest struct {
	UUID    string `json:"uuid,omitempty"`
	OrderID string `json:"order_id,omitempty"`
}

type RecurringPaymentInformationResponse struct {
	*HTTPResponse
	Result RecurringPaymentData `json:"result,omitempty"`
}

// RecurringPaymentInformation returns information about a recurring payment.
//
// Details: https://doc.cryptomus.com/business/recurring/info
//
// Example:
//
//	result, err := sdk.RecurringPaymentInformation(RecurringPaymentInformationRequest{
//		UUID: "4b1b3b7b-7b1b-4b1b-7b1b-4b1b3b7b1b4b",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) RecurringPaymentInformation(request RecurringPaymentInformationRequest) (*RecurringPaymentInformationResponse, error) {
	return sdk.RecurringPaymentInformationWithContext(context.Background(), request)
}

// RecurringPaymentInformationWithContext returns information about a recurring payment.
//
// Details: https://doc.cryptomus.com/business/recurring/info
//
// To get the recurring payment status you need to pass one of the required parameters, if you pass both, the account will be identified by order_id
//
// Example:
//
//	result, err := sdk.RecurringPaymentInformationWithContext(ctx, RecurringPaymentInformationRequest{
//		UUID: "4b1b3b7b-7b1b-4b1b-7b1b-4b1b3b7b1b4b",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) RecurringPaymentInformationWithContext(ctx context.Context, payload RecurringPaymentInformationRequest) (*RecurringPaymentInformationResponse, error) {
	var result RecurringPaymentInformationResponse

	reqByte, err := ToJSON(payload)
	if err != nil {
		return nil, err
	}

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, payload)).
		SetBody(reqByte).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(RecurringPaymentInformationEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}

type ListRecurringPaymentsRequest struct {
	Cursor string `json:"cursor,omitempty"`
}

type ListRecurringPaymentsResponse struct {
	*HTTPResponse
	Result ListRecurringPaymentsData `json:"result,omitempty"`
}

type ListRecurringPaymentsData struct {
	Items    []RecurringPaymentData `json:"items"`
	Paginate *Pagination            `json:"paginate"`
}

// ListRecurringPayments returns a list of recurring payments.
//
// Details: https://doc.cryptomus.com/business/recurring/list
//
// Example:
//
//	result, err := sdk.ListRecurringPayments(ListRecurringPaymentsRequest{
//		Cursor: "eyJpZCI6MjEyLCJfcG9pbnRzVzVG9OZhXh0SXRlbXMiOnRydWV9",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) ListRecurringPayments(request ListRecurringPaymentsRequest) (*ListRecurringPaymentsResponse, error) {
	return sdk.ListRecurringPaymentsWithContext(context.Background(), request)
}

// ListRecurringPaymentsWithContext returns a list of recurring payments.
//
// Details: https://doc.cryptomus.com/business/recurring/list
//
// Example:
//
//	result, err := sdk.ListRecurringPaymentsWithContext(ctx, ListRecurringPaymentsRequest{
//		Cursor: "eyJpZCI6MjEyLCJfcG9pbnRzVzFG9OZhXh0SXRlbXMiOnRydWV9",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) ListRecurringPaymentsWithContext(ctx context.Context, payload ListRecurringPaymentsRequest) (*ListRecurringPaymentsResponse, error) {
	var result ListRecurringPaymentsResponse

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, nil)).
		SetQueryParam("cursor", payload.Cursor).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(ListRecurringPaymentsEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}

type CancelRecurringPaymentRequest struct {
	UUID    string `json:"uuid,omitempty"`
	OrderID string `json:"order_id,omitempty"`
}

type CancelRecurringPaymentResponse struct {
	*HTTPResponse
	Result RecurringPaymentData `json:"result,omitempty"`
}

// CancelRecurringPayment cancels a recurring payment.
//
// Details: https://doc.cryptomus.com/business/recurring/cancel
//
// To cancel the recurring payment you need to pass one of the required parameters, if you pass both, the account will be identified by order_id
//
// Example:
//
//	result, err := sdk.CancelRecurringPayment(CancelRecurringPaymentRequest{
//		UUID: "4b1b3b7b-7b1b-4b1b-7b1b-4b1b3b7b1b4b",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) CancelRecurringPayment(payload CancelRecurringPaymentRequest) (*CancelRecurringPaymentResponse, error) {
	return sdk.CancelRecurringPaymentWithContext(context.Background(), payload)
}

// CancelRecurringPaymentWithContext cancels a recurring payment.
//
// Details: https://doc.cryptomus.com/business/recurring/cancel
//
// To cancel the recurring payment you need to pass one of the required parameters, if you pass both, the account will be identified by order_id
//
// Example:
//
//	result, err := sdk.CancelRecurringPaymentWithContext(ctx, CancelRecurringPaymentRequest{
//		UUID: "4b1b3b7b-7b1b-4b1b-7b1b-4b1b3b7b1b4b",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) CancelRecurringPaymentWithContext(ctx context.Context, payload CancelRecurringPaymentRequest) (*CancelRecurringPaymentResponse, error) {
	var result CancelRecurringPaymentResponse

	reqByte, err := ToJSON(payload)
	if err != nil {
		return nil, err
	}

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, payload)).
		SetBody(reqByte).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(CancelRecurringPaymentEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}
