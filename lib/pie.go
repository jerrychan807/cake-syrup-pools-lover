package lib

import (
	"cake-syrup-pools-lover/util"
	"os"
	"path/filepath"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// @title 生成pieChart绘制所需数据结构
func GenerateSyrupPoolPieItems(SyrupPools *SyrupPools) []opts.PieData {
	//pools := []string{"SFUND", "PSTAKE", "PEEL", "SHELL"}
	//values := []float64{0.79, 0.21, 0.18, 0.19}
	//items := make([]opts.PieData, 0)
	//for i := 0; i < 4; i++ {
	//	items = append(items, opts.PieData{
	//		Name:  pools[i],
	//		Value: values[i]})
	//}
	items := make([]opts.PieData, 0)
	for index, _ := range SyrupPools.SyPools {
		sPool := SyrupPools.SyPools[index]
		items = append(items, opts.PieData{
			Name:  sPool.Token.Symbol,
			Value: util.BigFloat4Decimal(sPool.HundredCakeDailyEarn.String())})
	}
	return items
}

// @title 生成饼状图的html文件
func CreatePieChart(SyrupPools *SyrupPools) {
	// create a new pie instance
	pie := charts.NewPie()
	nowStr := time.Now().Format("2006-01-02 15:04:05") //获取当前时间
	pie.SetGlobalOptions(
		charts.WithTitleOpts(
			opts.Title{
				Title:    "Syrup Pools Daily Earn(USD)",
				Subtitle: nowStr,
			},
		),
	)
	pie.SetSeriesOptions()
	pie.AddSeries("Monthly revenue",
		GenerateSyrupPoolPieItems(SyrupPools)).
		SetSeriesOptions(
			charts.WithPieChartOpts(opts.PieChart{
				Radius:   []string{"40%", "75%"},
				RoseType: "area",
			}),
			charts.WithLabelOpts(
				opts.Label{
					Show:      true,
					Formatter: "{b}: {c}",
				},
			),
		)

	htmlFileName := "/public/pie.html"
	AllConfig := GetConfig()
	htmlFilePath := filepath.Join(AllConfig.ProjectFolder, htmlFileName)
	f, _ := os.Create(htmlFilePath)
	defer f.Close()
	_ = pie.Render(f)
}
