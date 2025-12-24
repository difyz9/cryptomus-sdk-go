package cryptomus

type Endpoint string

var (
	Host = "https://api.cryptomus.com"
)

var (
	// 支付相关

	CreateInvoiceEndpoint                 Endpoint = "/v1/payment"
	CreateStaticWalletEndpoint            Endpoint = "/v1/wallet"
	GenerateQRCodeWalletEndpoint          Endpoint = "/v1/wallet/qr"
	GenerateQRCodeInvoiceEndpoint         Endpoint = "/v1/payment/qr"
	BlockStaticWalletEndpoint             Endpoint = "/v1/wallet/block-address"
	RefundPaymentOnBlockedAddressEndpoint Endpoint = "/v1/wallet/blocked-address-refund"
	PaymentInformationEndpoint            Endpoint = "/v1/payment/info"
	RefundEndpoint                        Endpoint = "/v1/payment/refund"
	ResendWebhookEndpoint                 Endpoint = "/v1/payment/resend"
	PaymentListOfServicesEndpoint         Endpoint = "/v1/payment/services"
	PaymentHistoryEndpoint                Endpoint = "/v1/payment/list"

	// Webhook 相关

	TestingWebhookPaymentEndpoint Endpoint = "/v1/test-webhook/payment"
	TestingWebhookPayoutEndpoint  Endpoint = "/v1/test-webhook/payout"
	TestingWebhookWalletEndpoint  Endpoint = "/v1/test-webhook/wallet"

	// 提现相关

	CreatePayoutEndpoint             Endpoint = "/v1/payout"
	PayoutInformationEndpoint        Endpoint = "/v1/payout/info"
	PayoutHistoryEndpoint            Endpoint = "/v1/payout/list"
	PayoutListOfServicesEndpoint     Endpoint = "/v1/payout/services"
	TransferToPersonalWalletEndpoint Endpoint = "/v1/transfer/to-personal"
	TransferToBusinessWalletEndpoint Endpoint = "/v1/transfer/to-business"

	// 循环支付相关

	CreateRecurringPaymentEndpoint      Endpoint = "/v1/recurrence/create"
	RecurringPaymentInformationEndpoint Endpoint = "/v1/recurrence/info"
	ListRecurringPaymentsEndpoint       Endpoint = "/v1/recurrence/list"
	CancelRecurringPaymentEndpoint      Endpoint = "/v1/recurrence/cancel"

	// 汇率相关

	ExchangeRateListEndpoint Endpoint = "/v1/exchange-rate/%s/list"

	// 折扣支付相关

	ListOfDiscountsEndpoint            Endpoint = "/v1/payment/discount/list"
	SetDiscountToPaymentMethodEndpoint Endpoint = "/v1/payment/discount/set"

	// 余额相关

	BalanceEndpoint Endpoint = "/v1/balance"
)

func (e Endpoint) String() string {
	return string(e)
}

func (e Endpoint) URL() string {
	return Host + e.String()
}
