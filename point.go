package pgeo

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func (p Point) Value() (driver.Value, error) {
	return fmt.Sprintf(`(%v,%v)`, p.X, p.Y), nil
}

func (p *Point) Scan(src interface{}) error {
	if src == nil {
		p.X, p.Y = 0, 0
		return nil
	}

	var val string
	var err error

	switch src.(type) {
	case string:
		val = src.(string)
	case []byte:
		val = string(src.([]byte))
	default:
		return errors.New("incompatible type for point")
	}

	ped := regexp.MustCompile(`^\((-?[0-9]+(?:\.[0-9]+)?),(-?[0-9]+(?:\.[0-9]+)?)\)$`).FindStringSubmatch(val)
	if len(ped) != 3 {
		return errors.New("wrong point")
	}

	p.X, err = strconv.ParseFloat(ped[1], 64)
	if err != nil {
		return err
	}

	p.Y, err = strconv.ParseFloat(ped[2], 64)
	if err != nil {
		return err
	}

	return nil
}
