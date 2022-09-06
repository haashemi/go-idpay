package idpay

import (
	"strconv"
	"strings"
)

type Number int

func (jsonInt *Number) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(int(*jsonInt))), nil
}

func (jsonInt *Number) UnmarshalJSON(b []byte) error {
	intStr := strings.Replace(string(b), "\"", "", 2)

	number, err := strconv.Atoi(intStr)
	if err != nil {
		return err
	}

	*jsonInt = Number(number)
	return nil
}
