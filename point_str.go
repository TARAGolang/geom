package geom


//String creates a wkt string format of point
func (self *Point) String() string {
    return WriteWKT(
        NewWKTParserObj(GeoType_Point, [][2]float64{self.ToArray()}),
    )
}