# Cryptomus SDK Go

[![Go Reference](https://pkg.go.dev/badge/github.com/difyz9/cryptomus-sdk-go.svg)](https://pkg.go.dev/github.com/difyz9/cryptomus-sdk-go)
[![Go Version](https://img.shields.io/github/go-mod/go-version/Aldiwildan77/cryptomus-sdk-go)](https://golang.org/dl/)
[![License](https://img.shields.io/github/license/Aldiwildan77/cryptomus-sdk-go)](LICENSE)

Cryptomus SDK Go 是一个用于与 Cryptomus API 交互的 Go 语言库。Cryptomus 是一个加密货币支付网关，支持多种加密货币的支付和提现功能。

完整的 API 文档请访问：[https://doc.cryptomus.com/](https://doc.cryptomus.com/)

## ✨ 功能特性

- 💰 **支付集成** - 创建发票、查询发票、获取发票列表、取消发票等
- 💸 **提现集成** - 创建提现、查询提现、获取提现列表、取消提现等
- 🔄 **循环支付集成** - 创建循环支付、查询循环支付、获取循环支付列表、取消循环支付等
- 🔔 **Webhook 集成** - 验证 Webhook、重发 Webhook
- 👛 **静态钱包** - 生成固定的加密货币收款地址
- 📊 **余额查询** - 查询商户余额信息
- 💱 **汇率查询** - 获取加密货币实时汇率
- 🎁 **折扣管理** - 设置和管理支付方式折扣

## 📦 安装

使用以下命令安装 Cryptomus SDK Go：

```bash
go get github.com/difyz9/cryptomus-sdk-go
```

## 🚀 快速开始

### 初始化 SDK

```go
package main

import (
    cryptomus "github.com/difyz9/cryptomus-sdk-go"
)

func main() {
    // 创建 SDK 实例
    sdk := cryptomus.New(
        cryptomus.WithMerchant("your-merchant-id"),
        cryptomus.WithPaymentToken("your-payment-token"),
        cryptomus.WithPayoutToken("your-payout-token"), // 可选，提现功能需要
    )
}
```

### 创建发票

```go
result, err := sdk.CreateInvoice(&cryptomus.CreateInvoiceRequest{
    Amount:   "15",
    Currency: "USD",
    OrderID:  "123456",
    Lifetime: 3600,
})

if err != nil {
    log.Fatal(err)
}

log.Println("发票创建成功")
log.Printf("发票信息: %#+v", result)
```

### 查询余额

```go
balance, err := sdk.Balance()
if err != nil {
    log.Fatal(err)
}

log.Printf("余额信息: %#+v", balance)
```

### 获取汇率

```go
rates, err := sdk.ExchangeRateList("BTC")
if err != nil {
    log.Fatal(err)
}

log.Printf("BTC 汇率: %#+v", rates)
```

### 创建提现

```go
payout, err := sdk.CreatePayout(&cryptomus.CreatePayoutRequest{
    Amount:      "5",
    Currency:    "USDT",
    Network:     "TRON",
    Address:     "TJ5Zrj8z6bJ7bk9Kf8fz1yQFbJ7b7bJ7b7b",
    OrderID:     "payout-001",
    URLCallback: "https://example.com/callback",
    IsSubtract:  true,
})

if err != nil {
    log.Fatal(err)
}

log.Printf("提现创建成功: %#+v", payout)
```

## 📚 更多示例

在 [examples](examples) 目录中可以找到更多完整的示例代码：

- [创建发票](examples/payments/create-invoice/)
- [查询余额](examples/balance/)
- [获取汇率](examples/exchange-rates/)
- [循环支付](examples/recurring-payments/)
- [折扣管理](examples/discount-payment/)

## 🔧 高级用法

### 使用 Context

所有 API 方法都提供了带 Context 的版本，支持超时控制和取消操作：

```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

result, err := sdk.CreateInvoiceWithContext(ctx, &cryptomus.CreateInvoiceRequest{
    Amount:   "15",
    Currency: "USD",
    OrderID:  "123456",
})
```

### 自定义 HTTP 客户端

```go
import (
    "github.com/imroc/req/v3"
    cryptomus "github.com/difyz9/cryptomus-sdk-go"
)

client := req.C().
    SetTimeout(30 * time.Second).
    SetCommonRetryCount(3)

sdk := cryptomus.New(
    cryptomus.WithHttpClient(client),
    cryptomus.WithMerchant("your-merchant-id"),
    cryptomus.WithPaymentToken("your-payment-token"),
)
```

## 🔐 安全性

- 所有请求都使用 HMAC-MD5 签名验证
- 支持 Webhook 签名验证
- 建议在生产环境中使用环境变量存储 API 密钥

```go
sdk := cryptomus.New(
    cryptomus.WithMerchant(os.Getenv("CRYPTOMUS_MERCHANT_ID")),
    cryptomus.WithPaymentToken(os.Getenv("CRYPTOMUS_PAYMENT_TOKEN")),
)
```

## 📖 API 文档

详细的 API 文档和参数说明，请参考：

- [Cryptomus 官方文档](https://doc.cryptomus.com/)
- [Go Package 文档](https://pkg.go.dev/github.com/difyz9/cryptomus-sdk-go)

## 🤝 贡献

欢迎贡献代码！请随时提交 Pull Request。

在提交 PR 之前，请确保：

1. 代码通过了所有测试
2. 添加了必要的注释和文档
3. 遵循 Go 语言编码规范

## 📝 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件。

## 👤 作者

[Muhammad Wildan Aldiansyah](https://aldiwildan.me)

## 🙏 致谢

感谢 [Cryptomus](https://cryptomus.com/) 提供优秀的加密货币支付服务。

---

如有问题或建议，请提交 [Issue](https://github.com/difyz9/cryptomus-sdk-go/issues)。
