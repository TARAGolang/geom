package geom

import (
    "github.com/intdxdt/simplex/struct/rtree"
    "math"
)

type LineString struct {
    chains      []*MonoMBR
    coordinates []*Point
    monosize    int
    bucketsize  int
    index       *rtree.RTree
    bbox        *MonoMBR
    length      float64
}

//New LineString from a given coordinates {Array} [[x,y], ....[x,y]]
func NewLineString(coordinates []*Point) *LineString {
    if len(coordinates) < 2 {
        panic("a linestring must have at least 2 coordinate")
    }
    self := &LineString{}
    self.chains = make([]*MonoMBR, 0)

    //copy coordinates
    self.coordinates = CloneCoordinates(coordinates)

    //init
    self.monosize = int(math.Log2(float64(len(coordinates)) + 1.0))
    self.bucketsize = 9
    self.index = rtree.NewRTree(self.bucketsize)

    self.process_chains(0, 0)
    self.build_index()
    return self
}

//New line string from array
func NewLineStringFromArray(array [][2]float64) *LineString {
    return NewLineString(AsPointArray(array))
}

//create a new linestring from wkt string
//empty wkt are not allowed
func NewLineStringFromWKT(wkt_geom string) *LineString {
    return NewLineStringFromArray(
        ReadWKT(wkt_geom).ToArray()[0],
    )
}

//Point to LineString
func NewLineStringFromPoint(pt *Point) *LineString {
    return NewLineString([]*Point{pt.Clone(), pt.Clone()})
}

//get copy of chains of linestring
func (self  *LineString) MonoChains() []*MonoMBR {
    chains := make([]*MonoMBR, len(self.chains))
    for i := range self.chains {
        chains[i] = self.chains[i].Clone()
    }
    return chains
}
