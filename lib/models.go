package lib

import "math/big"

// Config 配置
type Config struct {
	RpcUrl            string // rpcurl
	PancakeRouterAddr string // pancakeV2 Router合约地址
	BnbContractAddr   string // BNB合约地址
	UsdtContractAddr  string // USDT合约地址
	CakeContractAddr  string // Cake合约地址
	ProjectFolder     string // 项目主文件夹
	TgToken           string // TgBotToken
	BcPath            string // 数据库文件夹路径
	BcFile            string // 数据库文件路径
	ChatIdFile        string // Tg订阅用户数据库文件
}

// SyrupPools 多个糖浆池信息
type SyrupPools struct {
	SyPools []*SyrupPool // 糖浆池数组
	MD5     string
}

// SyrupPool 单个糖浆池信息
type SyrupPool struct {
	SousId       string // id
	ContractAddr string // 糖浆池合约地址
	//Apr                  uint       // 年华百分比
	MaxEndTime           string     // 加速奖励结束时间 utc+8
	StartTime            string     // 开始时间 utc+8
	StartBlockNum        uint64     // 开始区块数
	BonusEndBlockNum     uint64     // 奖励结束区块数
	EndTime              string     // 结束时间 utc+8
	LimitPerUser         string     // 每个用户质押cake限额
	EarningToken         string     // 挖出的token名
	TokenPerBlock        string     //
	RewardPerBlock       string     // 每区块奖励的token数量
	RewardDailyEarnUsd   *big.Float // 一天可以产出的U
	HundredCakeDailyEarn *big.Float // 100cake每日可赚取多少u
	HundredCakeWeekEarn  *big.Float // 100cake每周可赚取多少u
	HundredCakeMonthEarn *big.Float // 100cake每月可赚取多少u
	HundredCakeYearEarn  *big.Float // 100cake每年可赚取多少u
	AaccTokenPerShare    string     // 区间单位token奖励数量 累加值
	StakedCake           string     // 池子cake总质押数量
	Token                Token      // 奖励Token
}

// Token 糖浆池奖励Token信息
type Token struct {
	ContractAddr string     // Token合约地址
	Name         string     // Token名字
	Symbol       string     // Token符号
	Decimals     uint8      // Token小数位
	Price        *big.Float // 价格(usd)
	//TotalLiquidity            *big.Int   // Liquidity(usd)
	//DayPriceChangePercentage  uint       // 1日涨跌幅百分比
	//WeekPriceChangePercentage uint       // 7日涨跌幅百分比
}

type PancakeApiJson struct {
	UpdatedAt       uint            `json:"updated_at"`
	PancakeApiToken PancakeApiToken `json:"data"`
}

type PancakeApiToken struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Price    string `json:"price"`
	PriceBNB string `json:"price_BNB"`
}
