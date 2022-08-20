package lib

import (
	"testing"
)

var serializedPoolConfigStr = `export const livePools: SerializedPoolConfig[] = [
  {
    sousId: 0,
    stakingToken: serializedTokens.cake,
    earningToken: serializedTokens.cake,
    contractAddress: {
      97: '0xB4A466911556e39210a6bB2FaECBB59E4eB7E43d',
      56: '0xa5f8C5Dbd5F286960b9d90548680aE5ebFf07652',
    },
    poolCategory: PoolCategory.CORE,
    tokenPerBlock: '10',
    isFinished: false,
  },
  {
    sousId: 291,
    stakingToken: serializedTokens.cake,
    earningToken: serializedTokens.pstake,
    contractAddress: {
      56: '0x56D6955Ba6404647191DD7A5D65A5c9Fe43905e1',
      97: '',
    },
    poolCategory: PoolCategory.CORE,
    tokenPerBlock: '1.1574',
    version: 3,
  },
]`

func TestParseContractAddrField(t *testing.T) {
	sousIds := parseSousIdField(serializedPoolConfigStr)
	t.Logf("[*] sousIds: %s\n", sousIds)
	if len(sousIds) == 0 {
		t.Errorf("%s \n", sousIds)
	}
}
