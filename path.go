package pgeo

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"regexp"
)

type Path struct {
	Points []Point
	Closed bool `json:"closed"`
}

func (p Path) Value() (driver.Value, error) {
	return valuePath(p)
}

func (p *Path) Scan(src interface{}) error {
	return scanPath(p, src)
}

func valuePath(p Path) (driver.Value, error) {
	var val string
	if p.Closed {
		val = fmt.Sprintf(`(%s)`, formatPoints(p.Points))
	} else {
		val = fmt.Sprintf(`[%s]`, formatPoints(p.Points))
	}
	return val, nil
}

func scanPath(p *Path, src interface{}) error {
	if src == nil {
		return nil
	}

	val, err := iToS(src)
	if err != nil {
		return err
	}

	(*p).Points, err = parsePoints(val)
	if err != nil {
		return err
	}

	if len((*p).Points) < 2 {
		return errors.New("wrong path")
	}

	(*p).Closed = regexp.MustCompile(`^\(\(`).MatchString(val)

	return nil
}
