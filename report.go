package main

import (
	"fmt"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// Report ...
func (r *RdbReport) Report() {
	r.reportMem()
	r.reportBid()
}

func (r *RdbReport) reportBid() {
	var total uint64
	for _, v := range r.bidSizeMap {
		total += v
	}

	for bid, l := range r.bidSizeMap {
		fmt.Println(bid, ": ", l, float64(l)/float64(total))
	}

	for bid, s := range r.bidNoExpiryMap {
		fmt.Println("no expiry: ", bid, ": ", s)
	}
}

func (r *RdbReport) reportMem() {
	total := r.keyLen + r.valueLen
	//128  4K  32K   128K  512K  2M    32M   128M  512M  INF
	//0-6 7-11 12-14 15-16 17-18 19-20 21-24 25-26 27-28

	gString := plotter.Values{
		float64(r.mString.accumSum(0, 6)) / float64(total) * 100,
		float64(r.mString.accumSum(7, 11)) / float64(total) * 100,
		float64(r.mString.accumSum(12, 14)) / float64(total) * 100,
		float64(r.mString.accumSum(15, 16)) / float64(total) * 100,
		float64(r.mString.accumSum(17, 18)) / float64(total) * 100,
		float64(r.mString.accumSum(19, 20)) / float64(total) * 100,
		float64(r.mString.accumSum(21, 24)) / float64(total) * 100,
		float64(r.mString.accumSum(25, 26)) / float64(total) * 100,
		float64(r.mString.accumSum(27, 28)) / float64(total) * 100,
		float64(r.mString.accumSum(29, 59)) / float64(total) * 100,
	}

	gHash := plotter.Values{
		float64(r.mHash.accumSum(0, 6)) / float64(total) * 100,
		float64(r.mHash.accumSum(7, 11)) / float64(total) * 100,
		float64(r.mHash.accumSum(12, 14)) / float64(total) * 100,
		float64(r.mHash.accumSum(15, 16)) / float64(total) * 100,
		float64(r.mHash.accumSum(17, 18)) / float64(total) * 100,
		float64(r.mHash.accumSum(19, 20)) / float64(total) * 100,
		float64(r.mHash.accumSum(21, 24)) / float64(total) * 100,
		float64(r.mHash.accumSum(25, 26)) / float64(total) * 100,
		float64(r.mHash.accumSum(27, 28)) / float64(total) * 100,
		float64(r.mHash.accumSum(29, 59)) / float64(total) * 100,
	}

	gList := plotter.Values{
		float64(r.mList.accumSum(0, 6)) / float64(total) * 100,
		float64(r.mList.accumSum(7, 11)) / float64(total) * 100,
		float64(r.mList.accumSum(12, 14)) / float64(total) * 100,
		float64(r.mList.accumSum(15, 16)) / float64(total) * 100,
		float64(r.mList.accumSum(17, 18)) / float64(total) * 100,
		float64(r.mList.accumSum(19, 20)) / float64(total) * 100,
		float64(r.mList.accumSum(21, 24)) / float64(total) * 100,
		float64(r.mList.accumSum(25, 26)) / float64(total) * 100,
		float64(r.mList.accumSum(27, 28)) / float64(total) * 100,
		float64(r.mList.accumSum(29, 59)) / float64(total) * 100,
	}

	gSet := plotter.Values{
		float64(r.mSet.accumSum(0, 6)) / float64(total) * 100,
		float64(r.mSet.accumSum(7, 11)) / float64(total) * 100,
		float64(r.mSet.accumSum(12, 14)) / float64(total) * 100,
		float64(r.mSet.accumSum(15, 16)) / float64(total) * 100,
		float64(r.mSet.accumSum(17, 18)) / float64(total) * 100,
		float64(r.mSet.accumSum(19, 20)) / float64(total) * 100,
		float64(r.mSet.accumSum(21, 24)) / float64(total) * 100,
		float64(r.mSet.accumSum(25, 26)) / float64(total) * 100,
		float64(r.mSet.accumSum(27, 28)) / float64(total) * 100,
		float64(r.mSet.accumSum(29, 59)) / float64(total) * 100,
	}

	gZset := plotter.Values{
		float64(r.mZset.accumSum(0, 6)) / float64(total) * 100,
		float64(r.mZset.accumSum(7, 11)) / float64(total) * 100,
		float64(r.mZset.accumSum(12, 14)) / float64(total) * 100,
		float64(r.mZset.accumSum(15, 16)) / float64(total) * 100,
		float64(r.mZset.accumSum(17, 18)) / float64(total) * 100,
		float64(r.mZset.accumSum(19, 20)) / float64(total) * 100,
		float64(r.mZset.accumSum(21, 24)) / float64(total) * 100,
		float64(r.mZset.accumSum(25, 26)) / float64(total) * 100,
		float64(r.mZset.accumSum(27, 28)) / float64(total) * 100,
		float64(r.mZset.accumSum(29, 59)) / float64(total) * 100,
	}

	p, _ := plot.New()
	p.Title.Text = fmt.Sprintf("value mem/total mem: %.2f%%", float64(r.valueLen)/float64(total)*100)
	p.Y.Label.Text = "mem/total * 100%"

	w := vg.Points(10)

	barsString, _ := plotter.NewBarChart(gString, w)
	barsString.LineStyle.Width = vg.Length(0)
	barsString.Color = plotutil.Color(0)
	barsString.Offset = -w * 2

	barsHash, _ := plotter.NewBarChart(gHash, w)
	barsHash.LineStyle.Width = vg.Length(0)
	barsHash.Color = plotutil.Color(1)
	barsHash.Offset = -w

	barsList, _ := plotter.NewBarChart(gList, w)
	barsList.LineStyle.Width = vg.Length(0)
	barsList.Color = plotutil.Color(2)

	barsSet, _ := plotter.NewBarChart(gSet, w)
	barsSet.LineStyle.Width = vg.Length(0)
	barsSet.Color = plotutil.Color(3)
	barsSet.Offset = w

	barsZset, _ := plotter.NewBarChart(gZset, w)
	barsZset.LineStyle.Width = vg.Length(0)
	barsZset.Color = plotutil.Color(4)
	barsZset.Offset = w * 2

	p.Add(barsString, barsHash, barsList, barsSet, barsZset)
	p.Legend.Add("string", barsString)
	p.Legend.Add("hash", barsHash)
	p.Legend.Add("list", barsList)
	p.Legend.Add("set", barsSet)
	p.Legend.Add("zset", barsZset)
	p.Legend.Top = true

	//128  4K  32K   128K  512K  2M    32M   128M  512M  INF
	p.NominalX("1 ~ 128", "128 ~ 4K", "4K ~ 32K", "32K ~ 128K", "128K ~ 512K", "512K ~ 2M", "2M ~ 32M", "32M ~ 128M", "128M ~ 512M", "512M ~ INF")
	if err := p.Save(12*vg.Inch, 7*vg.Inch, "barchart.png"); err != nil {
		panic(err)
	}
}
