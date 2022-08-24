package lib

import (
	"cake-syrup-pools-lover/util"
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
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

	// 查询最新糖浆池信息
	b.Handle("/syrup_pools", func(m *tb.Message) {
		poolMsg := QuerySyrupPoolStr()
		b.Send(m.Sender, poolMsg)
	})

	fmt.Println("[*] TgBot StartServer!!!")
	b.Start()
}

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

func (SyrupPools *SyrupPools) GenerateTgMsg() string {
	var msgStr string
	nowStr := time.Now().Format("2006-01-02 15:04:05") //获取当前时间
	msgStr = msgStr + "Updated at: " + nowStr + "\n"
	for index, _ := range SyrupPools.SyPools {
		msgStr = msgStr + "===============\n"
		sPool := SyrupPools.SyPools[index]
		// 保留4位小数
		sPoolStr := fmt.Sprintf("SyrupPool ID:   %s", sPool.SousId) + "\n" + fmt.Sprintf("RewardToken Name/Symbol:   %s/%s", sPool.Token.Name, sPool.Token.Symbol) + "\n" + fmt.Sprintf("RewardToken Price(USD):   %s", util.BigFloat4Decimal(sPool.Token.Price.String())) + "\n" + fmt.Sprintf("RewardToken ContractAddr:   %s", sPool.Token.ContractAddr) + "\n" + fmt.Sprintf("Daily/Weekly/Monthly/Yearly Profit(USD)(Per100Cake):   %s/%s/%s/%s", util.BigFloat4Decimal(sPool.HundredCakeDailyEarn.String()), util.BigFloat4Decimal(sPool.HundredCakeWeekEarn.String()), util.BigFloat4Decimal(sPool.HundredCakeMonthEarn.String()), util.BigFloat4Decimal(sPool.HundredCakeYearEarn.String())) + "\n" + fmt.Sprintf("SyrupPool TotalStaked Cake:   %s", util.BigFloat4Decimal(sPool.StakedCake)) + "\n" + fmt.Sprintf("SyrupPool EndTime:   %s", sPool.EndTime) + "\n"
		msgStr = msgStr + sPoolStr
	}
	return msgStr
}
