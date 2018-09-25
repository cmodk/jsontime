package jsontime

import (
	"encoding/json"
	"log"
	"testing"
	"time"
)

func TestISO8601(t *testing.T) {
	str := "2018-08-30T17:30:00.000"

	pt, err := time.Parse(ISO8601Format, str)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("ISO8601: %s -> %s\n", str, pt.Format(time.RFC3339))
}

func TestISO8601_json(t *testing.T) {
	str := `{"timestamp": "2018-08-30T17:30:00.000"}`

	ts := struct {
		Timestamp ISO8601 `json:"timestamp"`
	}{}

	if err := json.Unmarshal([]byte(str), &ts); err != nil {
		t.Fatal(err)
	}

	log.Printf("ISO8601: %s -> %s\n", str, ts.Timestamp.Time().Format(time.RFC3339))

}
