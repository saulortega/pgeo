package pgeo

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
)

var closedPathRegexp = regexp.MustCompile(`^\(\(`)

// Path is represented by list of connected points.
// Paths can be open, where the first and last points in the list are considered not connected,
// or closed, where the first and last points are considered connected.
type Path struct {
	Points []Point
	Closed bool
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

	if len((*p).Points) < 1 {
		return errors.New("wrong path")
	}

	(*p).Closed = closedPathRegexp.MatchString(val)

	return nil
}

func (p *Path) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.Points)
}

func (p *Path) UnmarshalJSON(data []byte) error {
	var err = json.Unmarshal(data, p.Points)
	if p.Points != nil && len(p.Points) > 1 {
		p.Closed = p.Points[0] == p.Points[len(p.Points)-1]
	}

	return err
}
