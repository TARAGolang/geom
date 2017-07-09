package main

import (
	"simplex/geom"
	"fmt"
)

func main() {
	// ln := "LINESTRING ( 810 540, 790 570, 800 580, 820 580, 860 570, 880 600, 870 610, 850 610, 800 610, 810 650, 890 640, 900 640, 920 600, 930 580, 930 540, 920 500, 880 490, 860 520, 810 510, 750 520, 780 460, 730 410, 830 440, 890 410, 940 450, 970 500, 1040 510, 1050 570, 1080 620, 1040 660, 1020 720, 950 720, 840 680, 760 690, 690 720, 710 640, 630 620 )"
	// ln := "LINESTRING ( 810 540, 790 570, 800 580, 820 580, 860 570, 880 600, 870 610, 850 610, 800 610, 810 650, 890 640, 900 640, 920 600, 930 580, 930 540, 920 500, 880 490, 860 520, 810 510, 750 520, 780 460, 730 410, 710 480, 680 400, 810 360, 970 500, 1040 510, 1050 570, 1080 620, 1040 660, 1020 720, 950 720, 840 680, 760 690, 690 720, 710 640, 630 620 )"
	// ln := "LINESTRING ( 810 540, 790 570, 800 580, 820 580, 860 570, 880 600, 870 610, 850 610, 800 610, 810 650, 890 640, 900 640, 920 600, 930 580, 930 540, 920 500, 880 490, 860 520, 810 510, 750 520, 780 460, 730 410, 710 480, 770 430, 810 430, 970 500, 1040 510, 1050 570, 1080 620, 1040 660, 1020 720, 950 720, 840 680, 760 690, 690 720, 710 640, 630 620 )"
	// ln := "LINESTRING (330.5770802442433 581.2490753170464, 289.3831572774428 587.7108671549759, 256.2664741080541 582.8645232765288, 197.441310580554 568.7784526288414, 194.9264255896295 552.0343728591217, 202.11368678741684 562.5981327919068, 211.82048369391282 548.4791554733671, 216.23266410595645 527.5948348563605, 218.58582699237974 516.711456506653, 185.64154658245394 504.3573513529308, 176.52304039756376 493.17982764242026, 172.1108599855201 481.4140132103039, 172.69915070712594 461.706274036509, 157.69773730617757 464.3535822837352, 137.9899981323827 464.941873005341)"
	// ln := "LINESTRING ( 330.5770802442433 581.2490753170464, 289.3831572774428 587.7108671549759, 256.2664741080541 582.8645232765288, 197.441310580554 568.7784526288414, 194.9264255896295 552.0343728591217, 202.11368678741684 562.5981327919068, 211.74199479333123 560.2115030178386, 227.0827680621383 567.7253511494993, 236.3364918395976 564.4237400821951, 249.31123545163425 568.9776591714426, 241.34572392737135 584.7737454387759, 235.22277020477063 586.8230484841366, 227.70892207311002 596.8415126596841, 222.69968998533628 614.3738249668921, 180 620 )"
	// ln := "LINESTRING ( 330.5770802442433 581.2490753170464, 289.3831572774428 587.7108671549759, 256.2664741080541 582.8645232765288, 197.441310580554 568.7784526288414, 194.9264255896295 552.0343728591217, 203.77648326906825 558.3187384752208, 215.0472554665592 551.7441213600177, 227.0827680621383 567.7253511494993, 236.3364918395976 564.4237400821951, 237.11918435331225 572.5637422248274, 241.34572392737135 584.7737454387759, 209.2553308650708 583.2083604113466, 207.2203303294127 590.2525930347784, 204.40263728003995 601.9929807404982, 189.2184025139758 610.2895213858734 )"
	ln := "LINESTRING ( 230.0195128864249 579.5695574923628, 217.3543939499459 574.8349335908753, 210.8442860854006 568.916653714016, 197.441310580554 568.7784526288414, 194.9264255896295 552.0343728591217, 203.77648326906825 558.3187384752208, 215.0472554665592 551.7441213600177, 227.0827680621383 567.7253511494993, 236.3364918395976 564.4237400821951, 237.11918435331225 572.5637422248274, 241.34572392737135 584.7737454387759, 209.2553308650708 583.2083604113466, 207.2203303294127 590.2525930347784, 204.40263728003995 601.9929807404982, 189.2184025139758 610.2895213858734 )"
	wkt := "POINT ( 130 540 )"
	//wkt := "POINT ( 270 540 )"

	pt := geom.NewPointFromWKT(wkt)
	g := geom.NewLineStringFromWKT(ln)
	coords := g.Coordinates()
	i, j := geom.TangentPointToPoly(pt, coords)
	fmt.Println(coords[i].WKT())
	fmt.Println(coords[j].WKT())
	fmt.Println(g)

}
