package lib

import (
	"cake-syrup-pools-lover/util"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"os"
	"path/filepath"
)

// @title 生成日收益pieChart绘制所需数据结构
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

// @title 生成糖浆池总质押cake数量pieChart绘制所需数据结构
func GenerateStakedCakePieItems(SyrupPools *SyrupPools) []opts.PieData {
	items := make([]opts.PieData, 0)
	for index, _ := range SyrupPools.SyPools {
		sPool := SyrupPools.SyPools[index]
		items = append(items, opts.PieData{
			Name:  sPool.Token.Symbol,
			Value: util.BigFloat4Decimal(sPool.StakedCake)})
	}
	return items
}

// @title 生成饼状图的html文件
func CreateRosePieChart(SyrupPools *SyrupPools) {
	// create a new pie instance
	pie := charts.NewPie()
	nowStr := util.GetNowTimeStr()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(
			opts.Title{
				Title:    "Syrup Pools Daily Earn Per100Cake(USD)",
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

// @title 生成各糖浆池质押cake总数量饼状图的html文件
func CreateStakedPieChart(SyrupPools *SyrupPools) {
	// create a new pie instance
	pie := charts.NewPie()
	nowStr := util.GetNowTimeStr()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(
			opts.Title{
				Title:    "SyrupPools Total Staked Cake Amount",
				Subtitle: nowStr,
			},
		),
	)
	pie.SetSeriesOptions()
	pie.AddSeries("pie", GenerateStakedCakePieItems(SyrupPools)).
		SetSeriesOptions(charts.WithLabelOpts(
			opts.Label{
				Show:      true,
				Formatter: "{b}: {c}",
			}),
		)

	htmlFileName := "/public/stakedPie.html"
	AllConfig := GetConfig()
	htmlFilePath := filepath.Join(AllConfig.ProjectFolder, htmlFileName)
	f, _ := os.Create(htmlFilePath)
	defer f.Close()
	_ = pie.Render(f)
}
