package lib

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

// @title 通过DexGuru的Api查询Token的市场信息
// @param shitTokenAddr string "代币合约地址"
func GetTokenInfoByDexguru(jsonData string) DexguruApiJson {
	//var input = `{"total":1,"data":[{"marketType":"token","address":"0xe5ba47fd94cb645ba4119222e34fb33f59c7cd90","id":"0xe5ba47fd94cb645ba4119222e34fb33f59c7cd90-bsc","symbols":["SAFUU"],"underlyingAddresses":null,"name":"Safuu","description":"Safuu/SAFUU","txns24h":1429,"txns24hChange":-0.03184281842818428,"verified":true,"decimals":5,"volume24h":0.0,"volume24hUSD":508233.6669300458,"volume24hETH":1810.1267414374197,"liquidityUSD":4940465.7026604,"liquidityETH":17127.018663631,"priceUSD":4.7668026489695094,"priceETH":0.016524984252149117,"priceUSDChange24h":-0.028118112497613614,"priceETHChange24h":-0.06342087850789545,"timestamp":1646175168,"blockNumber":0,"AMM":"uniswap","network":"bsc","tokenListsNames":["CoinGecko"],"marketCap":93527812.49009514,"marketCapChange24h":-2.929026625005399,"liquidityUSDChange24h":0.01555776285112074,"liquidityETHChange24h":-0.021331491422338894,"volumeUSDChange24h":-0.38769260704658837,"volumeETHChange24h":-0.38912142283559364,"logoURI":["https://assets.dex.guru/icons/0xe5ba47fd94cb645ba4119222e34fb33f59c7cd90-bsc.jpg"]}]}`

	var p DexguruApiJson
	// 解析json数据到post中
	err := json.Unmarshal([]byte(jsonData), &p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("[*] Dexguru Api Json Response: %+v\n", p)
	//fmt.Printf("Txns24h: %s\n", p.DexguruApiToken[0].Txns24hChange)
	//fmt.Printf("LogoURI: %s\n", p.DexguruApiToken[0].LogoURI[0])
	//return p.PancakeApiToken.Price
	return p
}

// @title 获得通过DexGuru的Api返回的json数据
func GetDexGuruApiJsonData(shitTokenAddr string) string {
	c := colly.NewCollector()

	var postData = fmt.Sprintf(`{"ids":["%s-bsc"],"network":"eth,optimism,bsc,polygon,fantom,arbitrum,celo,avalanche"}`, shitTokenAddr)

	var jsonData string

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Content-Type", "application/json;charset=UTF-8")
	})

	// extract status code
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("[*] response received", r.StatusCode)
		jsonData = string(r.Body)
		//fmt.Println(jsonData)
	})

	c.PostRaw("https://api.dex.guru/v3/tokens", []byte(postData))
	return jsonData
}
