package lib

import (
	"cake-syrup-pools-lover/util"
	"fmt"
	"github.com/klarkxy/gohtml"
	"os"
	"path/filepath"
	"time"
)

// @title 生成糖浆池信息的表格html
func GenerateSyrupPoolTableHtml(SyrupPools *SyrupPools) {
	/*
		// 初始化html文件
		bootstrap := gohtml.NewHtml()
		bootstrap.Html().Lang("zh-CN")
		// Meta部分
		bootstrap.Meta().Charset("utf-8")
		bootstrap.Meta().Http_equiv("X-UA-Compatible").Content("IE=edge")
		bootstrap.Meta().Name("viewport").Content("width=device-width, initial-scale=1")
		// Css引入
		bootstrap.Link().Href("https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/css/bootstrap.min.css").Rel("stylesheet")
		// Js引入
		bootstrap.Script().Src("https://cdn.bootcss.com/jquery/3.4.1/jquery.min.js")
		bootstrap.Script().Src("https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/js/bootstrap.min.js")
		// Head
		bootstrap.Head().Title().Text("Bootstrap 101 Template")
		// Body
		bootstrap.Body().H1().Text("SyrupPools Info")

		table := bootstrap.Body().Tag("table").Class("table table-bordered") //.Border("1")
	*/
	// 设置时间
	nowStr := time.Now().Format("2006-01-02 15:04:05") //获取当前时间
	timeTag := gohtml.Tag("p").Text(fmt.Sprintf("Updated at: %s", nowStr))
	// 将糖浆池数组信息 塞进表格里
	table := gohtml.Tag("table").Border("1")
	// 设置表头
	title := table.Tag("tr")
	title.Tag("th").Text("SyrupPool ID")
	title.Tag("th").Text("RewardToken Name/Symbol")
	title.Tag("th").Text("RewardToken Price(USD)")
	//title.Tag("th").Text("RewardToken ContractAddr")
	title.Tag("th").Text("Daily/Weekly/Monthly/Yearly Profit(USD)(Per100Cake)")
	title.Tag("th").Text("SyrupPool TotalStaked Cake")
	title.Tag("th").Text("SyrupPool StartTime")
	title.Tag("th").Text("SyrupPool EndTime")
	// 设置中文表头
	//cntitle := table.Tag("tr")
	//cntitle.Tag("th").Text("糖浆池ID")
	//cntitle.Tag("th").Text("奖励代币 名字/符号")
	//cntitle.Tag("th").Text("奖励代币 价格(USD)")
	//cntitle.Tag("th").Text("奖励代币 合约地址")
	//cntitle.Tag("th").Text("日/周/月/年 收益(USD)(每100Cake)")
	//cntitle.Tag("th").Text("糖浆池总质押Cake数量")
	//cntitle.Tag("th").Text("糖浆池开始时间")
	//cntitle.Tag("th").Text("糖浆池结束时间")
	// 循环填充数据
	for index, _ := range SyrupPools.SyPools {
		sPool := SyrupPools.SyPools[index]
		tr := table.Tag("tr")
		tr.Tag("th").Text(sPool.SousId)
		tr.Tag("th").Text(fmt.Sprintf("%s/%s", sPool.Token.Name, sPool.Token.Symbol))
		tr.Tag("th").Text(util.BigFloat4Decimal(sPool.Token.Price.String()))
		//tr.Tag("th").Text(sPool.Token.ContractAddr)
		tr.Tag("th").Text(fmt.Sprintf("%s/%s/%s/%s", util.BigFloat4Decimal(sPool.HundredCakeDailyEarn.String()), util.BigFloat4Decimal(sPool.HundredCakeWeekEarn.String()), util.BigFloat4Decimal(sPool.HundredCakeMonthEarn.String()), util.BigFloat4Decimal(sPool.HundredCakeYearEarn.String())))
		tr.Tag("th").Text(fmt.Sprintf("%s", util.BigFloat4Decimal(sPool.StakedCake)))
		tr.Tag("th").Text(fmt.Sprintf("%s", sPool.StartTime))
		tr.Tag("th").Text(fmt.Sprintf("%s", sPool.EndTime))
	}

	// 写进html文件
	htmlFileName := "/public/table.html"
	AllConfig := GetConfig()
	htmlFilePath := filepath.Join(AllConfig.ProjectFolder, htmlFileName)
	f, _ := os.Create(htmlFilePath)
	defer f.Close()
	f.Write([]byte(timeTag.String() + table.String()))
}
