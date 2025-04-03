package kqueue

type ThirdPaymentUpdatePayStatusNotifyMessage struct {
	PayStatus int64  `json:"payStatus"`
	OrderSn   string `json:"orderSn"`
}
