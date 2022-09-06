package idpay

import "time"

type PostCallback struct {
	Status       Number    `json:"status"`         // Status is Transaction Status
	TrackID      Number    `json:"track_id"`       // TrackIS is IDPay's track id
	ID           string    `json:"id"`             // ID is Transaction ID which generated on creating the transaction
	OrderID      string    `json:"order_id"`       // OrderID is the orderID which sent on creating the transaction
	Amount       Number    `json:"amount"`         // Amount is submitted amount on creating the transaction
	CardNo       string    `json:"card_no"`        // CardNo is payer's card number in the format of 123456******1234
	HashedCardNo string    `json:"hashed_card_no"` // HashedCardNo is hashed card number with SHA256 algorithm
	Date         time.Time `json:"date"`           // Date is the transaction payment time
}

type GetCallback struct {
	Status  Number `json:"status"`   // Status is Transaction Status
	TrackID Number `json:"track_id"` // TrackIS is IDPay's track id
	ID      string `json:"id"`       // ID is Transaction ID which generated on creating the transaction
	OrderID string `json:"order_id"` // OrderID is the orderID which sent on creating the transaction
}
