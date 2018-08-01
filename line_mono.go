package geom

import (
	"github.com/intdxdt/mbr"
	"github.com/intdxdt/geom/mono"
	"github.com/intdxdt/math"
)

//build xymonotone chain, perimeter length,
//monotone build starts from i and ends at j, designed for
//appending new points to the end of line
func (self *LineString) processChains() *LineString {
	var dx, dy float64
	var cur_x, cur_y, prev_x, prev_y int
	var i, j = 0, self.Coordinates.Len()-1
	var monoLimit = int(math.Log2(float64(j+1) + 1.0))

	//init chains
	self.chains = make([]mono.MBR, 0, 2*monoLimit)

	prev_x, prev_y = null, null
	var a, b *Point
	a = self.Coordinates.Pt(i)
	var box = mbr.MBR{a[X], a[Y], a[X], a[Y]}

	self.bbox = mono.CreateMonoMBR(box)
	var mbox = mono.CreateMonoMBR(box)

	self.xyMonobox(&mbox, i, i)
	self.chains = append(self.chains, mbox)

	var mono_size = 0
	var m_index = len(self.chains) - 1

	for i = i + 1; i <= j; i += 1 {
		a, b = self.Coordinates.Pt(i-1), self.Coordinates.Pt(i)
		dx = b[X] - a[X]
		dy = b[Y] - a[Y]

		cur_x = xySign(dx)
		cur_y = xySign(dy)

		if prev_x == null {
			prev_x = cur_x
		}

		if prev_y == null {
			prev_y = cur_y
		}

		//((cur_x + prev_x > 0) && (prev_y + cur_y > 0))
		mono_size += 1
		if prev_x == cur_x && prev_y == cur_y && mono_size <= monoLimit {
			self.xyMonobox(&self.chains[m_index], i, null)
		} else {
			mono_size = 1
			prev_x, prev_y = cur_x, cur_y
			a, b = self.Coordinates.Pt(i-1), self.Coordinates.Pt(i)
			mbox = mono.CreateMonoMBR(
				mbr.CreateMBR(a[X], a[Y], b[X], b[Y]),
			)
			self.xyMonobox(&mbox, i-1, i)
			self.chains = append(self.chains, mbox)
			m_index = len(self.chains) - 1
		}
	}
	return self
}

//compute bbox of x or y mono chain
func (self *LineString) xyMonobox(mono *mono.MBR, i, j int) {
	if i != null {
		var pt = self.Coordinates.Pt(i)
		mono.ExpandIncludeXY(pt[X], pt[Y])
		if j == null {
			mono.J = i
		} else {
			mono.I, mono.J = i, j
		}

		self.bbox.ExpandIncludeMBR(&mono.MBR)
		if self.bbox.I == null {
			self.bbox.I, self.bbox.J = mono.I, mono.J
		} else {
			if mono.J > self.bbox.J {
				self.bbox.J = mono.J
			}
		}
	}
}

//find the sign of value -1, 0 , 1
func xySign(v float64) int {
	var i int
	if v > 0 {
		i = 1
	} else if v < 0 {
		i = -1
	}
	return i
}
