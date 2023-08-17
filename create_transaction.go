package idpay

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type NewTransaction struct {
	ID   string `json:"id"`
	Link string `json:"link"`
}

type TransactionOptions struct {
	// Name is Payer name. up to 255 characters long accepted.
	Name string

	// Phone is payer phone number. should be 11 characters long.
	// valid examples: 9382198592, 09382198592, 989382198592
	Phone string

	// Mail is payer mail. up to 255 characters long accepted.
	Mail string

	// Description is transaction's description. up to 255 characters long accepted.
	Description string
}

// CreateTransaction created new idpay transaction
//
// orderID: Acceptor's order number. up to 50 characters.
//
// callbackURL:
//
// amount: The desired amount is in Rials. it must be between 1,000 and 500,000,000 Rials.
func (i IDPay) CreateTransaction(orderID, callbackURL string, amount int, opts *TransactionOptions) (*NewTransaction, error) {
	body := map[string]interface{}{
		"order_id": orderID,
		"amount":   amount,
		"callback": callbackURL,
	}

	if opts != nil {
		if opts.Name != "" {
			body["name"] = opts.Name
		}

		if opts.Phone != "" {
			body["phone"] = opts.Phone
		}

		if opts.Mail != "" {
			body["mail"] = opts.Mail
		}

		if opts.Description != "" {
			body["desc"] = opts.Description
		}
	}

	payload, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, HOST+"/payment", bytes.NewBuffer(payload))
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

	if resp.StatusCode != http.StatusCreated {
		return nil, getError(resp.Body)
	}

	data := &NewTransaction{}
	err = json.NewDecoder(resp.Body).Decode(data)
	return data, err
}
