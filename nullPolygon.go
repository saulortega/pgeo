package pgeo

import (
	"database/sql/driver"
	"encoding/json"
)

type NullPolygon struct {
	Polygon
	Valid bool
}

func (p NullPolygon) Value() (driver.Value, error) {
	if !p.Valid {
		return nil, nil
	}

	return valuePolygon(p.Polygon)
}

func (p *NullPolygon) Scan(src interface{}) error {
	if src == nil {
		p.Polygon, p.Valid = NewPolygon([]Point{Point{}, Point{}, Point{}, Point{}}), false
		return nil
	}

	p.Valid = true
	return scanPolygon(&p.Polygon, src)
}

func (p *NullPolygon) MarshalJSON() ([]byte, error) {
	if !p.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(p.Polygon)
}

func (p *NullPolygon) UnmarshalJSON(data []byte) error {
	if string(data) == "" || string(data) == "null" {
		p.Valid = false
		return nil
	}

	var err = json.Unmarshal(data, p.Polygon)
	p.Valid = err == nil
	return err
}
