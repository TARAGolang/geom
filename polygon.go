package geom

type Polygon struct {
	Shell *LinearRing
	Holes []*LinearRing
}

//New polygon from points
func NewPolygon(coordinates ...Coords) *Polygon {
	return newPolygonFromRings(lnrRings(coordinates)...)
}

//New Polygon from rings
func newPolygonFromRings(rings ...*LinearRing) *Polygon {
	var holes []*LinearRing
	if len(rings) > 1 {
		holes = rings[1:]
	}
	return &Polygon{Shell: rings[0], Holes: holes}
}

//create a new linestring from wkt string
//empty wkt are not allowed
func NewPolygonFromWKT(wkt string) *Polygon {
	return NewPolygon(ReadWKT(wkt, GeoTypePolygon).ToCoordinates()...)
}

//get geometry type
func (self *Polygon) Type() GeoType {
	return GeoType(GeoTypePolygon)
}

//get geometry interface
func (self *Polygon) Geometry() Geometry {
	return self
}

//ConvexHull computes slice of vertices as points forming convex hull
func (self *Polygon) ConvexHull() *Polygon {
	return NewPolygon(ConvexHull(self.Shell.Coordinates))
}

//As line strings
func (self *Polygon) AsLinearRings() []*LinearRing {
	var rings = make([]*LinearRing, 0, len(self.Holes)+1)
	rings = append(rings, self.Shell)
	for i := range self.Holes {
		rings = append(rings, self.Holes[i])
	}
	return rings
}

//polygon lnrRings
func lnrRings(coords []Coords) []*LinearRing {
	var rings = make([]*LinearRing, 0, len(coords))
	for i := range coords {
		rings = append(rings, NewLinearRing(coords[i]))
	}
	return rings
}
