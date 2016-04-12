package geom

func IsNullGeometry(geom Geometry) bool {
    if g, ok := IsPoint(geom); ok {
        return g == nil
    }else if g, ok := IsLineString(geom); ok {
        return g == nil
    }else if g, ok := IsPolygon(geom); ok {
        return g == nil
    }
    //type unknown treat as null
    return true
}

//Is point
func IsPoint(g Geometry) (*Point, bool) {
    pt, ok := g.(*Point)
    return pt, ok
}

//Is polygon
func IsPolygon(g Geometry) (*Polygon, bool) {
    ply, ok := g.(*Polygon)
    return ply, ok
}

//Is linestring
func IsLineString(g Geometry) (*LineString, bool) {
    ln, ok := g.(*LineString)
    return ln, ok
}

//Is linearing
func IsLinearRing(g Geometry) (*LinearRing, bool) {
    ln, ok := g.(*LinearRing)
    return ln, ok
}

//Type Values
type values []float64

func (s values) Len() int {
    return len(s)
}

func (s values) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s values) Less(i, j int) bool {
    return s[i] < s[j]
}
