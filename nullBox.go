package pgeo

import (
	"database/sql/driver"
	"encoding/json"
)

type NullBox struct {
	Box
	Valid bool
}

func (b NullBox) Value() (driver.Value, error) {
	if !b.Valid {
		return nil, nil
	}

	return valueBox(b.Box)
}

func (b *NullBox) Scan(src interface{}) error {
	if src == nil {
		b.Box, b.Valid = NewBox(Point{}, Point{}), false
		return nil
	}

	b.Valid = true
	return scanBox(&b.Box, src)
}

func (b *NullBox) MarshalJSON() ([]byte, error) {
	if !b.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(b.Box)
}

func (b *NullBox) UnmarshalJSON(data []byte) error {
	if string(data) == "" || string(data) == "null" {
		b.Valid = false
		return nil
	}

	var err = json.Unmarshal(data, b.Box)
	b.Valid = err == nil
	return err
}
