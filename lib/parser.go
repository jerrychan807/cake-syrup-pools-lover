package lib

import (
	"fmt"
	"regexp"
	"strings"
)

func ParsePoolsInfo(content string) ([]string, []string, []string, []string) {
	fmt.Println("[*] ParsePoolsInfo through SerializedPoolConfig Strings ")
	sousIds := parseSousIdField(content)
	earningTokens := parseEarningTokenField(content)
	contractAddrs := parseContractAddrField(content)
	tokenPerBlocks := parseTokenPerBlockField(content)
	//fmt.Printf("[*] sousIds: %s\n", sousIds)
	//fmt.Printf("[*] earningTokens: %s\n", earningTokens)
	//fmt.Printf("[*] contractAddrs: %s\n", contractAddrs)
	//fmt.Printf("[*] tokenPerBlocks: %s\n", tokenPerBlocks)

	return sousIds, earningTokens, contractAddrs, tokenPerBlocks

}

func parseSousIdField(content string) []string {
	var re = regexp.MustCompile(`(?m)sousId: (\d*)`)
	var sousIds []string
	for _, match := range re.FindAllString(content, -1) {
		sousId := strings.Replace(match, "sousId: ", "", 1)
		sousIds = append(sousIds, sousId)
	}
	return sousIds
}

func parseEarningTokenField(content string) []string {
	var re = regexp.MustCompile(`(?m)earningToken: serializedTokens\.(\S*),`)
	var tempStrs []string
	for _, match := range re.FindAllString(content, -1) {
		tempStr := strings.Replace(match, "earningToken: serializedTokens.", "", 1)
		tempStr = strings.Replace(tempStr, ",", "", 1)
		tempStrs = append(tempStrs, tempStr)
	}
	return tempStrs
}

func parseContractAddrField(content string) []string {
	var re = regexp.MustCompile(`(?m)56: (\S*),`)
	var tempStrs []string
	for _, match := range re.FindAllString(content, -1) {
		tempStr := strings.Replace(match, "56: ", "", 1)
		tempStr = strings.Replace(tempStr, "'", "", 2)
		tempStr = strings.Replace(tempStr, ",", "", 1)
		tempStrs = append(tempStrs, tempStr)
	}
	return tempStrs
}

func parseTokenPerBlockField(content string) []string {
	var re = regexp.MustCompile(`(?m)tokenPerBlock: (\S*),`)
	var tempStrs []string
	for _, match := range re.FindAllString(content, -1) {
		tempStr := strings.Replace(match, "tokenPerBlock: ", "", 1)
		tempStr = strings.Replace(tempStr, "'", "", 2)
		tempStr = strings.Replace(tempStr, ",", "", 1)
		tempStrs = append(tempStrs, tempStr)
	}
	return tempStrs
}
