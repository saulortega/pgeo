# pgeo
Package pgeo implements [geometric types for Postgres](https://www.postgresql.org/docs/current/static/datatype-geometric.html)

It uses Scanner and Valuer interfaces from `database/sql`


# Types
```go
//Points are the fundamental two-dimensional building block for geometric types.
//X and Y are the respective coordinates, as floating-point numbers
type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

//Circles are represented by a center point and radius.
type Circle struct {
	Point
	Radius float64 `json:"radius"`
}

//Line represents a infinite line with the linear equation Ax + By + C = 0, where A and B are not both zero.
type Line struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
	C float64 `json:"c"`
}

//Paths are represented by lists of connected points.
//Paths can be open, where the first and last points in the list are considered not connected,
//or closed, where the first and last points are considered connected.
type Path struct {
	Points []Point
	Closed bool `json:"closed"`
}

//Boxes are represented by pairs of points that are opposite corners of the box.
type Box [2]Point

//Line segments are represented by pairs of points that are the endpoints of the segment.
type Lseg [2]Point

//Polygons are represented by lists of points (the vertexes of the polygon).
type Polygon []Point
```

### Null Types
```go
type NullPoint struct {
	Point
	Valid bool `json:"valid"`
}

//Same for others...
```


# Functions

### Create Types
```go
NewPoint(X, Y float64) Point

NewLine(A, B, C float64) Line

NewLseg(A, B Point) Lseg

NewBox(A, B Point) Box

NewPath(P []Point, Closed bool) Path

NewPolygon(P []Point) Polygon

NewCircle(P Point, Radius float64) Circle
```

### Null Types
```go
NewNullPoint(P Point, valid bool) NullPoint

NewNullLine(L Line, valid bool) NullLine

NewNullLseg(L Lseg, valid bool) NullLseg

NewNullBox(B Box, valid bool) NullBox

NewNullPath(P Path, valid bool) NullPath

NewNullPolygon(P Polygon, valid bool) NullPolygon

NewNullCircle(C Circle, valid bool) NullCircle
```

### Rand Types
```go
NewRandPoint() Point

NewRandLine() Line

NewRandLseg() Lseg

NewRandBox() Box

NewRandPath() Path

NewRandPolygon() Polygon

NewRandCircle() Circle
```

### Zero Point
```go
//Returns a Point X=0, Y=0
NewZeroPoint() Point
```
