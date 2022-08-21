package lib

import (
	erc20 "cake-syrup-pools-lover/abi/erc20"
	pancakeRouter "cake-syrup-pools-lover/abi/pancakeRouter"
	"cake-syrup-pools-lover/abi/syrupPool"
	"cake-syrup-pools-lover/util"
	"encoding/json"
	"github.com/gocolly/colly"

	//Token "cake-syrup-pools-lover/types"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
)

func QueryBalance() {
	client, err := ethclient.Dial("https://bsc.mytokenpocket.vip")
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress("0x6d80fdaa5cdbff3db45a853d41d779509da85da3")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance) // 25893180161173005034

	blockNumber := big.NewInt(5532993)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(balanceAt) // 25729324269165216042
	//
	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue) // 25.729324269165216041
	//
	//pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	//fmt.Println(pendingBalance) // 25729324269165216042
}

func getRpcClient(rpcUrl string) *ethclient.Client {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

// @title
// @description 获取erc20Token的一个contract实例
// @param client *ethclient.Client "rpc连接client"
func getTokenInstance(contractAddr common.Address, client *ethclient.Client) *erc20.Erc20 {
	instance, err := erc20.NewErc20(contractAddr, client)
	if err != nil {
		log.Fatal(err)
	}
	return instance
}

// @title
// @description 获取pancakeRouter的一个contract实例
func getRouterContractInstance(contractAddr common.Address, client *ethclient.Client) *pancakeRouter.PancakeRouter {
	instance, err := pancakeRouter.NewPancakeRouter(contractAddr, client)
	if err != nil {
		log.Fatal(err)
	}
	return instance
}

// @title
// @description 获取SyrupPool的一个contract实例
func GetSyrupPoolContractInstance(contractAddr common.Address, client *ethclient.Client) *syrupPool.SyrupPool {
	instance, err := syrupPool.NewSyrupPool(contractAddr, client)
	if err != nil {
		log.Fatal(err)
	}
	return instance
}

// @title
// @description 获取erc20Token的代币名称
// @param instance *erc20.Erc20 "erc20实例"
func (t *Token) getTokenName(instance *erc20.Erc20) string {
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	return name
}

// @title 查询erc20Token的代币符号
// @param instance *erc20.Erc20 "erc20实例"
func (t *Token) getTokenSymbol(instance *erc20.Erc20) string {
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	return symbol
}

// @title 查询erc20Token的小数点
// @param instance *erc20.Erc20 "erc20实例"
func (t *Token) getTokenDecimals(instance *erc20.Erc20) uint8 {
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	return decimals
}

/*
// @title 查询erc20Token的价格(usd)
// @description
// @param instance *erc20.Erc20 "erc20实例"
func (t *Token) getTokenPriceUsd(instance *erc20.Erc20) uint8 {
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	return decimals
}
*/

// @title
// @description 查询某地址的erc20Token余额
// @param instance *erc20.Erc20 "erc20实例"
func (t *Token) getAddrBalance(instance *erc20.Erc20, Addr common.Address) string {
	balanceWei, err := instance.BalanceOf(&bind.CallOpts{}, Addr)
	if err != nil {
		log.Fatal(err)
	}

	balance := util.WeiToEther(balanceWei)
	//fmt.Printf("[*] balance: %s \n", balance)
	return balance
}

// @title 通过pancakeSwapContract 计算出BnbPrice
// @description Refs: https://gist.github.com/Linch1/ede03999f483f2b1d5fcac9e8b312f2c
// @param
func CalcBNBPriceByPancakeSwap() string {
	AllConfig := GetConfig()
	client := getRpcClient(AllConfig.RpcUrl)

	BNBTokenAddr := common.HexToAddress(AllConfig.BnbContractAddr)
	UsdtTokenAddr := common.HexToAddress(AllConfig.UsdtContractAddr)

	var path []common.Address
	path = append(path, BNBTokenAddr)
	path = append(path, UsdtTokenAddr)
	//fmt.Println("[*] path: ")
	//fmt.Println(path)

	bnbToSellAmount := new(big.Float).SetFloat64(1.0)
	bnbToSellAmountWei := util.EtherToWei(bnbToSellAmount, 18)
	//fmt.Printf("[*] bnbToSellAmountWei: %d \n", bnbToSellAmountWei)

	contractAddr := common.HexToAddress(AllConfig.PancakeRouterAddr)
	routerIns := getRouterContractInstance(contractAddr, client)

	amountOut, err := routerIns.GetAmountsOut(&bind.CallOpts{}, bnbToSellAmountWei, path)
	if err != nil {
		log.Fatal(err)
	}
	bnbPriceUsd := util.WeiToEther(amountOut[1])

	fmt.Println(bnbPriceUsd)
	return bnbPriceUsd
}

// @title 通过pancakeSwapContract 计算出BnbPrice
// @description Refs: https://gist.github.com/Linch1/ede03999f483f2b1d5fcac9e8b312f2c
// @param shitTokenAddr string "erc20实例"
func CalcTokenPriceByPancakeSwap(shitTokenAddr string) string {
	AllConfig := GetConfig()
	client := getRpcClient(AllConfig.RpcUrl)
	BNBTokenAddr := common.HexToAddress(AllConfig.BnbContractAddr)

	var ShitToken Token
	ShitTokenAddr := common.HexToAddress(shitTokenAddr)
	shitTokenIns := getTokenInstance(ShitTokenAddr, client)
	ShitToken.Decimals = ShitToken.getTokenDecimals(shitTokenIns)

	var path []common.Address
	path = append(path, ShitTokenAddr)
	path = append(path, BNBTokenAddr)

	tokensToSellAmount := new(big.Float).SetFloat64(1.0)
	tokensToSellAmountWei := util.EtherToWei(tokensToSellAmount, int(ShitToken.Decimals))
	fmt.Printf("[*] tokensToSellAmountWei: %d \n", tokensToSellAmountWei)

	contractAddr := common.HexToAddress("0x10ED43C718714eb63d5aA57B78B54704E256024E")
	routerIns := getRouterContractInstance(contractAddr, client)
	amountOut, err := routerIns.GetAmountsOut(&bind.CallOpts{}, tokensToSellAmountWei, path)
	if err != nil {
		log.Fatal(err)
	}
	// amountOut[1] Shit等于多少个bnb,一个bnb约300$
	shitTokenPriceBnb := util.WeiToEther(amountOut[1])

	fmt.Println(shitTokenPriceBnb)
	return shitTokenPriceBnb
}

// @title 通过pancakeSwapApi 查询出Token价格
// @description Refs: https://api.pancakeswap.info/api/v2/tokens/0xc84c177ac200461e6a7208a3a1073538036d0779
// https://stackoverflow.com/questions/53293506/how-to-handle-response-json-have-custom-field-with-out-key
// @param jsonData string "API返回的json数据"
func GetTokenPriceByPancakeApi(jsonData string) (string, string, string) {
	var p PancakeApiJson
	// 解析json数据到post中
	err := json.Unmarshal([]byte(jsonData), &p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("[*] Api Json Data After handled: %+v \n", p)
	return p.PancakeApiToken.Name, p.PancakeApiToken.Symbol, p.PancakeApiToken.Price
}

func GetPancakeApiJsonData(apiFullUrl string) string {
	c := colly.NewCollector()
	var jsonData string

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Content-Type", "application/json;charset=UTF-8")
	})

	// extract status code
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("[*] response received", r.StatusCode)
		// fmt.Println(r.Ctx.Get("url"))
		// 打印body html源码
		//fmt.Println("[*] Source Code", string(r.Body))
		jsonData = string(r.Body)
	})

	c.Visit(apiFullUrl)
	return jsonData
}

