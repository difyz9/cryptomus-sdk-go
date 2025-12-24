package cryptomus

import "encoding/json"

// ToJSON 将结构体转换为 JSON 字节数组。
//
// 示例：
//
//	type ExchangeRate struct {
//		From   string `json:"from,omitempty"`
//		To     string `json:"to,omitempty"`
//		Course string `json:"course,omitempty"`
//	}
//
//	exchangeRate := ExchangeRate{
//		From:   "ETH",
//		To:     "BTC",
//		Course: "0.0001",
//	}
//
//	resByte, err := ToJSON[ExchangeRate](exchangeRate)
//	if err != nil {
//		log.Fatal(err)
//	}
func ToJSON[T any](v T) ([]byte, error) {
	resByte, err := json.Marshal(v)
	if err != nil {
		return []byte{}, err
	}

	return resByte, nil
}

// FromJSONString 将 JSON 字符串转换为结构体。
//
// 示例：
//
//	type ExchangeRate struct {
//		From   string `json:"from,omitempty"`
//		To     string `json:"to,omitempty"`
//		Course string `json:"course,omitempty"`
//	}
//
//	exchangeRate, err := FromJSONString[ExchangeRate](`{"from":"ETH","to":"BTC","course":"0.0001"}`)
//	if err != nil {
//		log.Fatal(err)
//	}
func FromJSONString[T any](v string) (T, error) {
	var res T
	err := json.Unmarshal([]byte(v), &res)
	if err != nil {
		return res, err
	}

	return res, nil
}
