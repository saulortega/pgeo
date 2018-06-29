package pgeo

import (
	"database/sql/driver"
	"encoding/json"
)

type NullPath struct {
	Path
	Valid bool
}

func (p NullPath) Value() (driver.Value, error) {
	if !p.Valid {
		return nil, nil
	}

	return valuePath(p.Path)
}

func (p *NullPath) Scan(src interface{}) error {
	if src == nil {
		p.Path, p.Valid = NewPath([]Point{Point{}, Point{}}, false), false
		return nil
	}

	p.Valid = true
	return scanPath(&p.Path, src)
}

func (p *NullPath) MarshalJSON() ([]byte, error) {
	if !p.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(p.Path)
}

func (p *NullPath) UnmarshalJSON(data []byte) error {
	if string(data) == "" || string(data) == "null" {
		p.Valid = false
		return nil
	}

	var err = json.Unmarshal(data, p.Path)
	p.Valid = err == nil
	return err
}
