package geom

import (
    . "simplex/geom/mbr"
    "simplex/struct/sset"
)

//segments in range
//xor - altenate segments if nothing is in range of box
func (self *LineString) segs_inrange(seglist []*Segment,
box *MBR, i, j int, extend, xor bool) []*Segment {

    //extend or refresh list
    if !extend {
        seglist = make([]*Segment, 0)
    }

    altsegs := make([]*Segment, 0)//, bool, seg
    for ; i < j; i++ {
        inters := box.IntersectsBounds(
            self.coordinates[i][:],
            self.coordinates[i + 1][:],
        )
        var seg = &Segment{
            self.coordinates[i],
            self.coordinates[i + 1],
        }
        if inters {
            seglist = append(seglist, seg)
        } else {
            altsegs = append(altsegs, seg)
        }
    }
    if xor && len(seglist) == 0 {
        seglist = append(seglist, altsegs...)
    }
    return seglist
}

//Segment - Segment intersection of slice of arrays
func (self *LineString) segseg_intersection(segsa, segsb []*Segment, ptset *sset.SSet, extend bool) {
    if !extend {
        ptset.Empty()
    }

    for a := 0; a < len(segsa); a++ {
        for b := 0; b < len(segsb); b++ {
            if coord, ok := segsa[a].Intersection(
                segsb[b], false); ok {
                for _, pt := range coord {
                    ptset.Add(pt)
                }
            }
        }
    }
}



