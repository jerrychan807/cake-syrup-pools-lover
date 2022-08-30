package lib

import (
	"cake-syrup-pools-lover/util"
	"fmt"
	"github.com/gocolly/colly"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

// @title 从pools.tsx的response body匹配出糖浆池配置字符串
// @description 1.https请求 https://github.com/pancakeswap/pancake-frontend/blob/develop/src/config/constants/pools.tsx
// 2.糖浆池池子信息在livePools: SerializedPoolConfig[]
func GetSerializedPoolConfigStr(url string) string {
	c := colly.NewCollector()
	var serializedPoolConfigStr string

	// extract status code
	c.OnResponse(func(r *colly.Response) {

		fmt.Println("[*] response received", r.StatusCode)
		// fmt.Println(r.Ctx.Get("url"))
		// 打印body html源码
		//fmt.Println("Visited", string(r.Body))

		var re = regexp.MustCompile(`(?s)livePools: SerializedPoolConfig(.*)const finishedPools`)
		match := re.FindString(string(r.Body))
		//fmt.Println(match, "found at index", i)
		serializedPoolConfigStr = match
	})

	c.Visit(url)
	return serializedPoolConfigStr
}

// @title 获取糖浆池信息
// @description 解析糖浆池配置列表,得到糖浆池信息
func GetSyrupPoolsByBscToolsWeb(url string) string {
	c := colly.NewCollector()
	var serializedPoolConfigStr string

	// extract status code
	c.OnResponse(func(r *colly.Response) {

		fmt.Println("[*] response received", r.StatusCode)
		// fmt.Println(r.Ctx.Get("url"))
		// 打印body html源码
		//fmt.Println("Visited", string(r.Body))

		var re = regexp.MustCompile(`(?s)livePools: SerializedPoolConfig(.*)const finishedPools`)
		match := re.FindString(string(r.Body))
		//fmt.Println(match, "found at index", i)
		serializedPoolConfigStr = match

	})

	c.Visit(url)
	return serializedPoolConfigStr
}

// @title 保存html渲染后的图片
// @description 使用本地的http和html2img服务,生成图片,保存在download文件夹中
func GetHTML2Image(url string, picPath string) string {
	//pic := base
	//idx := strings.LastIndex(url, "/")
	AllConfig := GetConfig()
	//picSavePath := filepath.Join(AllConfig.ProjectFolder, "/download/pie.png")
	picSavePath := filepath.Join(AllConfig.ProjectFolder, picPath)
	v, err := http.Get(url)
	if err != nil {
		fmt.Printf("Http get [%v] failed! %v", url, err)
		log.Fatal(err)
	}
	defer v.Body.Close()
	content, err := ioutil.ReadAll(v.Body)
	if err != nil {
		fmt.Printf("Read http response failed! %v", err)
		log.Fatal(err)
	}
	err = ioutil.WriteFile(picSavePath, content, 0666)
	if err != nil {
		fmt.Printf("Save to file failed! %v", err)
		log.Fatal(err)
	}
	return picSavePath
}

// @title 获取糖浆池信息
// @description 解析糖浆池配置列表,得到初始化糖浆池信息
// @param serializedPoolConfigStr string "糖浆池配置列表"
func GenerateSyrupPools(serializedPoolConfigStr string) SyrupPools {
	var SyrupPools SyrupPools
	sousIds, earningTokens, contractAddrs, tokenPerBlocks := ParsePoolsInfo(serializedPoolConfigStr)
	/*
		{0 0xa5f8C5Dbd5F286960b9d90548680aE5ebFf07652 0   0 0   cake 10  {   0 <nil> <nil> 0 0}}
		{291 0x56D6955Ba6404647191DD7A5D65A5c9Fe43905e1 0   0 0   pstake 1.1574  {   0 <nil> <nil> 0 0}}
		{290 0x288d1aD79c113552B618765B4986f7DE679367Da 0   0 0   peel 4.34  {   0 <nil> <nil> 0 0}}
		{289 0x595B7AF4F1828AB4953792482b01B2AFC4A46b72 0   0 0   shell 89.699  {   0 <nil> <nil> 0 0}}
		{288 0x28cc814bE3B994187B7f8Bfed10516A84A671119 0   0 0   high 0.1332  {   0 <nil> <nil> 0 0}}
		{287 0xda6F750be1331963E5772BEe757062f6bddcEA4C 0   0 0   ole 2.2569  {   0 <nil> <nil> 0 0}}
		{286 0x86471019Bf3f403083390AC47643062e15B0256e 0   0 0   trivia 4.976  {   0 <nil> <nil> 0 0}}
	*/
	for i := 0; i < len(sousIds); i++ {
		// 去掉第一个mastercheif合约
		if sousIds[i] != "0" {
			sp := SyrupPool{SousId: sousIds[i], EarningToken: earningTokens[i], ContractAddr: contractAddrs[i], TokenPerBlock: tokenPerBlocks[i]}
			//fmt.Println(sp)
			// 初始化糖浆池数组
			SyrupPools.SyPools = append(SyrupPools.SyPools, &sp)
		}

	}
	return SyrupPools
}

// @title 生成md5
// @description 糖浆池数组的ContractAddr列表,生成此次糖浆池配置信息的MD5哈希值
func (SyrupPools *SyrupPools) CalcSyrupPoolsMd5() {
	var baseStr string
	for _, syrupPool := range SyrupPools.SyPools {
		//fmt.Printf("[*] contractAddr: %s \n", syrupPool.ContractAddr)
		baseStr = baseStr + syrupPool.ContractAddr
	}
	SyrupPools.MD5 = util.Md5V(baseStr)
	fmt.Printf("[*] SyrupPools.MD5 : %s \n", SyrupPools.MD5)
}

// @title 填充糖浆池信息
// @description 1.获取token价格 2.计算出Roi 投资回报率
func (SyrupPools *SyrupPools) UpdateSyrupPoolsTokenInfo() {
	for index, _ := range SyrupPools.SyPools {
		fmt.Println("=================================")
		SyrupPools.SyPools[index].GetSyrupPoolInfo(SyrupPools.SyPools[index].ContractAddr)

		// 通过PancakeswapApi 获取token价格
		apiBaseUrl := "https://api.pancakeswap.info/api/v2/tokens/"
		apiFullUrl := fmt.Sprintf("%s%s", apiBaseUrl, SyrupPools.SyPools[index].Token.ContractAddr)
		//fmt.Println("[*] apiFullUrl", apiFullUrl)
		jsonData := GetPancakeApiJsonData(apiFullUrl)
		tokenName, tokenSymbol, priceUsd := GetTokenPriceByPancakeApi(jsonData)

		f1 := new(big.Float)
		f1.SetString(priceUsd)
		SyrupPools.SyPools[index].Token.Price = f1
		SyrupPools.SyPools[index].Token.Name = tokenName
		SyrupPools.SyPools[index].Token.Symbol = tokenSymbol

		fmt.Printf("[*] syrupPool.Token.Price:%s \n", SyrupPools.SyPools[index].Token.Price.String())

		// 通过DexGuruApi 获取token市场信息
		jsonData1 := GetDexGuruApiJsonData(SyrupPools.SyPools[index].Token.ContractAddr)
		p := GetTokenInfoByDexguru(jsonData1)
		SyrupPools.SyPools[index].Token.Txns24h = p.DexguruApiToken[0].Txns24h
		SyrupPools.SyPools[index].Token.Txns24hChange = p.DexguruApiToken[0].Txns24hChange
		SyrupPools.SyPools[index].Token.Volume24hUSD = p.DexguruApiToken[0].Volume24hUSD
		SyrupPools.SyPools[index].Token.VolumeUSDChange24h = p.DexguruApiToken[0].VolumeUSDChange24h
		SyrupPools.SyPools[index].Token.LiquidityUSD = p.DexguruApiToken[0].LiquidityUSD
		SyrupPools.SyPools[index].Token.LiquidityUSDChange24h = p.DexguruApiToken[0].LiquidityUSDChange24h
		SyrupPools.SyPools[index].Token.PriceUSDChange24h = p.DexguruApiToken[0].PriceUSDChange24h
		SyrupPools.SyPools[index].Token.MarketCap = p.DexguruApiToken[0].MarketCap
		SyrupPools.SyPools[index].Token.MarketCapChange24h = p.DexguruApiToken[0].MarketCapChange24h
		SyrupPools.SyPools[index].Token.LogoURI = p.DexguruApiToken[0].LogoURI[0]

		// 100cake daily earn usd 每日利润usd
		SyrupPools.SyPools[index].CalcRewadTokenDailyEarn()
		SyrupPools.SyPools[index].CalcHundredCakeDailyEarn()
		fmt.Printf("[*] syrupPool %+v \n", SyrupPools.SyPools[index])
		fmt.Println("=================================")
		//break
	}

	//fmt.Printf("[*] SyrupPools %+v \n", SyrupPools)
	// 获取token信息
	//QueryErc20TokenInfo()

}

// @title 字符串写入文件
func WriteStrInFile(content string) {
	//AllConfig := GetConfig()
	filePath := getLastMd5FilePath()
	dstFile, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	dstFile.WriteString(content)
}

// @title 获得md5文件的路径
func getLastMd5FilePath() string {
	fileName := "lastMd5.log"
	AllConfig := GetConfig()
	//currentPath, _ := os.Getwd()
	//currentPath := AllConfig.ProjectFolder + fileName
	filePath := filepath.Join(AllConfig.ProjectFolder, fileName)
	return filePath
}

// @title 从文件读取字符串
func ReadLastMd5() string {
	filePath := getLastMd5FilePath()
	stream, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	readString := string(stream)
	return readString
}

// @title 通过比较两次MD5值,查看糖浆池配置信息有没发生改变
func CheckSyrupPoolsContractsChange(currentMd5 string) bool {
	// 查询上次的md5值
	lastMD5 := ReadLastMd5()
	fmt.Printf("[*] lastMD5: %s \n", lastMD5)
	return (lastMD5 != currentMd5)
	//return (lastMD5 == currentMd5)
}
