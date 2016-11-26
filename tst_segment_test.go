package geom

import (
    . "github.com/franela/goblin"
    . "simplex/util/math"
    "math"
    "testing"
)

func TestSegment(t *testing.T) {
    g := Goblin(t)

    g.Describe("Segment", func() {
        g.It("should test segment intersection", func() {
            a := NewPointXY(0, 0)
            b := NewPointXY(-3, 4)
            c := NewPointXY(1.5, -2)
            d := NewPointXY(-1.5, 2)
            e := NewPointXY(0.5, 3)
            //f := &Point{-2, -2}
            gk := &Point{-1.5, -2.5}
            h := &Point{0.484154648492778, -0.645539531323704}
            i := &Point{0.925118053504632, -1.233490738006176}
            k := &Point{2, 2}
            n := &Point{1, 5}

            seg_ab := NewSegment(a, b)
            seg_de := &Segment{d, e}

            seg_cd := &Segment{c, d}
            seg_gkh := &Segment{gk, h}
            seg_hi := &Segment{h, i}
            seg_ak := &Segment{a, k}
            seg_kn := &Segment{k, n}

            pts, ok := seg_ab.Intersection(seg_de, false)
            g.Assert(ok).Equal(true)
            g.Assert(pts[0]).Equal(&Point{-1.5, 2})

            pts, ok = seg_ab.Intersection(seg_cd, false)
            g.Assert(ok).Equal(true)
            g.Assert(pts[0]).Equal(&Point{-1.5, 2})
            g.Assert(pts[1]).Equal(&Point{0.0, 0.0})

            pts, ok = seg_gkh.Intersection(seg_cd, false)
            g.Assert(ok).Equal(true)
            g.Assert(len(pts)).Equal(1) //at h

            pts, ok = seg_hi.Intersection(seg_cd, false)
            g.Assert(ok).Equal(true)
            g.Assert(len(pts)).Equal(2) //at h, i

            pts, ok = seg_hi.Intersection(seg_ab, false)
            g.Assert(seg_hi.Intersects(seg_ab, false)).Equal(ok)
            g.Assert(ok).Equal(false)
            g.Assert(len(pts)).Equal(0) //empty

            pts, ok = seg_ak.Intersection(seg_kn, false)
            g.Assert(seg_ak.Intersects(seg_kn, false)).Equal(ok)
            g.Assert(ok).Equal(true)
            g.Assert(len(pts)).Equal(1)//at k
            g.Assert(pts[0]).Equal(k) //k
        })
    })

}

