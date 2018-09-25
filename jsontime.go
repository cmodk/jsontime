package jsontime

import (
	"strings"
	"time"
)

type JSONTime time.Time

func (jt *JSONTime) Time() time.Time {
	return time.Time(*jt)
}

func (jt *JSONTime) TimePtr() *time.Time {
	t := time.Time(*jt)

	return &t
}

type RFC3339Nano struct {
	JSONTime
}

func (ct *RFC3339Nano) UnmarshalJSON(input []byte) error {

	str := strings.Replace(string(input), "\"", "", -1)

	t, err := time.Parse(time.RFC3339Nano, str)
	if err != nil {
		return err
	}

	*ct = RFC3339Nano{JSONTime(t)}

	return nil
}

func (ct *RFC3339Nano) MarshalJSON() ([]byte, error) {
	t := ct.Time()
	return []byte(t.Format(time.RFC3339Nano)), nil
}

var (
	ISO8601Format = "2006-01-02T15:04:05.999999999"
)

type ISO8601 struct {
	JSONTime
}

func (ct *ISO8601) UnmarshalJSON(input []byte) error {

	//Need to strip string for "
	str := strings.Replace(string(input), "\"", "", -1)
	t, err := time.Parse(ISO8601Format, str)
	if err != nil {
		return err
	}

	*ct = ISO8601{JSONTime(t)}

	return nil
}

func (ct *ISO8601) MarshalJSON() ([]byte, error) {
	t := ct.Time()
	return []byte(t.Format(ISO8601Format)), nil
}
