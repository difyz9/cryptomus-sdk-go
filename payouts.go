package cryptomus

import (
	"context"
	"time"
)

// PayoutStatus 提现状态在某些方法的响应主体中返回，指示提现当前所处的阶段。
//
// 详情：https://doc.cryptomus.com/business/payouts/payout-statuses
type PayoutStatus string

const (
	PayoutStatusProcess    PayoutStatus = "process"     // 提现处理中
	PayoutStatusCheck      PayoutStatus = "check"       // 提现正在验证
	PayoutStatusPaid       PayoutStatus = "paid"        // 提现成功
	PayoutStatusFail       PayoutStatus = "fail"        // 提现失败
	PayoutStatusCancel     PayoutStatus = "cancel"      // 提现已取消
	PayoutStatusSystemFail PayoutStatus = "system_fail" // 系统错误
)

type CreatePayoutRequest struct {
	Amount       string `json:"amount"`
	Currency     string `json:"currency" validate:"required"`
	OrderID      string `json:"order_id"`
	Address      string `json:"address"`
	IsSubtract   bool   `json:"is_subtract"`
	Network      string `json:"network"`
	URLCallback  string `json:"url_callback"`
	ToCurrency   string `json:"to_currency"`
	CourseSource string `json:"course_source"`
	FromCurrency string `json:"from_currency"`
	Priority     string `json:"priority"`
	Memo         string `json:"memo"`
}

