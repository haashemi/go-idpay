package idpay

const HOST = "https://api.idpay.ir/v1.1"

type IDPay struct {
	apiKey    string
	isSandbox bool
}

func New(apiKey string) *IDPay {
	return &IDPay{apiKey: apiKey}
}

func NewSandbox(apiKey string) *IDPay {
	return &IDPay{apiKey: apiKey, isSandbox: true}
}

func (i IDPay) IsSandbox() string {
	if i.isSandbox {
		return "1"
	}
	return "0"
}
