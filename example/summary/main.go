package main

import (
	"fmt"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

var (
	//128  4K  32K   128K  512K  2M    32M   128M  512M  INF
	//0-6 7-11 12-14 15-16 17-18 19-20 21-24 25-26 27-28

	gString = plotter.Values{
		6.435880829,
		1.694300518,
		0.844343696,
		0.438039724,
		0,
		0,
		0,
		0,
		0,
		0,
	}

	gHash = plotter.Values{
		1.115716753,
		5.328151986,
		2.512953368,
		1.794905009,
		0.860103627,
		0.331390328,
		0,
		0,
		0,
		0,
	}

	gList = plotter.Values{
		0.016407599,
		0.243091537,
		0.064766839,
		0.046632124,
		0.466321244,
		0,
		0,
		0,
		0.035621762,
		0,
	}

	gSet = plotter.Values{
		0.142918826,
		0.681994819,
		0,
		0,
		0,
		0.002374784,
		0.018998273,
		0.130613126,
		0.04037133,
		0,
	}

	gZset = plotter.Values{
		1.804835924,
		5.764464594,
		15.81800518,
		11.31692573,
		14.19149396,
		6.91126943,
		0.970207254,
		1.255613126,
		3.062823834,
		0.55224525,
	}
)

func main() {
	reportSummary()
	reportType()
	reportSize()
}

func reportSize() {
	var gSize plotter.Values
	for i := range gString {
		gSize = append(gSize, gString[i]+gHash[i]+gList[i]+gSet[i]+gZset[i])
	}

	p, _ := plot.New()
	p.Title.Text = fmt.Sprintf("value mem/total mem: %.2f%%", 15.10621762)
	p.Y.Label.Text = "mem/total * 100%"

	w := vg.Points(20)

	barSize, _ := plotter.NewBarChart(gSize, w)
	barSize.LineStyle.Width = vg.Length(0)
	barSize.Color = plotutil.Color(6)

	p.Add(barSize)

	//128  4K  32K   128K  512K  2M    32M   128M  512M  INF
	p.NominalX("1 ~ 128", "128 ~ 4K", "4K ~ 32K", "32K ~ 128K", "128K ~ 512K", "512K ~ 2M", "2M ~ 32M", "32M ~ 128M", "128M ~ 512M", "512M ~ INF")
	if err := p.Save(12*vg.Inch, 7*vg.Inch, "size.png"); err != nil {
		panic(err)
	}
}

func reportType() {
	var tstring, thash, tlist, tset, tzset float64
	for _, v := range gString {
		tstring += v
	}
	for _, v := range gHash {
		thash += v
	}
	for _, v := range gList {
		tlist += v
	}
	for _, v := range gSet {
		tset += v
	}
	for _, v := range gZset {
		tzset += v
	}

	gSummary := plotter.Values{tstring, thash, tlist, tset, tzset}

	p, _ := plot.New()
	p.Title.Text = fmt.Sprintf("value mem/total mem: %.2f%%", 15.10621762)
	p.Y.Label.Text = "mem/total * 100%"

	w := vg.Points(30)

	bargSummary, _ := plotter.NewBarChart(gSummary, w)
	bargSummary.LineStyle.Width = vg.Length(0)
	bargSummary.Color = plotutil.Color(5)

	p.Add(bargSummary)

	p.NominalX("string", "hash", "list", "set", "zset")
	if err := p.Save(12*vg.Inch, 7*vg.Inch, "type.png"); err != nil {
		panic(err)
	}
}

func reportSummary() {
	p, _ := plot.New()
	p.Title.Text = fmt.Sprintf("value mem/total mem: %.2f%%", 15.10621762)
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
	if err := p.Save(12*vg.Inch, 7*vg.Inch, "summary.png"); err != nil {
		panic(err)
	}
}
