package lib

import (
	"cake-syrup-pools-lover/abi/syrupPool"
	"cake-syrup-pools-lover/config"
	"cake-syrup-pools-lover/util"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
	"log"
	"math/big"
	"time"
)

// @title 返回rpcUrl
// @description 读取配置,返回rpcUrl
func GetConfig() Config {
	var AllConfig Config
	config.InitConfig()
	// web3配置
	AllConfig.RpcUrl = viper.GetString("bsc.rpcUrl")
	AllConfig.BnbContractAddr = viper.GetString("bsc.bnbContractAddr")
	AllConfig.PancakeRouterAddr = viper.GetString("bsc.pancakeRouterAddr")
	AllConfig.UsdtContractAddr = viper.GetString("bsc.usdtContractAddr")
	AllConfig.CakeContractAddr = viper.GetString("bsc.cakeContractAddr")
	// TgBot配置
	AllConfig.TgToken = viper.GetString("telegram.token")
	// Database配置
	AllConfig.BcPath = viper.GetString("project.bcPath")
	AllConfig.BcFile = viper.GetString("project.bcFile")
	AllConfig.ChatIdFile = viper.GetString("project.chatIdFile")

	AllConfig.ProjectFolder = viper.GetString("project.projectFolder")
	//fmt.Printf("Viper get rpcUrl config -->%s \n", AllConfig.RpcUrl)
	return AllConfig
}

// @title 通过SyrupPoolContract查询糖浆池合约里的信息
// @description: 重要的糖浆池字段
// 				accTokenPerShare: 187624189728(Wei) 区间单位token奖励数量 累加值
//				bonusEndBlock: 22016415(Wei) 结束奖励的区块数
//
// example: SmartChefInitializable https://bscscan.com/address/0x56D6955Ba6404647191DD7A5D65A5c9Fe43905e1#code
// @param syrupPoolAddr string "糖浆池合约地址"
func (syrupPool *SyrupPool) GetSyrupPoolInfo(syrupPoolAddr string) {
	AllConfig := GetConfig()
	client := getRpcClient(AllConfig.RpcUrl)
	contractAddr := common.HexToAddress(syrupPoolAddr)
	syrupPoolIns := GetSyrupPoolContractInstance(contractAddr, client)
	// 开始区块数
	syrupPool.StartBlockNum = GetSyrupPoolStartBlock(syrupPoolIns)
	// 奖励结束区块数
	syrupPool.BonusEndBlockNum = GetSyrupPoolBonusEndBlock(syrupPoolIns)
	// 奖励结束utc8时间
	BonusEndBlockNumUTCTime := CalcFutureBlockNumTime(syrupPool.BonusEndBlockNum)
	syrupPool.EndTime = BonusEndBlockNumUTCTime.Format("2006-01-02 15:04:05")
	// 每个用户质押cake限额
	syrupPool.LimitPerUser = GetSyrupPoolLimitPerUser(syrupPoolIns)
	// 奖励token的合约地址
	syrupPool.Token.ContractAddr = GetSyrupPoolRewardTokenContractAddr(syrupPoolIns)
	// Token精度 小数位
	syrupPool.Token.Decimals = GetErc20TokenDecimals(syrupPool.Token.ContractAddr)
	// 每区块奖励的token数量
	syrupPool.RewardPerBlock = GetSyrupPoolRewardPerBlock(syrupPoolIns, syrupPool.Token.Decimals)

	//fmt.Printf("[*] syrupPool.Token.ContractAddr: %s \n", syrupPool.Token.ContractAddr)
}