// @title 查询Token的liquidity信息
// @param shitTokenAddr string "代币合约地址"
// TODO:
/*
func GetTokenLiquidity(shitTokenAddr string) string {

	var input = `{"updated_at":1660728906554,"data":{"name":"Meta Apes Peel","symbol":"PEEL","price":"0.126858198411895729756863679343","price_BNB":"0.0004009516337587390261547463939483"}}`

	var p pancakeApiJson
	// 解析json数据到post中
	err := json.Unmarshal([]byte(input), &p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Api Json Response: %+v", p)
	return p.PancakeApiToken.Price
}
*/

func QueryErc20TokenInfo(tokenAddr string) Token {
	var Erc20Token Token
	AllConfig := GetConfig()
	client := getRpcClient(AllConfig.RpcUrl)
	contractAddr := common.HexToAddress(tokenAddr)
	tokenIns := getTokenInstance(contractAddr, client)
	Erc20Token.Name = Erc20Token.getTokenName(tokenIns)
	Erc20Token.Symbol = Erc20Token.getTokenSymbol(tokenIns)
	Erc20Token.Decimals = Erc20Token.getTokenDecimals(tokenIns)
	//fmt.Printf("name: %s\n", name)         // "name: Golem Network"
	//fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"

	//fmt.Printf("wei: %s\n", bal) // "wei: 74605500647408739782407023"

	//fbal := new(big.Float)
	//fbal.SetString(bal.String())
	//value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
	//
	//fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
	return Erc20Token
}

func GetErc20TokenDecimals(tokenAddr string) uint8 {
	var Erc20Token Token
	AllConfig := GetConfig()
	client := getRpcClient(AllConfig.RpcUrl)
	contractAddr := common.HexToAddress(tokenAddr)
	tokenIns := getTokenInstance(contractAddr, client)
	return Erc20Token.getTokenDecimals(tokenIns)

}

// @title 查询用户的cake余额
// @param shitTokenAddr string "代币合约地址"
func QueryAddrCakeBalance(userAddr string) string {
	var Cake Token
	AllConfig := GetConfig()
	client := getRpcClient(AllConfig.RpcUrl)
	contractAddr := common.HexToAddress(AllConfig.CakeContractAddr)
	tokenIns := getTokenInstance(contractAddr, client)
	addr := common.HexToAddress(userAddr)
	balance := Cake.getAddrBalance(tokenIns, addr)
	//fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
	return balance
}
