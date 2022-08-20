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
// startserver命令:开启tgBot,接受指令。
// updatespool命令:1.更新数据库里的Syrup信息 2.有新池则通知所有订阅了的用户

func (cli *CommandLine) printUsage() {
	fmt.Println("Pancake Syrup Pool Lover, usage is as follows:")
	fmt.Println("--------------------------------------------------------------------------------------------------------------")
	fmt.Println("startserver  ----> Start tg bot Server")
	fmt.Println("updatespool  ----> Update Syrup Pool Info,try to send alert to users")
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

	startServerCmd := flag.NewFlagSet("startserver", flag.ExitOnError)
	updateSyrupPoolCmd := flag.NewFlagSet("updatespool", flag.ExitOnError)
	//onlyTestCmd := flag.NewFlagSet("onlytest", flag.ExitOnError)

	switch os.Args[1] {
	case "startserver":
		err := startServerCmd.Parse(os.Args[2:])
		util.Handle(err)
	case "updatespool":
		err := updateSyrupPoolCmd.Parse(os.Args[2:])
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
	if updateSyrupPoolCmd.Parsed() {
		cli.updateSyrupPool()
	}
	//if onlyTestCmd.Parsed() {
	//	cli.onlyTest()
	//}
}

func (cli *CommandLine) startServer() {

	TgBotStartServer()
}

func (cli *CommandLine) updateSyrupPool() {
	poolsTsxfileUrl := "https://raw.githubusercontent.com/pancakeswap/pancake-frontend/develop/src/config/constants/pools.tsx"
	fmt.Printf("[*] Request Url: %s\n", poolsTsxfileUrl)
	// 获取pancake github上的糖浆池配置列表
	serializedPoolConfigStr := GetSerializedPoolConfigStr(poolsTsxfileUrl)
	// 解析糖浆池配置列表,获取SyrupPool数组
	SyrupPools := GenerateSyrupPools(serializedPoolConfigStr)
	// 生成此次糖浆池配置信息的MD5哈希值
	SyrupPools.CalcSyrupPoolsMd5()
	// 比较md5
	ifChange := CheckSyrupPoolsContractsChange(SyrupPools.MD5)
	fmt.Printf("[*] ifChange : %t \n", ifChange)
	// 糖浆池配置信息更新,有新的糖浆池
	//ifChange = true

	if ifChange {
		// 填充糖浆池信息
		SyrupPools.UpdateSyrupPoolsTokenInfo()

		fmt.Println("$$$$$$$$$$$$$$$$$$$$$$")
		for _, syrupPool := range SyrupPools.SyPools {
			fmt.Printf("[*] syrupPool %+v \n", syrupPool)
		}

		msg := SyrupPools.GenerateTgMsg()
		SaveSyrupPoolStr(msg)
		fmt.Println("$$$$$$$$$$$ msg $$$$$$$$$$$")
		fmt.Println(msg)
		uids := QueryAllChatId()
		// 发生tg提醒
		for _, uid := range uids {
			TgBotSendMsgToUser(uid, msg)
		}
		// 记录本次md5
		WriteStrInFile(SyrupPools.MD5)
	}
}