// @title 计算糖浆池每日可产生多少矿币对应的usd
// @description 矿币一天可以产出的U = 一天区块数 * 每区块奖励的token数量 * token价格
// refs: https://bsctools.xyz/pancakeswap/pools/
// @return rewardDailyEarnUsd "矿币一天可以产出的U"
func (syrupPool *SyrupPool) CalcRewadTokenDailyEarn() {
	// Bsc一天约可以出28800块
	BscDayBlockNum := new(big.Float)
	BscDayBlockNum.SetString("28800")

	// 矿币一天可以产出的token数量 = 一天区块数 * 每区块奖励的token数量
	RewardPerBlock := new(big.Float)
	RewardPerBlock.SetString(syrupPool.RewardPerBlock)
	rewardDailyTokenNum := new(big.Float)
	rewardDailyTokenNum = rewardDailyTokenNum.Mul(RewardPerBlock, BscDayBlockNum)
	// 矿币一天可以产出的U = 矿币一天可以产出的token数量 * token价格(usd)
	rewardDailyEarnUsd := new(big.Float)
	rewardDailyEarnUsd = rewardDailyEarnUsd.Mul(rewardDailyTokenNum, syrupPool.Token.Price)
	syrupPool.RewardDailyEarnUsd = rewardDailyEarnUsd
	fmt.Printf("%s","[*] rewardDailyEarnUsd  ")
	fmt.Println(rewardDailyEarnUsd)
	//return rewardDailyEarnUsd
}

// @title 计算糖浆池100cake每日可赚取多少u
// @description 矿币一天可以产出的U = 一天区块数 * 每区块奖励的token数量 * token价格
// refs: https://bsctools.xyz/pancakeswap/pools/
// @param
// @return
func (syrupPool *SyrupPool) CalcHundredCakeDailyEarn() {
	var prec uint = 1024 // 512
	// 糖浆池总质押的cake数量
	syrupPool.StakedCake = QueryAddrCakeBalance(syrupPool.ContractAddr)
	stakedCake := new(big.Float)
	stakedCake.SetString(syrupPool.StakedCake)

	hundredCake := new(big.Float)
	hundredCake.SetString("100")
	// 100cake质押后所占的比例
	rate := new(big.Float)
	rate = rate.Quo(hundredCake, stakedCake.Add(stakedCake, hundredCake))
	// 糖浆池100cake每日可赚取多少u
	HundredCakeDailyEarn := syrupPool.RewardDailyEarnUsd.Mul(syrupPool.RewardDailyEarnUsd, rate)
	syrupPool.HundredCakeDailyEarn = HundredCakeDailyEarn
	// 每周可赚取多少u
	WeekFloat, _ := new(big.Float).SetPrec(prec).SetString("7")
	tmp := new(big.Float).SetPrec(prec)
	syrupPool.HundredCakeWeekEarn = tmp.Mul(HundredCakeDailyEarn, WeekFloat)
	// 每月可赚取多少u
	MonthFloat, _ := new(big.Float).SetPrec(prec).SetString("30")
	tmp = new(big.Float).SetPrec(prec)
	syrupPool.HundredCakeMonthEarn = tmp.Mul(HundredCakeDailyEarn, MonthFloat)
	// 每年可赚取多少u
	YearFloat, _ := new(big.Float).SetPrec(prec).SetString("365")
	tmp = new(big.Float).SetPrec(prec)
	syrupPool.HundredCakeYearEarn = tmp.Mul(HundredCakeDailyEarn, YearFloat)
	fmt.Println("%s","[*] HundredCakeDailyEarn  ")
	fmt.Println(HundredCakeDailyEarn)

}

