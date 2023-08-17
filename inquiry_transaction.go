package idpay

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type InquiryTransaction struct {
	Status  Number   `json:"status"`   // Status is Transaction Status
	TrackID Number   `json:"track_id"` // TrackIS is IDPay's track id
	ID      string   `json:"id"`       // ID is Transaction ID which generated on creating the transaction
	OrderID string   `json:"order_id"` // OrderID is the orderID which sent on creating the transaction
	Amount  Number   `json:"amount"`   // Amount is submitted amount on creating the transaction
	Date    UnixTime `json:"date"`     // Date is the transaction creation time
	Wage    struct {
		By     string `json:"by"`     // By is whose pays the wage. payer or payee
		Type   string `json:"type"`   // Type is the transaction's wage type. Amount, Percent, Stair
		Amount Number `json:"amount"` // Amount of the transaction's wage
	} `json:"wage"` // Wage is information about the transaction's wage
	Payer struct {
		Name  string `json:"name"`  // Name is Payer's name
		Phone string `json:"phone"` // Phone is Payer's phone number
		Mail  string `json:"mail"`  // Mail is Payer's mail address
		Desc  string `json:"desc"`  // Desc is Payer's description
	} `json:"payer"` // Payer contains the transaction's payer info
	Payment struct {
		TrackID      string   `json:"track_id"`       // TrackIS is IDPay's track id
		Amount       Number   `json:"amount"`         // Amount is the amount payable
		CardNo       string   `json:"card_no"`        // CardNo is payer's card number in the format of 123456******1234
		HashedCardNo string   `json:"hashed_card_no"` // HashedCardNo is hashed card number with SHA256 algorithm
		Date         UnixTime `json:"date"`           // Date is the transaction payment time
	} `json:"payment"` // Payment contains the transaction's payment info
	Verify struct {
		Date UnixTime `json:"date"` // Date is the transaction verification time
	} `json:"verify"` // Verify contains payment verification info
	Settlement struct {
		TrackID Number   `json:"track_id"` // TrackIS is IDPay's deposit track id
		Amount  Number   `json:"amount"`   // Amount to be deposited
		Date    UnixTime `json:"date"`     // Date is the time of depositing the transaction to the receiver's bank account
	} `json:"settlement"` // Settlement contains the transaction's settlement info
}

// InquiryTransaction returns the latest transaction status.
//
// id is the unique transaction id from the CreateTransaction response
//
// orderID is the orderID sent on the CreateTransaction step
func (i IDPay) InquiryTransaction(id, orderID string) (*InquiryTransaction, error) {
	body := map[string]string{
		"id":       id,
		"order_id": orderID,
	}

	payload, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, HOST+"/payment/inquiry", bytes.NewBuffer(payload))
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
		return nil, getError(resp.Body)
	}

	data := &InquiryTransaction{}
	err = json.NewDecoder(resp.Body).Decode(data)
	return data, err
}
