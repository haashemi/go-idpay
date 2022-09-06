package idpay

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type VerifyTransaction struct {
	Status  Number    `json:"status"`   // Status is Transaction Status
	TrackID Number    `json:"track_id"` // TrackIS is IDPay's track id
	ID      string    `json:"id"`       // ID is Transaction ID which generated on creating the transaction
	OrderID string    `json:"order_id"` // OrderID is the orderID which sent on creating the transaction
	Amount  Number    `json:"amount"`   // Amount is submitted amount on creating the transaction
	Date    time.Time `json:"date"`     // Date is the transaction creation time
	Payment struct {
		TrackID      string    `json:"track_id"`       // TrackIS is IDPay's track id
		Amount       Number    `json:"amount"`         // Amount is the amount payable
		CardNo       string    `json:"card_no"`        // CardNo is payer's card number in the format of 123456******1234
		HashedCardNo string    `json:"hashed_card_no"` // HashedCardNo is hashed card number with SHA256 algorithm
		Date         time.Time `json:"date"`           // Date is the transaction payment time
	} `json:"payment"` // Payment contains the transaction's payment info
	Verify struct {
		Date time.Time `json:"date"` // Date is the transaction verification time
	} `json:"verify"` // Verify contains payment verification info
}

// VerifyTransaction verifies the transaction to complete the payment.
//
// id is the unique transaction id from the CreateTransaction response
//
// orderID is the orderID sent on the CreateTransaction step
func (i IDPay) VerifyTransaction(id, orderID string) (*VerifyTransaction, error) {
	body := map[string]string{
		"id":       id,
		"order_id": orderID,
	}

	payload, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, HOST+"/payment/verify", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", i.apiKey)
	req.Header.Set("X-SANDBOX", i.IsSandbox())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, getError(req.Body)
	}

	data := &VerifyTransaction{}
	err = json.NewDecoder(resp.Body).Decode(data)
	return data, err
}
