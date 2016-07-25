package geom

import (
    . "github.com/franela/goblin"
    "testing"
)

func TestLineStringSelf(t *testing.T) {
    g := Goblin(t)
    wkt0 := "LINESTRING(520.3891360357894 542.3912033070129,506.3024618690045 551.4232473315985,499.8456492240652 555.3948968460392,492.961552805167 552.5004635914114,489.3155900796462 547.0315195031302,494.7910190818659 540.6453203655232,503.2430235819369 542.0539877822016,506.3024618690045 551.4232473315985,505.72509579166825 560.3502151427206,505.2252456091915 568.0786679640912,509.1982380315184 573.1744625927278,510.22036282943196 578.5066706448671,506.9538876603224 582.6378010057998,500.5170101211947 582.253509809434,492.2547493993293 573.7991034893856,486.1060902574759 569.7640459275444,481.13065522694603 565.0658697854723,477.55561113833613 565.056478772063,477.07661476306924 570.8172868053748,478.90063032561653 576.393069064855,488.0275462393051 588.113950554013,494.11875860924937 590.8324257774667,503.20704849575554 590.8039889285739,513.3907651994501 588.2100233531045,519.7227675184544 584.0205966373184,528.1782382221293 579.2651445083028,528.2675859252843 570.3930616942043,532.7013978168333 560.3489116165815,531.7413381584337 549.6831287580544)"
    wkt1 := "LINESTRING ( -876.7342911153115 -1179.1875803402643, -1379.2842911153114 -1560.0675803402642, -1495.6642911153112 -983.4575803402644, -1003.6942911153114 -666.0575803402645, -1035.4342911153115 -909.3975803402644, 125.35053019388499 -1134.0185052232773, -210.1942911153118 -1390.7875803402642, 398.1557088846879 -1401.367580340264, -130.84429111531185 -882.9475803402644, -670.4242911153116 -819.4675803402644, -62.07429111531189 -671.3475803402645, 768.4557088846877 -1242.6675803402643, 202.425708884688 -47.12758034026476, 1371.5157088846875 -454.4575803402646, 1419.1257088846874 -1126.2875803402642, 2228.4957088846872 -734.8275803402645, 1789.4257088846873 -1729.347580340264, 995.9257088846877 -1454.2675803402642, 667.9457088846877 -2237.1875803402636, -876.7342911153115 -1179.1875803402643 )"
    wkt2 := "LINESTRING ( -1003.6942911153114 -36.54758034026476, 943.0257088846877 1370.5924196597346, 1457.889355803078 807.0600697988099, 2006.3157088846872 206.79241965973512, 2362.7803242693026 741.4893427366578, 2624.7844100209995 665.4332995584509, 2794.525708884687 333.7524196597351, 3074.895708884687 534.772419659735, 3006.125708884687 735.792419659735, 3085.475708884687 968.5524196597348, 2558.9172473462254 1035.6947273520427, 2895.0357088846868 1539.8724196597345, 3688.5357088846863 598.252419659735, 4196.375708884686 1555.7424196597344, 2096.2457088846872 995.0024196597348, 2773.3657088846867 -295.75758034026467, 1889.9357088846873 -533.8075803402645, 1069.9857088846877 -142.34758034026473, 874.2557088846877 698.7624196597349, -72.65429111531188 1169.5724196597348 )"

    g.Describe("Linestring-Self Intersection", func() {
        g.It("test self intersection", func() {
            ln0 := NewLineStringFromWKT(wkt0)
            g.Assert(ln0.IsSimple()).IsFalse()
            g.Assert(ln0.IsComplex()).IsTrue()
            g.Assert(len(ln0.SelfIntersection())).Equal(1)

            ln1 := NewLineStringFromWKT(wkt1)
            g.Assert(ln1.IsSimple()).IsFalse()
            g.Assert(ln1.IsComplex()).IsTrue()
            g.Assert(len(ln1.SelfIntersection())).Equal(1)

            ln2 := NewLineStringFromWKT(wkt2)
            g.Assert(ln2.IsSimple()).IsFalse()
            g.Assert(ln2.IsComplex()).IsTrue()
            g.Assert(len(ln2.SelfIntersection())).Equal(4)
        })
    })

}