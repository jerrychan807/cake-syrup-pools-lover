package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/shopspring/decimal"
)

// IsValidAddress validate hex address
func IsValidAddress(iaddress interface{}) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	switch v := iaddress.(type) {
	case string:
		return re.MatchString(v)
	case common.Address:
		return re.MatchString(v.Hex())
	default:
		return false
	}
}

// IsZeroAddress validate if it's a 0 address
func IsZeroAddress(iaddress interface{}) bool {
	var address common.Address
	switch v := iaddress.(type) {
	case string:
		address = common.HexToAddress(v)
	case common.Address:
		address = v
	default:
		return false
	}

	zeroAddressBytes := common.FromHex("0x0000000000000000000000000000000000000000")
	addressBytes := address.Bytes()
	return reflect.DeepEqual(addressBytes, zeroAddressBytes)
}

// ToDecimal wei to decimals
func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)

	return result
}

// ToWei decimals to wei
func ToWei(iamount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}

// CalcGasCost calculate gas cost given gas limit (units) and gas price (wei)
func CalcGasCost(gasLimit uint64, gasPrice *big.Int) *big.Int {
	gasLimitBig := big.NewInt(int64(gasLimit))
	return gasLimitBig.Mul(gasLimitBig, gasPrice)
}

// SigRSV signatures R S V returned as arrays
func SigRSV(isig interface{}) ([32]byte, [32]byte, uint8) {
	var sig []byte
	switch v := isig.(type) {
	case []byte:
		sig = v
	case string:
		sig, _ = hexutil.Decode(v)
	}

	sigstr := common.Bytes2Hex(sig)
	rS := sigstr[0:64]
	sS := sigstr[64:128]
	R := [32]byte{}
	S := [32]byte{}
	copy(R[:], common.FromHex(rS))
	copy(S[:], common.FromHex(sS))
	vStr := sigstr[128:130]
	vI, _ := strconv.Atoi(vStr)
	V := uint8(vI + 27)

	return R, S, V
}

func EtherToWei(eth *big.Float, parm int) *big.Int {
	truncInt, _ := eth.Int(nil)
	potencial := strings.Repeat("0", parm)
	i, _ := strconv.ParseInt("1"+potencial, 10, 64)
	truncInt = new(big.Int).Mul(truncInt, big.NewInt(i))
	fracInt, _ := new(big.Int).SetString(potencial, parm)
	wei := new(big.Int).Add(truncInt, fracInt)
	return wei
}

func WeiToEther(wei *big.Int) string {
	fbal := new(big.Float)
	fbal.SetString(wei.String())
	ether := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(18)))
	sprintf := fmt.Sprintf("%.18f", ether)
	return sprintf
}
func WeiToEtherSpecificDecimal(wei *big.Int, decimals uint8) string {
	fbal := new(big.Float)
	fbal.SetString(wei.String())
	//fmt.Printf("[*] decimals: %d \n", int(decimals))
	ether := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
	sprintf := fmt.Sprintf("%.18f", ether)
	return sprintf
}

func BigFloatMulBigFloat(s1 string, s2 string) string {
	f1 := new(big.Float)
	f1.SetString(s1)
	f2 := new(big.Float)
	f2.SetString(s2)

	f1.Mul(f1, f2)
	sprintf := fmt.Sprintf("%.18f", f1)
	return sprintf
}

func BigFloat4Decimal(s1 string) string {
	f1 := new(big.Float)
	f1.SetString(s1)
	sprintf := fmt.Sprintf("%.4f", f1)
	return sprintf
}

// @title float64转换为百分比%,保留两位小数
func Float64ToPercentage(value float64) string {
	return strconv.FormatFloat(value*100, 'f', 2, 64) + "%"
}

// @title float64转换字符串,保留x位小数
// @param value float64 "浮点数"
// @param prec int "保留小数位"
func Float64ToStr(value float64, prec int) string {
	return strconv.FormatFloat(value, 'f', prec, 64)
}

// @title 字符串生成md5字符串
func Md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// @title 检查文件地址下文件是否存在的函数
func FileExists(fileAddr string) bool {
	if _, err := os.Stat(fileAddr); os.IsNotExist(err) {
		return false
	}
	return true
}

// @title 错误处理函数
func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// @title 错误处理函数
func GetNowTimeStr() string {
	nowStr := time.Now().Format("2006-01-02 15:04:05") //获取当前时间
	return nowStr
}

// @title 变化百分比<0为红色，>0为绿色
func GetColorByChange(change float64) string {
	if change < 0 {
		return "red"
	} else {
		return "green"
	}
}
