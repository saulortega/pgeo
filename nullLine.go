package pgeo

import (
	"database/sql/driver"
	"encoding/json"
)

type NullLine struct {
	Line
	Valid bool
}

func (l NullLine) Value() (driver.Value, error) {
	if !l.Valid {
		return nil, nil
	}

	return valueLine(l.Line)
}

func (l *NullLine) Scan(src interface{}) error {
	if src == nil {
		l.Line, l.Valid = NewLine(0, 0, 0), false
		return nil
	}

	l.Valid = true
	return scanLine(&l.Line, src)
}

func (l *NullLine) MarshalJSON() ([]byte, error) {
	if !l.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(l.Line)
}

func (l *NullLine) UnmarshalJSON(data []byte) error {
	if string(data) == "" || string(data) == "null" {
		l.Valid = false
		return nil
	}

	var err = json.Unmarshal(data, l.Line)
	l.Valid = err == nil
	return err
}
