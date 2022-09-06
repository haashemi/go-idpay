package idpay

import (
	"strconv"
	"time"
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

type UnixTime time.Time

func (c *UnixTime) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(strconv.FormatInt(time.Time(*c).Unix(), 10))), nil
}

func (c *UnixTime) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		s = string(b)
	}
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*c = UnixTime(time.Unix(i, 0))
	return nil
}

func (c UnixTime) AsTime() time.Time {
	return time.Time(c)
}
