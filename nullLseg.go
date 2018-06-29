package pgeo

import (
	"database/sql/driver"
	"encoding/json"
)

type NullLseg struct {
	Lseg
	Valid bool
}

func (l NullLseg) Value() (driver.Value, error) {
	if !l.Valid {
		return nil, nil
	}

	return valueLseg(l.Lseg)
}

func (l *NullLseg) Scan(src interface{}) error {
	if src == nil {
		l.Lseg, l.Valid = NewLseg(Point{}, Point{}), false
		return nil
	}

	l.Valid = true
	return scanLseg(&l.Lseg, src)
}

func (l *NullLseg) MarshalJSON() ([]byte, error) {
	if !l.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(l.Lseg)
}

func (l *NullLseg) UnmarshalJSON(data []byte) error {
	if string(data) == "" || string(data) == "null" {
		l.Valid = false
		return nil
	}

	var err = json.Unmarshal(data, l.Lseg)
	l.Valid = err == nil
	return err
}
