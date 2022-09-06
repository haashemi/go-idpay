package idpay

import (
	"encoding/json"
	"io"
)

type Error struct {
	Code    int    `json:"error_code"`
	Message string `json:"error_message"`
}

func (e Error) Error() string {
	return e.Message
}

func getError(resp io.Reader) error {
	idpayErr := &Error{}

	if err := json.NewDecoder(resp).Decode(idpayErr); err != nil {
		return err
	}

	return idpayErr
}