type PayoutData struct {
	UUID          string       `json:"uuid"`
	Amount        string       `json:"amount"`
	Currency      string       `json:"currency"`
	Network       string       `json:"network"`
	Address       string       `json:"address"`
	TxID          string       `json:"txid"`
	Status        PayoutStatus `json:"status"`
	IsFinal       bool         `json:"is_final"`
	Balance       int64        `json:"balance"`
	PayerCurrency string       `json:"payer_currency"`
	PayerAmount   int64        `json:"payer_amount"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
}

type CreatePayoutResponse struct {
	*HTTPResponse
	Result *PayoutData `json:"result"`
}

// CreatePayout 创建提现。
//
// 详情：https://doc.cryptomus.com/business/payout/creating
//
// 通过 API 提现仅从您的商户钱包余额中扣除。
//
// 提现可以通过以下不同方式进行：
//
// - 您可以选择以特定加密货币接收提现，提现将自动以该特定加密货币处理。为此，请确保您在该特定货币中有足够的余额来支付所有相关费用。
//
// - 或者，您可以选择以法定货币指定提现金额。在这种情况下，金额将自动从您的可用余额转换为特定加密货币。例如，如果您请求以 LTC 提现 20 美元，等值金额将从您的 LTC 余额中扣除。重要的是在相应的加密货币中有足够的资金来支付所有适用费用。
//
// - 另一种可能性是以法定货币指定提现金额，该金额将使用您的 USDT 余额自动转换为特定加密货币。当您启用自动转换时，此选项特别有用，因为来自发票的资金会自动转换为 USDT。例如，如果您想以 LTC 提现 20 美元，但只有 USDT 余额，请确保您有足够的 USDT 资金来支付所有费用。
//
// - 此外，您可以选择以您偏好的任何加密货币指定提现金额。然后，提现将自动以该特定加密货币处理，使用您可用的 USDT 余额。拥有足够的 USDT 余额来支付所有相关费用至关重要。
//
// 示例：
//
//	result, err := sdk.CreatePayout(&CreatePayoutRequest{
//		Amount:   "5",
//		Currency: "USDT",
//		Network:  "TRON",
//		Address:  "TJ5Zrj8z6bJ7bk9Kf8fz1yQFbJ7b7bJ7b7b",
//		OrderID: "1",
//		URLCallback: "https://example.com/callback",
//		IsSubtract: true,
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) CreatePayout(payload *CreatePayoutRequest) (*CreatePayoutResponse, error) {
	return sdk.CreatePayoutWithContext(context.Background(), payload)
}

// CreatePayoutWithContext 创建提现。
//
// 详情：https://doc.cryptomus.com/business/payout/creating
//
// 通过 API 提现仅从您的商户钱包余额中扣除。
//
// 提现可以通过以下不同方式进行：
//
// - 您可以选择以特定加密货币接收提现，提现将自动以该特定加密货币处理。为此，请确保您在该特定货币中有足够的余额来支付所有相关费用。
//
// - 或者，您可以选择以法定货币指定提现金额。在这种情况下，金额将自动从您的可用余额转换为特定加密货币。例如，如果您请求以 LTC 提现 20 美元，等值金额将从您的 LTC 余额中扣除。重要的是在相应的加密货币中有足够的资金来支付所有适用费用。
//
// - 另一种可能性是以法定货币指定提现金额，该金额将使用您的 USDT 余额自动转换为特定加密货币。当您启用自动转换时，此选项特别有用，因为来自发票的资金会自动转换为 USDT。例如，如果您想以 LTC 提现 20 美元，但只有 USDT 余额，请确保您有足够的 USDT 资金来支付所有费用。
//
// - 此外，您可以选择以您偏好的任何加密货币指定提现金额。然后，提现将自动以该特定加密货币处理，使用您可用的 USDT 余额。拥有足够的 USDT 余额来支付所有相关费用至关重要。
//
// 示例：
//
//	result, err := sdk.CreatePayout(ctx, &CreatePayoutRequest{
//		Amount:   "5",
//		Currency: "USDT",
//		Network:  "TRON",
//		Address:  "TJ5Zrj8z6bJ7bk9Kf8fz1yQFbJ7b7bJ7b7b",
//		OrderID: "1",
//		URLCallback: "https://example.com/callback",
//		IsSubtract: true,
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) CreatePayoutWithContext(ctx context.Context, payload *CreatePayoutRequest) (*CreatePayoutResponse, error) {
	var result CreatePayoutResponse

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PayoutToken, payload)).
		SetBody(payload).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(CreatePayoutEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}

type PayoutInformationRequest struct {
	UUID    string `json:"uuid,omitempty"`
	OrderID string `json:"order_id,omitempty"`
}

type PayoutInformationResponse struct {
	*HTTPResponse
	Result *PayoutData `json:"result"`
}

// PayoutInformation returns information about a payout.
//
// Details: https://doc.cryptomus.com/business/payouts/payout-information
//
// To get the payout information you need to pass one of the parameters, if you pass both, the payout will be identified by order_id
//
// Example:
//
//	result, err := sdk.PayoutInformation(&PayoutInformationRequest{
//		UUID: "4b1b3b7b-7b1b-4b1b-7b1b-4b1b3b7b1b4b",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) PayoutInformation(payload *PayoutInformationRequest) (*PayoutInformationResponse, error) {
	return sdk.PayoutInformationWithContext(context.Background(), payload)
}

// PayoutInformationWithContext returns information about a payout.
//
// Details: https://doc.cryptomus.com/business/payouts/payout-information
//
// To get the payout information you need to pass one of the parameters, if you pass both, the payout will be identified by order_id
//
// Example:
//
//	result, err := sdk.PayoutInformationWithContext(ctx, &PayoutInformationRequest{
//		UUID: "4b1b3b7b-7b1b-4b1b-7b1b-4b1b3b7b1b4b",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) PayoutInformationWithContext(ctx context.Context, payload *PayoutInformationRequest) (*PayoutInformationResponse, error) {
	var result PayoutInformationResponse

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PayoutToken, payload)).
		SetBody(payload).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(PayoutInformationEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}

type PayoutHistoryRequest struct {
	DateFrom string `json:"date_from,omitempty"`
	DateTo   string `json:"date_to,omitempty"`
}

type PayoutHistoryData struct {
	MerchantUUID string        `json:"merchant_uuid"`
	Items        []*PayoutData `json:"items"`
	Paginate     *Pagination   `json:"paginate"`
}

type PayoutHistoryResponse struct {
	*HTTPResponse
	Result *PayoutHistoryData `json:"result"`
}

// PayoutHistory returns a list of payouts.
//
// Details: https://doc.cryptomus.com/business/payouts/payout-history
//
// Example:
//
//	result, err := sdk.PayoutHistory(&PayoutHistoryRequest{
//		DateFrom: "2022-01-01 00:00:00",
//		DateTo: "2022-01-31 23:59:59",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) PayoutHistory(payload *PayoutHistoryRequest) (*PayoutHistoryResponse, error) {
	return sdk.PayoutHistoryWithContext(context.Background(), payload)
}

// PayoutHistoryWithContext returns a list of payouts.
//
// Details: https://doc.cryptomus.com/business/payouts/payout-history
//
// Example:
//
//	result, err := sdk.PayoutHistoryWithContext(ctx, &PayoutHistoryRequest{
//		DateFrom: "2022-01-01 00:00:00",
//		DateTo: "2022-01-31 23:59:59",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) PayoutHistoryWithContext(ctx context.Context, payload *PayoutHistoryRequest) (*PayoutHistoryResponse, error) {
	var result PayoutHistoryResponse

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PayoutToken, payload)).
		SetBody(payload).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(PayoutHistoryEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}

type PayoutListOfServiceLimit struct {
	MinAmount string `json:"min_amount"`
	MaxAmount string `json:"max_amount"`
}

type PayoutListOfServiceCommission struct {
	FeeAmount string `json:"fee_amount"`
	Percent   string `json:"percent"`
}

type PayoutListOfServicesData struct {
	Network     string                         `json:"network"`
	Currency    string                         `json:"currency"`
	IsAvailable bool                           `json:"is_available"`
	Limit       PaymentListOfServiceLimit      `json:"limit"`
	Commission  PaymentListOfServiceCommission `json:"commission"`
}

type PayoutListOfServicesResponse struct {
	*HTTPResponse
	Result []*PayoutListOfServicesData `json:"result"`
}

// PayoutListOfServices returns a list of services for payouts.
//
// Details: https://doc.cryptomus.com/business/payouts/list-of-services
//
// Returns a list of available payout services. Payout services store settings that are taken into account when creating a payout. For example. currencies, networks, minimum and maximum limits, commissions.
//
// Example:
//
//	result, err := sdk.PayoutListOfServices()
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) PayoutListOfServices() (*PayoutListOfServicesResponse, error) {
	return sdk.PayoutListOfServicesWithContext(context.Background())
}

// PayoutListOfServicesWithContext returns a list of services for payouts.
//
// Details: https://doc.cryptomus.com/business/payouts/list-of-services
//
// Returns a list of available payout services. Payout services store settings that are taken into account when creating a payout. For example. currencies, networks, minimum and maximum limits, commissions.
//
// Example:
//
//	result, err := sdk.PayoutListOfServicesWithContext(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) PayoutListOfServicesWithContext(ctx context.Context) (*PayoutListOfServicesResponse, error) {
	var result PayoutListOfServicesResponse

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PayoutToken, nil)).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Get(PayoutListOfServicesEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}

type TransferToPersonalWalletRequest struct {
	Amount   string `json:"amount" validate:"required"`
	Currency string `json:"currency" validate:"required"`
}

type TransferToPersonalWalletData struct {
	UserWalletTransactionUUID string `json:"user_wallet_transaction_uuid"`
	UserWalletBalance         string `json:"user_wallet_balance"`
	MerchantTransactionUUID   string `json:"merchant_transaction_uuid"`
	MerchantBalance           string `json:"merchant_balance"`
}

type TransferToPersonalWalletResponse struct {
	*HTTPResponse
	Result *TransferToPersonalWalletData `json:"result"`
}

// TransferToPersonalWallet transfers funds from a business wallet to a personal wallet.
//
// Details: https://doc.cryptomus.com/business/payouts/transfer-to-personal
//
// Example:
//
//	result, err := sdk.TransferToPersonalWallet(&TransferToPersonalWalletRequest{
//		Amount:   "5",
//		Currency: "USDT",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) TransferToPersonalWallet(payload *TransferToPersonalWalletRequest) (*TransferToPersonalWalletResponse, error) {
	return sdk.TransferToPersonalWalletWithContext(context.Background(), payload)
}

// TransferToPersonalWalletWithContext transfers funds from a business wallet to a personal wallet.
//
// Details: https://doc.cryptomus.com/business/payouts/transfer-to-personal
//
// Example:
//
//	result, err := sdk.TransferToPersonalWalletWithContext(ctx, &TransferToPersonalWalletRequest{
//		Amount:   "5",
//		Currency: "USDT",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) TransferToPersonalWalletWithContext(ctx context.Context, payload *TransferToPersonalWalletRequest) (*TransferToPersonalWalletResponse, error) {
	var result TransferToPersonalWalletResponse

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PayoutToken, payload)).
		SetBody(payload).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(TransferToPersonalWalletEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}

type TransferToBusinessWalletRequest struct {
	Amount   string `json:"amount" validate:"required"`
	Currency string `json:"currency" validate:"required"`
}

type TransferToBusinessWalletData struct {
	UserWalletTransactionUUID string `json:"user_wallet_transaction_uuid"`
	UserWalletBalance         string `json:"user_wallet_balance"`
	MerchantTransactionUUID   string `json:"merchant_transaction_uuid"`
	MerchantBalance           string `json:"merchant_balance"`
}

type TransferToBusinessWalletResponse struct {
	*HTTPResponse
	Result *TransferToBusinessWalletData `json:"result"`
}

// TransferToBusinessWallet transfers funds from a personal wallet to a business wallet.
//
// Details: https://doc.cryptomus.com/business/payouts/transfer-to-business
//
// Example:
//
//	result, err := sdk.TransferToBusinessWallet(&TransferToBusinessWalletRequest{
//		Amount:   "5",
//		Currency: "USDT",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) TransferToBusinessWallet(payload *TransferToBusinessWalletRequest) (*TransferToBusinessWalletResponse, error) {
	return sdk.TransferToBusinessWalletWithContext(context.Background(), payload)
}

// TransferToBusinessWalletWithContext transfers funds from a personal wallet to a business wallet.
//
// Details: https://doc.cryptomus.com/business/payouts/transfer-to-business
//
// Example:
//
//	result, err := sdk.TransferToBusinessWalletWithContext(ctx, &TransferToBusinessWalletRequest{
//		Amount:   "5",
//		Currency: "USDT",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) TransferToBusinessWalletWithContext(ctx context.Context, payload *TransferToBusinessWalletRequest) (*TransferToBusinessWalletResponse, error) {
	var result TransferToBusinessWalletResponse

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PayoutToken, payload)).
		SetBody(payload).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(TransferToBusinessWalletEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}
