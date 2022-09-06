package idpay

import (
	"strconv"
)

type Number int64

func (c *Number) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(strconv.FormatInt(int64(*c), 10))), nil
}

func (c *Number) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		s = string(b)
	}
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*c = Number(i)
	return nil
}

func (c Number) Int64() int64 {
	return int64(c)
}
