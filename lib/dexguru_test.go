package lib

import (
	"testing"
)

func TestGetTokenInfoByDexguru(t *testing.T) {
	shitTokenAddr := "0xe5ba47fd94cb645ba4119222e34fb33f59c7cd90"
	GetTokenInfoByDexguru(shitTokenAddr)

	//t.Logf("[*] Erc20Token: %+v \n", Erc20Token)
	if 1 == (2) {
		t.Error("error")
	}
}
