package lib

import "testing"

func TestGetBlockTime(t *testing.T) {
	blockNumber := uint64(21873915)
	tUnix := GetBlockTime(blockNumber)
	utcTime := BlockTimestampToTimeStr(tUnix)
	t.Logf("[*] utcTime: %s \n", utcTime)
}

//func TestGetBlockNumNow(t *testing.T) {
//
//	blockNumber := GetBlockNumNow()
//	utcTime := BlockTimestampToTimeStr(tUnix)
//	t.Logf("[*] utcTime: %s \n", utcTime)
//}

//func TestCalcFutureBlockNumTime(t *testing.T) {
//	syrupPool := GetSyrupPoolInfo("0x288d1aD79c113552B618765B4986f7DE679367Da")
//	// 获取奖励结束的区块数
//	t.Logf("[*] syrupPool.BonusEndBlockNum: %d \n", syrupPool.BonusEndBlockNum)
//	FutureBlockNumUTCTime := CalcFutureBlockNumTime(syrupPool.BonusEndBlockNum)
//
//	t.Logf("[*] FutureBlockNumUTCTime: %s \n", FutureBlockNumUTCTime)
//}
