package pgeo

import (
	"database/sql/driver"
	"encoding/json"
)

type NullCircle struct {
	Circle
	Valid bool
}

func (c NullCircle) Value() (driver.Value, error) {
	if !c.Valid {
		return nil, nil
	}

	return valueCircle(c.Circle)
}

func (c *NullCircle) Scan(src interface{}) error {
	if src == nil {
		c.Circle, c.Valid = NewCircle(Point{}, 0), false
		return nil
	}

	c.Valid = true
	return scanCircle(&c.Circle, src)
}

func (c *NullCircle) MarshalJSON() ([]byte, error) {
	if !c.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(c.Circle)
}

func (c *NullCircle) UnmarshalJSON(data []byte) error {
	if string(data) == "" || string(data) == "null" {
		c.Valid = false
		return nil
	}

	var err = json.Unmarshal(data, c.Circle)
	c.Valid = err == nil
	return err
}