// @title 查询糖浆池合约里的StartBlock 开始区块数
// @param syrupPoolIns *syrupPool.SyrupPool "糖浆池合约实例"
func GetSyrupPoolStartBlock(syrupPoolIns *syrupPool.SyrupPool) uint64 {
	startBlock, err := syrupPoolIns.StartBlock(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	sBlock := startBlock.Uint64()
	//fmt.Printf("[*] startBlock: %d\n", startBlock)
	return sBlock
}

// @title 查询糖浆池合约的每区块奖励的token数量
// @param syrupPoolIns *syrupPool.SyrupPool "糖浆池合约实例"
func GetSyrupPoolRewardPerBlock(syrupPoolIns *syrupPool.SyrupPool, decimal uint8) string {
	rewardBlockWei, err := syrupPoolIns.RewardPerBlock(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	rewardBlock := util.WeiToEtherSpecificDecimal(rewardBlockWei, decimal)
	//rewardBlock := util.WeiToEther(rewardBlockWei)
	fmt.Printf("[*] rewardBlock: %s \n", rewardBlock)
	return rewardBlock
}

// @title 查询糖浆池合约里的奖励token合约地址
// @param syrupPoolIns *syrupPool.SyrupPool "糖浆池合约实例"
func GetSyrupPoolRewardTokenContractAddr(syrupPoolIns *syrupPool.SyrupPool) string {
	rewardTokenContractAddr, err := syrupPoolIns.RewardToken(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("[*] rewardTokenContractAddr: %s\n", rewardTokenContractAddr)
	return rewardTokenContractAddr.String()
}

// @title 查询糖浆池合约里的bonusEndBlock 奖励结束区块数
// @param syrupPoolIns *syrupPool.SyrupPool "糖浆池合约实例"
func GetSyrupPoolBonusEndBlock(syrupPoolIns *syrupPool.SyrupPool) uint64 {
	blockNum, err := syrupPoolIns.BonusEndBlock(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	blockNum64 := blockNum.Uint64()
	//fmt.Printf("[*] bonusEndBlock: %d\n", blockNum64)
	return blockNum64
}

// @title 查询糖浆池合约里的LimitPerUser  每个用户质押cake限额
// @param syrupPoolIns *syrupPool.SyrupPool "糖浆池合约实例"
func GetSyrupPoolLimitPerUser(syrupPoolIns *syrupPool.SyrupPool) string {
	limitWei, err := syrupPoolIns.PoolLimitPerUser(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	limitEther := util.WeiToEther(limitWei)
	//fmt.Printf("[*] limitEther: %d\n", limitEther)
	return limitEther
}

// @title 预估bonusEndBlock 奖励结束区块数的结束时间
// @description 通过与当前区块数的差值、BSC出块时间、预估出未来区块数的时间
func CalcFutureBlockNumTime(futureBlockNum uint64) time.Time {
	// 计算区块数之差
	blockNumberNow := GetBlockNumNow()
	tUnix := GetBlockTime(blockNumberNow)
	blocktimeNow := time.Unix(int64(tUnix), 0)
	diffBlockNum := futureBlockNum - blockNumberNow

	// 约3s一个block
	diffSeconds := diffBlockNum * uint64(3)
	FutureBlockNumUTCTime := blocktimeNow.Add(time.Second * time.Duration(diffSeconds))
	//fmt.Println(FutureBlockNumUTCTime)
	return FutureBlockNumUTCTime
}

// @title 通过blockNum查询区块时间戳
// @description :仅支持过去的区块数
// refs: https://goethereumbook.org/zh/block-query/
// @param blockNum uint64 "区块数"
func GetBlockTime(blockNum uint64) uint64 {
	AllConfig := GetConfig()
	client := getRpcClient(AllConfig.RpcUrl)
	blockNumber := big.NewInt(int64(blockNum))
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(block.Number().Uint64())     // 5671744
	//fmt.Println(block.Time()) // 1527211625
	//fmt.Println(block.Difficulty().Uint64()) // 3217000136609065
	//fmt.Println(block.Hash().Hex())          // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	//fmt.Println(len(block.Transactions()))   // 144
	tUnix := block.Time()
	//fmt.Printf("[*] tUnix: %d\n", tUnix)
	return tUnix
}

// @title 查询当前区块数
// @description
// refs: https://goethereumbook.org/zh/block-query/
// @param blockNum uint64 "区块数"
func GetBlockNumNow() uint64 {
	AllConfig := GetConfig()
	client := getRpcClient(AllConfig.RpcUrl)
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	blockNum := header.Number.Uint64()
	//fmt.Println(header.Number.String()) // 5671744
	return blockNum
}

// @title 区块时间戳转换为UTC+8时间字符串
// @description
// refs: https://goethereumbook.org/zh/block-query/
// @param tUnix uint64 "区块时间戳"
// @return nowStr string "UTC+8时间字符串"
func BlockTimestampToTimeStr(tUnix uint64) string {
	timeT := time.Unix(int64(tUnix), 0)
	//fmt.Printf("time.Time: %s\n", timeT)
	nowStr := timeT.Format("2006-01-02 15:04:05")
	//fmt.Printf("nowStr: %s\n", nowStr)
	return nowStr
}