func TestSegDist(t *testing.T) {
    g := Goblin(t)
    g.Describe("SegSeg and SegToPoint", func() {
        g.It("should test segment distance to point and seg", func() {
            var a = NewPointXY(-0.8, -2.6)
            var b = NewPointXY(-1, 1)
            var c = NewPointXY(-2, 3)
            var d = NewPointXY(7, -3)
            var e = NewPointXY(1.6, 0.6)
            var f = NewPointXY(-8, 4)
            var gi = NewPointXY(10, -8)
            var j =  NewPointXY(-3.5, 4)
            var k =  NewPointXY(-5, 5)
            var l =  NewPointXY(8.5, -4)
            var m =  NewPointXY(10, -5)
            var t =  NewPointXY(1, 6)
            var u =  NewPointXY(6, 4)


            var n =  NewPointXY(1, 3)
            var o =  NewPointXY(6, 5)
            var expects = Round(1.1094003924504583, 12);

            seg_ab := NewSegment(a, b)
            seg_ba := NewSegment(b, a)
            seg_cd := NewSegment(c, d)
            seg_dc := NewSegment(d, c)
            seg_dd := NewSegment(d, d)
            seg_ff := NewSegment(f, f)
            seg_ef := NewSegment(e, f)
            seg_fg := NewSegment(f, gi)
            seg_jk := NewSegment(j, k)
            seg_jj := NewSegment(j, j)
            seg_lm := NewSegment(l, m)
            seg_ll := NewSegment(l, l)
            seg_no := NewSegment(n, o)
            seg_tu := NewSegment(t, u)

            g.Assert(Round(seg_ab.Distance(seg_ab), 12)).Equal(0.0)
            g.Assert(Round(seg_ab.Distance(seg_cd), 12)).Equal(expects)
            g.Assert(Round(seg_ab.Distance(seg_dc), 12)).Equal(expects)
            g.Assert(Round(seg_ba.Distance(seg_cd), 12)).Equal(expects)
            g.Assert(Round(seg_cd.Distance(seg_ab), 12)).Equal(expects)

            g.Assert(Round(seg_dc.Distance(seg_ef), 12)).Equal(0.0)
            g.Assert(seg_dd.Distance(seg_ff)).Equal(d.Distance(f))
            g.Assert(Round(seg_dc.Distance(seg_fg), 12)).Equal(
                Round(2.496150883013531, 12),
            )

            g.Assert(Round(seg_dc.Distance(seg_lm), 12)).Equal(
                Round(d.Distance(l), 12),
            )

            g.Assert(Round(seg_cd.Distance(seg_lm), 12)).Equal(
                Round(d.Distance(l), 12),
            )
            g.Assert(Round(seg_dc.Distance(seg_ll), 12)).Equal(
                Round(d.Distance(l), 12),
            )
            g.Assert(Round(seg_cd.Distance(seg_ll), 12)).Equal(
                Round(d.Distance(l), 12),
            )

            g.Assert(Round(seg_dc.Distance(seg_jk), 12)).Equal(
                Round(c.Distance(j), 12),
            )
            g.Assert(Round(seg_dc.Distance(seg_jj), 12)).Equal(
                Round(c.Distance(j), 12),
            )
            g.Assert(Round(seg_jj.Distance(seg_dc), 12)).Equal(
                Round(c.Distance(j), 12),
            )


            g.Assert(Round(seg_ab.Distance(seg_no), 12)).Equal(
                Round(b.Distance(n), 12),
            )
            g.Assert(Round( seg_no.Distance(seg_ab), 12)).Equal(
                Round(n.Distance(b), 12),
            )
            //no intersects tu
            g.Assert(Round( seg_no.Distance(seg_tu), 12)).Equal(0.0)

            a = &Point{16.82295, 10.44635}
            b = &Point{28.99656, 15.76452}
            on_ab := &Point{25.32, 14.16}

            tpoints := []*Point{
                {30., 0.},
                {15.78786, 25.26468},
                {-2.61504, -3.09018},
                {28.85125, 27.81773},
                a, b, on_ab,
            }

            t_dists := []float64{14.85, 13.99, 23.69, 12.05, 0.00, 0.00, 0.00}
            tvect := NewSegment(a, b)
            dists := make([]float64, len(tpoints))

            for i, tp := range tpoints {
                dists[i] = tvect.segpt_mindist(tp)
            }

            pt1_out := NewPointFromWKT("POINT ( 49.8322373906287 49.1670033843562 )")
            pt2_out := NewPointFromWKT("POINT (  26.70508112717612 29.46609249326697 )")
            pnt3_in := NewPointFromWKT("POINT ( 27.439276564111122 38.76590136111034 )")
            poly := NewPolygonFromWKT("POLYGON (( 35 10, 45 45, 15 40, 10 20, 35 10 ), ( 20 30, 35 35, 30 20, 20 30 ))")
            var pt_online = NewPointXY(45.00000 , 45.000000000000000000000000001)
            ln := poly.Shell.AsLinear()[0]

            g.Assert(Round(ln.Distance(poly), 12)).Equal(Round(0, 12))
            g.Assert(Round(ln.Distance(pt_online), 12)).Equal(Round(0, 12))
            g.Assert(Round(pt1_out.Distance(ln), 12)).Equal(Round(6.380786425247758, 12))
            g.Assert(Round(ln.Distance(pt1_out), 12)).Equal(Round(6.380786425247758, 12))
            g.Assert(Round(pt1_out.Distance(poly), 12)).Equal(Round(6.380786425247758, 12))
            g.Assert(Round(poly.Distance(pt1_out), 12)).Equal(Round(6.380786425247758, 12))
            g.Assert(Round(pt2_out.Distance(poly), 12)).Equal(Round(2.626841960149983, 12))
            g.Assert(Round(poly.Distance(pt2_out), 12)).Equal(Round(2.626841960149983, 12))
            g.Assert(Round(pnt3_in.Distance(poly), 12)).Equal(Round(0.0, 12))
            g.Assert(Round(poly.Distance(pnt3_in), 12)).Equal(Round(0.0, 12))

            var null_pt *Point
            var null_poly *Polygon
            var null_ln *LineString
            g.Assert(math.IsNaN(poly.Distance(null_pt))).IsTrue()
            g.Assert(math.IsNaN(pt2_out.Distance(null_pt))).IsTrue()
            g.Assert(math.IsNaN(ln.Distance(null_pt))).IsTrue()
            g.Assert(math.IsNaN(pt2_out.Distance(null_poly))).IsTrue()
            g.Assert(math.IsNaN(poly.Distance(null_poly))).IsTrue()
            g.Assert(math.IsNaN(poly.Distance(null_ln))).IsTrue()


            var seg_aa = NewSegment(a, a)
            g.Assert(seg_aa.segpt_mindist(a)).Equal(0.0)
            g.Assert(a.SideOf(a, b).IsOn()).IsTrue()
            g.Assert(b.SideOf(a, b).IsOn()).IsTrue()

            seg_ab = NewSegment(a, b)
            g.Assert(seg_ab.SideOf(a).IsOn()).IsTrue()
            g.Assert(seg_ab.SideOf(b).IsOn()).IsTrue()

            for i, _ := range tpoints {
                g.Assert(Round(dists[i], 2)).Equal(Round(t_dists[i], 2))
            }
        });

    })

}

