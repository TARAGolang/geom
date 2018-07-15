package geom

import (
	"bytes"
	"strconv"
	"github.com/intdxdt/math"
)

//String creates a wkt string format of point
func (self *Point) WKT() string {
	var coords = self.ToArray()
	return WriteWKT(NewWKTParserObj(GeoType_Point, [][]float64{{coords[X], coords[Y]}}))
}

//String creates a wkt string format of point
func (self *Point) String() string {
	var buf bytes.Buffer
	buf.WriteString("[")
	buf.WriteString(strconv.FormatFloat(self[X], 'f', -1, 64) + ", ")
	buf.WriteString(strconv.FormatFloat(self[Y], 'f', -1, 64) + ", ")
	buf.WriteString(strconv.FormatFloat(self[Z], 'f', -1, 64))
	buf.WriteString("]")

	return buf.String()
}

//coordinate string
func coord_str(pt []float64) string {
	return math.FloatToString(pt[X]) + " " + math.FloatToString(pt[Y])
}
