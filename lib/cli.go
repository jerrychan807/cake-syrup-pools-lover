package lib

import (
	"cake-syrup-pools-lover/util"
	"flag"
	"fmt"
	"os"
	"runtime"
)

// 一个空结构体来表征我们的命令行
type CommandLine struct{}

// 命令行程序首先应打印所有的命令及其用法
// start_tg_server命令:开启tgBot,接受指令。
// start_http_server命令:1.更新数据库里的Syrup信息 2.有新池则通知所有订阅了的用户
//
func (cli *CommandLine) printUsage() {
	fmt.Println("Pancake Syrup Pool Lover, usage is as follows:")
	fmt.Println("--------------------------------------------------------------------------------------------------------------")
	fmt.Println("start_tg_server  ----> Start tg bot Server")
	fmt.Println("start_http_server  ----> Start http Server")
	fmt.Println("update_pool_msg  ----> Update Syrup Pool Info,try to send alert to users")
	fmt.Println("onlytest  ----> onlytest")
	fmt.Println("--------------------------------------------------------------------------------------------------------------")
}

func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

// 使用go语言自带的flag库将各个命令注册
func (cli *CommandLine) Run() {
	cli.validateArgs()

	startServerCmd := flag.NewFlagSet("start_tg_server", flag.ExitOnError)
	startHttpServerCmd := flag.NewFlagSet("start_http_server", flag.ExitOnError)
	updateSyrupPoolMsgCmd := flag.NewFlagSet("update_pool_msg", flag.ExitOnError)
	//onlyTestCmd := flag.NewFlagSet("onlytest", flag.ExitOnError)

	switch os.Args[1] {
	case "start_tg_server":
		err := startServerCmd.Parse(os.Args[2:])
		util.Handle(err)
	case "start_http_server":
		err := startHttpServerCmd.Parse(os.Args[2:])
		util.Handle(err)
	case "update_pool_msg":
		err := updateSyrupPoolMsgCmd.Parse(os.Args[2:])
		util.Handle(err)
	//case "onlytest":
	//	err := onlyTestCmd.Parse(os.Args[2:])
	//	utils.Handle(err)
	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if startServerCmd.Parsed() {
		cli.startServer()
	}
	if startHttpServerCmd.Parsed() {
		cli.startHttpServerCmd()
	}
	if updateSyrupPoolMsgCmd.Parsed() {
		cli.updateSyrupPool()
	}
	//if onlyTestCmd.Parsed() {
	//	cli.onlyTest()
	//}
}

func (cli *CommandLine) startServer() {
	TgBotStartServer()
}

func (cli *CommandLine) startHttpServerCmd() {
	StartSimpleServer()
}

func (cli *CommandLine) updateSyrupPool() {
	poolsTsxfileUrl := "https://raw.githubusercontent.com/pancakeswap/pancake-frontend/develop/src/config/constants/pools.tsx"
	fmt.Printf("[*] Request Url: %s\n", poolsTsxfileUrl)
	// 获取pancake github上的糖浆池配置列表
	serializedPoolConfigStr := GetSerializedPoolConfigStr(poolsTsxfileUrl)
	//fmt.Println(serializedPoolConfigStr)

	// 解析糖浆池配置列表,获取SyrupPool数组
	SyrupPools := GenerateSyrupPools(serializedPoolConfigStr)
	//fmt.Println("$$$$$$$$$$$$$$$$$$$$$$")
	//for _, syrupPool := range SyrupPools.SyPools {
	//	fmt.Printf("[*] syrupPool %+v \n", syrupPool)
	//	fmt.Printf("[*] syrupPool.HundredCakeDailyEarn %s \n", syrupPool.HundredCakeDailyEarn.Text('e', 1024))
	//}

	// 生成此次糖浆池配置信息的MD5哈希值
	SyrupPools.CalcSyrupPoolsMd5()
	// 比较md5
	ifChange := CheckSyrupPoolsContractsChange(SyrupPools.MD5)
	fmt.Printf("[*] ifChange : %t \n", ifChange)

	// 填充糖浆池信息
	SyrupPools.UpdateSyrupPoolsTokenInfo()

	// 生成piechart html文件
	CreateRosePieChart(&SyrupPools)
	CreateStakedPieChart(&SyrupPools)
	CreatePoolDailyYieldPieChart(&SyrupPools)
	// 生成糖浆池信息表格html文件
	GenerateSyrupPoolTableHtml(&SyrupPools)
	piePicSavePath := GetHTML2Image("http://127.0.0.1:8080/convert/html2image?u=doctron&p=lampnick&url=http://127.0.0.1:9090/pie.html&customClip=true&clipX=0&clipY=0&clipWidth=900&clipHeight=500&clipScale=2&format=jpeg&Quality=80", "/download/pie.png")
	stakedPiePicSavePath := GetHTML2Image("http://127.0.0.1:8080/convert/html2image?u=doctron&p=lampnick&url=http://127.0.0.1:9090/stakedPie.html&customClip=true&clipX=0&clipY=0&clipWidth=900&clipHeight=500&clipScale=2&format=jpeg&Quality=80", "/download/stakedPie.png")
	dailyYieldPiePicSavePath := GetHTML2Image("http://127.0.0.1:8080/convert/html2image?u=doctron&p=lampnick&url=http://127.0.0.1:9090/dailyYieldPie.html&customClip=true&clipX=0&clipY=0&clipWidth=900&clipHeight=500&clipScale=2&format=jpeg&Quality=80", "/download/dailyYieldPie.png")

	tablePicSavePath := GetHTML2Image("http://127.0.0.1:8080/convert/html2image?u=doctron&p=lampnick&url=http://127.0.0.1:9090/table.html", "/download/table.png")

	fmt.Printf("[*] Get SyrupPools DailyEarn Pie Image : %s \n", piePicSavePath)
	fmt.Printf("[*] Get SyrupPools Staked Cake Pie Image : %s \n", stakedPiePicSavePath)
	fmt.Printf("[*] Get SyrupPools daily Yield Pie Image : %s \n", dailyYieldPiePicSavePath)
	fmt.Printf("[*] Get SyrupPools table Image : %s \n", tablePicSavePath)

	// 生成通知信息
	msg := SyrupPools.GenerateTgFullMsg()
	//smsg := SyrupPools.GenerateTgShortlyMsg()
	// 更新数据库里的糖浆池信息
	// 完整信息模板
	SaveSyrupPoolStr("pools", msg)
	// 简要信息模板
	//SaveSyrupPoolStr("pools_shortly", smsg)

	//fmt.Println("$$$$$$$$$$$ smsg $$$$$$$$$$$")
	//fmt.Println(msg)

	// 糖浆池配置信息更新(合约地址有变动),有新的糖浆池
	// ifChange = true
	fmt.Printf("[*] need to alert user: %b \n", ifChange)
	if ifChange {
		uids := QueryAllChatId()
		// 发生tg提醒
		fmt.Println(uids)
		for _, uid := range uids {
			fmt.Printf("[*] new pool, try to alert user: %d", uid)
			TgBotSendImgToUser(uid)
		}
		// 记录本次md5
		WriteStrInFile(SyrupPools.MD5)
	}
}
