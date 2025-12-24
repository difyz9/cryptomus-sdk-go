# Cryptomus SDK Go

[![Go Reference](https://pkg.go.dev/badge/github.com/difyz9/cryptomus-sdk-go.svg)](https://pkg.go.dev/github.com/difyz9/cryptomus-sdk-go)
[![Go Version](https://img.shields.io/github/go-mod/go-version/Aldiwildan77/cryptomus-sdk-go)](https://golang.org/dl/)
[![License](https://img.shields.io/github/license/Aldiwildan77/cryptomus-sdk-go)](LICENSE)

> 🚀 强大、易用、类型安全的 Cryptomus 加密货币支付网关 Go SDK

Cryptomus SDK Go 是一个用于与 Cryptomus API 交互的 Go 语言库。Cryptomus 是一个专业的加密货币支付网关，支持多种加密货币的支付、提现和循环订阅功能。

**官方文档**：[https://doc.cryptomus.com/](https://doc.cryptomus.com/)

## 📑 目录

- [功能特性](#-功能特性)
- [环境要求](#-环境要求)
- [安装](#-安装)
- [快速开始](#-快速开始)
- [更多示例](#-更多示例)
- [高级用法](#-高级用法)
- [Webhook 集成](#-webhook-集成)
- [错误处理](#-错误处理)
- [常见问题](#-常见问题)
- [故障排查](#-故障排查)
- [API 文档](#-api-文档)
- [贡献指南](#-贡献)
- [许可证](#-许可证)

## ✨ 功能特性

### 核心功能

- 💰 **支付集成**
  - 创建一次性支付发票
  - 创建静态收款钱包
  - 生成支付二维码
  - 查询支付状态和详情
  - 获取支付历史记录
  - 支付退款功能

- 💸 **提现集成**
  - 创建提现请求
  - 查询提现状态
  - 获取提现历史
  - 提现到多种区块链网络
  - 自动汇率转换

- 🔄 **循环支付（订阅）**
  - 创建循环支付计划
  - 支持周/月/季度周期
  - 管理订阅状态
  - 取消循环支付

- 🔔 **Webhook 集成**
  - 实时支付通知
  - 签名验证
  - 重发 Webhook
  - 测试 Webhook

- 📊 **账户管理**
  - 查询商户余额
  - 查询个人余额
  - 多币种余额查询

- 💱 **汇率服务**
  - 实时加密货币汇率
  - 法币兑换率
  - 自动价格转换

- 🎁 **营销工具**
  - 设置支付折扣
  - 管理促销活动
  - 币种优惠配置

### 技术特性

- ✅ 完整的类型安全
- ✅ Context 超时控制
- ✅ 自动签名验证
- ✅ 详细的错误信息
- ✅ 并发安全
- ✅ 可自定义 HTTP 客户端
- ✅ 完整的中文文档和注释

## 🔧 环境要求

- **Go 版本**: 1.18 或更高
- **依赖**: 
  - `github.com/imroc/req/v3` - HTTP 客户端库
- **API 密钥**: 需要从 [Cryptomus 商户后台](https://cryptomus.com/) 获取

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
    � Webhook 集成

### 验证 Webhook 签名

```go
import (
    "crypto/md5"
    "encoding/base64"
    "encoding/hex"
    "encoding/json"
    "io"
    "net/http"
)

func handleWebhook(w http.ResponseWriter, r *http.Request) {
    // 读取请求体
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    
    // 获取签名
    sign := r.Header.Get("sign")
    
    // 验证签名
    encodedData := base64.StdEncoding.EncodeToString(body)
    expectedSign := md5Hash(encodedData + os.Getenv("CRYPTOMUS_PAYMENT_TOKEN"))
    
    if sign != expectedSign {
        http.Error(w, "Invalid signature", http.StatusUnauthorized)
        return
    }
    
    // 解析 Webhook 数据
    var webhookData map[string]interface{}
    if err := json.Unmarshal(body, &webhookData); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
    
    // 处理支付状态
    status := webhookData["status"].(string)
    switch status {
    case "paid":
        // 处理支付成功
        log.Println("支付成功！")
    case "wrong_amount":
        // 处理金额错误
        log.Println("支付金额不正确")
    case "cancel":
        // 处理取消支付
        log.Println("支付已取消")
    }
    
    w.WriteHeader(http.StatusOK)
}

func md5Hash(text string) string {
    hash := md5.Sum([]byte(text))
    return hex.EncodeToString(hash[:])
}
```

### 重发 Webhook

```go
### 官方资源

- [Cryptomus 官方文档](https://doc.cryptomus.com/) - 完整的 API 参考
- [Go Package 文档](https://pkg.go.dev/github.com/difyz9/cryptomus-sdk-go) - SDK API 文档
- [商户后台](https://cryptomus.com/merchant) - 获取 API 密钥和管理账户
- [支持中心](https://cryptomus.com/support) - 技术支持

我们欢迎并感谢所有形式的贡献！

### 贡献方式

- 🐛 [报告 Bug](https://github.com/difyz9/cryptomus-sdk-go/issues)
- 💡 [提出新功能建议](https://github.com/difyz9/cryptomus-sdk-go/issues)
- 📝 改进文档
- 🔧 提交代码修复
- ⭐ Star 本项目

### 提交 Pull Request

在提交 PR 之前，请确保：

1. ✅ 代码通过 `go test` 测试
2. ✅ 运行 `go fmt` 格式化代码
3. ✅ 运行 `go vet` 检查代码
4. ✅ 添加了必要的注释和文档
5. ✅ 更新了相关的 README 部分
6. ✅ 遵循 Go 语言编码规范

### 开发环境设置

```bash
# 克隆仓库
git clone https://github.com/difyz9/cryptomus-sdk-go.git
cd cryptomus-sdk-go

# 安装依赖
go mod download

# 运行测试
go test -v ./...

# 检查代码
go vet ./...
go fmt ./...
```ry` - 获取支付历史
- `Refund` - 退款
- `GenerateQRStaticWallet` / `GenerateQRCodeInvoice` - 生成二维码

#### 提现相关
- `CreatePayout` - 创建提现
- `PayoutInformation` - 查询提现信息
- `PayoutHistory` - 获取提现历史

#### 循环支付
- `CreateRecurringPayment` - 创建循环支付
- `RecurringPaymentInformation` - 查询循环支付信息
- `ListRecurringPayments` - 获取循环支付列表

#### 其他
- `Balance` - 查询余额
- `ExchangeRateList` - 获取汇率
- `ListOfDiscount` / `SetDiscountToPaymentMethod` - 折扣管理
- `ResendWebhook` - 重发 Webhook

if err != nil {
    log.Fatal(err)
}
```

## ⚠️ 错误处理

### 检查 API 错误

```go
result, err := sdk.CreateInvoice(&cryptomus.CreateInvoiceRequest{
    Amount:   "15",
    Currency: "USD",
    OrderID:  "123456",
})

if err != nil {
    log.Printf("网络错误: %v", err)
    return
}

// 检查 API 返回的业务错误
if result.State != 0 || result.Code != 0 {
    log.Printf("API 错误 - State: %d, Code: %d, Error: %s", 
        result.State, result.Code, result.Error)
    
    // 处理具体错误
    if result.Errors != nil {
        for field, messages := range result.Errors {
            log.Printf("字段 %s 错误: %v", field, messages)
        }
    }
    return
}

// 成功处理
log.Printf("发票创建成功: %s", result.Result.UUID)
```

### 常见错误码

| 错误码 | 说明 | 解决方案 |
|--------|------|----------|
| 401 | 签名验证失败 | 检查 API 密钥是否正确 |
| 400 | 请求参数错误 | 检查必填参数和参数格式 |
| 404 | 资源不存在 | 检查 UUID 或 OrderID 是否正确 |
| 429 | 请求过于频繁 | 实施请求限流 |
| 500 | 服务器错误 | 稍后重试或联系技术支持 |

## 🔐 安全性

### 最佳实践

1. **使用环境变量存储敏感信息**

```go
sdk := cryptomus.New(
    cryptomus.WithMerchant(os.Getenv("CRYPTOMUS_MERCHANT_ID")),
    cryptomus.WithPaymentToken(os.Getenv("CRYPTOMUS_PAYMENT_TOKEN")),
    cryptomus.WithPayoutToken(os.Getenv("CRYPTOMUS_PAYOUT_TOKEN")),
)
```

2. **始终验证 Webhook 签名**

所有 Webhook 请求都使用 HMAC-MD5 签名，必须验证签名以确保请求来自 Cryptomus。

3. **使用 HTTPS**

在生产环境中，回调 URL 必须使用 HTTPS。

4. **设置合理的超时时间**

```go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

result, err := sdk.CreateInvoiceWithContext(ctx, request)
```

5. **不要在日志中输出敏感信息**

避免记录 API 密钥、签名或完整的请求/响应数据。

## ❓ 常见问题

### Q: 支持哪些加密货币？

A: Cryptomus 支持 BTC、ETH、USDT、USDC、LTC、DOGE、TRX、BNB 等多种主流加密货币，以及多个区块链网络（如 Ethereum、BSC、Polygon、Tron、Arbitrum 等）。具体列表请查看[官方文档](https://doc.cryptomus.com/)。

### Q: 如何测试集成？

A: Cryptomus 提供测试网络支持。您可以：
1. 在测试环境创建商户账户
2. 使用测试网络的 API 密钥
3. 使用 `TestingWebhookPayment` 方法测试 Webhook

```go
// 测试 Webhook
result, err := sdk.TestingWebhookPayment(&cryptomus.TestingWebhookPaymentRequest{
    URLCallback: "https://your-domain.com/webhook",
    Currency:    "USDT",
    Network:     "tron",
    Status:      "paid",
})
```

### Q: 提现需要多长时间？

A: 提现处理时间取决于区块链网络：
- TRX/USDT(TRC20): 通常 1-5 分钟
- BTC: 10-60 分钟
- ETH/USDT(ERC20): 5-30 分钟
- BSC: 1-5 分钟

### Q: 如何处理支付金额误差？

A: 使用 `AccuracyPaymentPercent` 参数设置允许的支付金额误差范围（百分比）：

```go
result, err := sdk.CreateInvoice(&cryptomus.CreateInvoiceRequest{
    Amount:                 "100",
    Currency:               "USD",
    AccuracyPaymentPercent: 1, // 允许 ±1% 的误差
})
```

### Q: 如何实现定时任务检查支付状态？

A: 推荐使用 Webhook 接收实时通知。如果需要轮询：

```go
import "time"

func checkPaymentStatus(sdk *cryptomus.Cryptomus, orderID string) {
    ticker := time.NewTicker(30 * time.Second)
    defer ticker.Stop()
    
    timeout := time.After(1 * time.Hour)
    
    for {
        select {
        case <-ticker.C:
            info, err := sdk.PaymentInformation(&cryptomus.PaymentInformationRequest{
                OrderID: orderID,
            })
            if err != nil {
                log.Printf("查询失败: %v", err)
                continue
            }
            
            if info.Result.IsFinal {
                log.Printf("支付已完成，状态: %s", info.Result.PaymentStatus)
                return
            }
            
        case <-timeout:
            log.Println("支付超时")
            return
        }
    }
}
```

### Q: 能否自定义发票过期时间？

A: 可以，使用 `Lifetime` 参数（秒）：

```go
result, err := sdk.CreateInvoice(&cryptomus.CreateInvoiceRequest{
    Amount:   "100",
    Currency: "USD",
    Lifetime: 3600, // 1 小时后过期
})
```

## 🔍 故障排查

### 问题：签名验证失败

**症状**: 收到 401 错误或签名不匹配

**解决方案**:
1. 确认使用正确的 API 密钥
2. 检查是否混用了支付密钥和提现密钥
3. 确保请求体和签名计算使用相同的数据
4. 检查字符编码是否为 UTF-8

### 问题：Webhook 未收到通知

**症状**: 支付完成但未触发 Webhook

**解决方案**:
1. 确认创建发票时设置了 `URLCallback`
2. 检查回调 URL 是否可公网访问
3. 确保回调服务器返回 200 状态码
4. 查看 Cryptomus 后台的 Webhook 日志
5. 使用 `ResendWebhook` 手动重发

### 问题：金额精度问题

**症状**: 金额计算出现小数点偏差

**解决方案**:
1. 使用字符串类型传递金额，避免浮点数精度问题
2. 设置 `AccuracyPaymentPercent` 允许合理的误差范围
3. 后端使用 `decimal` 库处理货币计算

```go
// ✅ 正确
Amount: "100.50"

// ❌ 错误
Amount: fmt.Sprintf("%f", 100.50) // 可能产生精度问题 [创建发票](examples/payments/create-invoice/)
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
    � 项目状态

- ✅ 稳定版本
- ✅ 生产环境可用
- ✅ 持续维护
- ✅ 完整中文文档

## 🔗 相关链接

- [GitHub 仓库](https://github.com/difyz9/cryptomus-sdk-go)
- [问题反馈](https://github.com/difyz9/cryptomus-sdk-go/issues)
- [贡献指南](https://github.com/difyz9/cryptomus-sdk-go/pulls)

## 👤 作者

原作者：[Muhammad Wildan Aldiansyah](https://aldiwildan.me)

维护者：difyz9

## 🙏 致谢

- 感谢 [Cryptomus](https://cryptomus.com/) 提供优秀的加密货币支付服务
- 感谢所有贡献者的支持和帮助
- 感谢 [req](https://github.com/imroc/req) 提供强大的 HTTP 客户端库

## 📄 变更日志

查看 [CHANGELOG.md](CHANGELOG.md) 了解版本更新历史。

## ⚖️ 免责声明

本 SDK 仅供学习和开发使用。使用本 SDK 进行加密货币交易时，请确保：

1. 遵守当地法律法规
2. 了解加密货币交易风险
3. 妥善保管 API 密钥
4. 在生产环境充分测试

---

**需要帮助？**
- 📧 提交 [Issue](https://github.com/difyz9/cryptomus-sdk-go/issues)
- 💬 查看 [常见问题](#-常见问题)
- 📖 阅读 [官方文档](https://doc.cryptomus.com/)

**觉得有用？** 请给我们一个 ⭐ Star！
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
