package lib

import (
	"cake-syrup-pools-lover/util"
	"fmt"
	"testing"
)

//func TestQueryErc20(t *testing.T) {
//	client := rpcClient("https://bsc.mytokenpocket.vip")
//	QueryErc20(client)
//	t.Logf("[*] sousIds: %s\n", sousIds)
//	if len(sousIds) == 0 {
//		t.Errorf("%s \n", sousIds)
//	}
//}

//func TestQueryErc20TokenInfo(t *testing.T) {
//	Erc20Token := QueryErc20TokenInfo()
//	t.Logf("[*] Erc20Token: %+v \n", Erc20Token)
//	if Erc20Token == (Token{}) {
//		t.Error("Erc20Token Struct is empty")
//	}
//}

func TestCalcBNBPriceByPancakeSwap(t *testing.T) {
	amountOut := CalcBNBPriceByPancakeSwap()
	t.Logf("[*] amountOut: %+v \n", amountOut)
	if amountOut == "" {
		t.Error("amountOut is empty")
	}
}

func TestCalcTokenPriceByPancakeSwap(t *testing.T) {
	// 创建不同的子测试用例
	t.Run("peel_token", func(t *testing.T) {
		bnbPriceUsd := CalcBNBPriceByPancakeSwap()
		fmt.Printf("[*] bnbPriceUsd:%s \n", bnbPriceUsd)
		// peel
		shitTokenAddr := "0x734548a9e43d2d564600b1b2ed5be9c2b911c6ab"
		shitTokenPriceBnb := CalcTokenPriceByPancakeSwap(shitTokenAddr)
		shitTokenPriceUsd := util.BigFloatMulBigFloat(bnbPriceUsd, shitTokenPriceBnb)
		t.Logf("[*] shitTokenPriceUsd: %s \n", shitTokenPriceUsd)
	})

	t.Run("pancake_token", func(t *testing.T) {
		bnbPriceUsd := CalcBNBPriceByPancakeSwap()
		fmt.Printf("[*] bnbPriceUsd:%s \n", bnbPriceUsd)
		// peel
		shitTokenAddr := "0x0e09fabb73bd3ade0a17ecc321fd13a19e81ce82"
		shitTokenPriceBnb := CalcTokenPriceByPancakeSwap(shitTokenAddr)
		shitTokenPriceUsd := util.BigFloatMulBigFloat(bnbPriceUsd, shitTokenPriceBnb)
		t.Logf("[*] shitTokenPriceUsd: %s \n", shitTokenPriceUsd)
	})

}

/*
func TestGetTokenPriceByPancakeApi(t *testing.T) {
	var input = `{"updated_at":1660728906554,"data":{"name":"Meta Apes Peel","symbol":"PEEL","price":"0.126858198411895729756863679343","price_BNB":"0.0004009516337587390261547463939483"}}`

	//shitTokenAddr := "0x0e09fabb73bd3ade0a17ecc321fd13a19e81ce82"
	shitTokenPriceUsd := GetTokenPriceByPancakeApi(input)
	t.Logf("[*] shitTokenPriceUsd: %+v \n", shitTokenPriceUsd)
}
*/