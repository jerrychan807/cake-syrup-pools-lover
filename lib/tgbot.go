package lib

import (
	"cake-syrup-pools-lover/util"
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"path/filepath"
	"strconv"
	"time"
)

//tg机器人
func TgBotStartServer() {
	AllConfig := GetConfig()
	b, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		//URL: "http://195.129.111.17:8012",
		Token:  AllConfig.TgToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	// 订阅糖浆池信息通知
	b.Handle("/subscribe_syrup_pool", func(m *tb.Message) {
		b.Send(m.Sender, "Thank you for subscribing")
		SaveChatId(int(m.Sender.ID))
	})

	// 查询订阅人数
	b.Handle("/subscribe_nums", func(m *tb.Message) {
		uids := QueryAllChatId()
		msgStr := fmt.Sprintf("subscribes num: %d", len(uids))
		b.Send(m.Sender, msgStr)
		fmt.Println(uids)
	})

	// 查询糖浆池信息订阅是否成功
	b.Handle("/subscribe_result", func(m *tb.Message) {
		userId := QueryChatId(int(m.Sender.ID))
		userIdStr := strconv.Itoa(userId)
		msgStr := fmt.Sprintf("Subscribe Successfully, Your userId :%s", userIdStr)
		b.Send(m.Sender, msgStr)
	})

	// 查询糖浆池日收益的饼状图
	b.Handle("/syrup_pools_daily_earn_pie", func(m *tb.Message) {
		img := getImg("/download/pie.png")
		b.Send(m.Sender, img)
	})

	// 查询糖浆池所有信息的表格图
	b.Handle("/syrup_pools_table", func(m *tb.Message) {
		img := getImg("/download/table.png")
		b.Send(m.Sender, img)
	})

	// 查询糖浆池总质押cake的饼状图
	b.Handle("/syrup_pools_staked_cake_pie", func(m *tb.Message) {
		img := getImg("/download/stakedPie.png")
		b.Send(m.Sender, img)
	})

	// 查询糖浆池日产量Usd的饼状图
	b.Handle("/syrup_pools_daily_yield_pie", func(m *tb.Message) {
		img := getImg("/download/dailyYieldPie.png")
		b.Send(m.Sender, img)
	})

	// 查询最新糖浆池-完整信息
	b.Handle("/syrup_pools_full", func(m *tb.Message) {
		poolMsg := QuerySyrupPoolStr("pools")
		b.Send(m.Sender, poolMsg)
	})

	fmt.Println("[*] TgBot StartServer!!!")
	b.Start()
}

// @title 获取图片
// @param relativePath string "相对路径下的图片"
func getImg(relativePath string) *tb.Photo {
	AllConfig := GetConfig()
	picSavePath := filepath.Join(AllConfig.ProjectFolder, relativePath)
	img := &tb.Photo{File: tb.FromDisk(picSavePath)}
	return img
}

// @title TgBot主动发送信息给user
func TgBotSendMsgToUser(userId int, msgStr string) {
	AllConfig := GetConfig()
	b, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		//URL: "http://195.129.111.17:8012",
		Token:  AllConfig.TgToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	var user tb.User
	user.ID = int64(userId)
	b.Send(&user, msgStr)
}

// @title TgBot主动发送图片给user
func TgBotSendImgToUser(userId int) {
	AllConfig := GetConfig()
	b, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		//URL: "http://195.129.111.17:8012",
		Token:  AllConfig.TgToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	var user tb.User
	user.ID = int64(userId)
	pieImg := getImg("/download/pie.png")
	tableImg := getImg("/download/table.png")
	b.Send(&user, pieImg)
	b.Send(&user, tableImg)
}

// @title 生成完整的通知模板
// @return msgStr string "完整的糖浆池信息"
func (SyrupPools *SyrupPools) GenerateTgFullMsg() string {
	var msgStr string
	nowStr := time.Now().Format("2006-01-02 15:04:05") //获取当前时间
	msgStr = msgStr + "Updated at: " + nowStr + "\n"
	for index, _ := range SyrupPools.SyPools {
		msgStr = msgStr + "===============\n"
		sPool := SyrupPools.SyPools[index]
		// 保留4位小数
		sPoolStr := fmt.Sprintf("SyrupPool ID:   %s", sPool.SousId) + "\n" + fmt.Sprintf("RewardToken Name/Symbol:   %s/%s", sPool.Token.Name, sPool.Token.Symbol) + "\n" + fmt.Sprintf("RewardToken Price(USD):   %s", util.BigFloat4Decimal(sPool.Token.Price.String())) + "\n" + fmt.Sprintf("RewardToken ContractAddr:   %s", sPool.Token.ContractAddr) + "\n" + fmt.Sprintf("Daily/Weekly/Monthly/Yearly Profit(USD)(Per100Cake):   %s/%s/%s/%s", util.BigFloat4Decimal(sPool.HundredCakeDailyEarn.String()), util.BigFloat4Decimal(sPool.HundredCakeWeekEarn.String()), util.BigFloat4Decimal(sPool.HundredCakeMonthEarn.String()), util.BigFloat4Decimal(sPool.HundredCakeYearEarn.String())) + "\n" + fmt.Sprintf("SyrupPool TotalStaked Cake:   %s", util.BigFloat4Decimal(sPool.StakedCake)) + "\n" + fmt.Sprintf("SyrupPool StartTime/EndTime:   %s/%s", sPool.StartTime, sPool.EndTime) + "\n" + fmt.Sprintf("RewardToken VolumeUSDChange24h:   %s", util.Float64ToPercentage(sPool.Token.VolumeUSDChange24h)) + "\n"

		msgStr = msgStr + sPoolStr
	}
	return msgStr
}

// @title 生成简要的通知模板
// @description 基于需求:大多数时间关注的是一些关键的信息,少部分时间会需要查询更完整的细节信息
// @return msgStr string "简要的信息:Token名称、每周期收益"
func (SyrupPools *SyrupPools) GenerateTgShortlyMsg() string {
	var msgStr string
	nowStr := time.Now().Format("2006-01-02 15:04:05") //获取当前时间
	msgStr = msgStr + "Updated at: " + nowStr + "\n"
	for index, _ := range SyrupPools.SyPools {
		msgStr += "===============\n"
		sPool := SyrupPools.SyPools[index]
		sPoolStr := fmt.Sprintf("RewardToken Name/Symbol:   %s/%s", sPool.Token.Name, sPool.Token.Symbol) + "\n" + fmt.Sprintf("Profit(USD)(Per100Cake):   ") + "\n" + fmt.Sprintf("Daily:   %s", util.BigFloat4Decimal(sPool.HundredCakeDailyEarn.String())) + "\n" + fmt.Sprintf("Weekly:   %s", util.BigFloat4Decimal(sPool.HundredCakeWeekEarn.String())) + "\n" + fmt.Sprintf("Monthly:   %s", util.BigFloat4Decimal(sPool.HundredCakeMonthEarn.String())) + "\n" + fmt.Sprintf("Yearly :   %s", util.BigFloat4Decimal(sPool.HundredCakeYearEarn.String())) + "\n"
		msgStr += sPoolStr
	}
	return msgStr
}
