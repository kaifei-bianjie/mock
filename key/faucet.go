package key

import (
	"fmt"
	"github.com/kaifei-bianjie/mock/types"
	"github.com/kaifei-bianjie/mock/util/constants"
	"github.com/kaifei-bianjie/mock/util/helper"
	"github.com/kaifei-bianjie/mock/util/helper/account"
	"log"
	"strconv"
	"strings"
)

// create faucet sub accounts and transfer token to them.
// faucetName and faucetPasswd specified by flag,
// subAccNum is the num of sub account which need create.
func CreateFaucetSubAccount(faucetName, faucetPasswd, faucetAddr string, subAccNum int) ([]types.AccountInfo, error) {
	var (
		method               = "CreateFaucetSubAccount"
		createdAccs, subAccs []types.AccountInfo
	)

	keyChan := make(chan types.AccountInfo)

	// create sub account
	for i := 1; i <= subAccNum; i++ {
		keyName := fmt.Sprintf("%v_%v", faucetName, i)
		go CreateAccount(keyName, keyChan)
	}

	counter := 0
	for {
		accInfo := <-keyChan
		if accInfo.Address != "" {
			createdAccs = append(createdAccs, accInfo)
		}
		counter++
		if counter == subAccNum {
			log.Printf("%v: all create sub faucet key goroutine over\n", method)
			log.Printf("%v: except create %v accounts, successful create %v accounts",
				method, subAccNum, len(createdAccs))
			break
		}
	}

	// distribute token

	// get sender info
	senderInfo := types.AccountInfo{
		LocalAccountName: faucetName,
		Password:         faucetPasswd,
		Address:          faucetAddr,
	}
	accInfo, err := account.GetAccountInfo(senderInfo.Address)
	if err != nil {
		log.Printf("%v: get faucet info fail: %v\n", method, err)
		return subAccs, err
	}
	senderInfo.AccountNumber = accInfo.AccountNumber
	senderSequence, err := helper.ConvertStrToInt64(accInfo.Sequence)
	if err != nil {
		log.Printf("%v: convert sequence to int64 fail: %v\n", method, err)
		return subAccs, err
	}

	// get transfer amount which equal senderBalance / subAccNum
	amt, err := parseCoins(accInfo.Coins)
	if err != nil {
		log.Printf("%v: parse coin failed: %v\n", method, err)
		return subAccs, err
	}
	transferAmt := fmt.Sprintf("%v%s", parseFloat64ToStr(amt/float64(subAccNum+1)), constants.Denom)

	// distribute token to created accounts
	for _, acc := range createdAccs {
		senderInfo.Sequence = fmt.Sprintf("%v", senderSequence)
		acc, err := DistributeToken(senderInfo, acc, transferAmt)
		if err != nil {
			log.Printf("%v: distribute token to %v failed: %v\n",
				method, acc.LocalAccountName, err)
		} else {
			subAccs = append(subAccs, acc)
			senderSequence += 1
		}
	}

	return subAccs, err
}

func parseCoins(coins []string) (float64, error) {
	coin := coins[0]
	amtStr := strings.Replace(coin, constants.Denom, "", -1)
	amtFloat, err := strconv.ParseFloat(amtStr, 64)

	if err != nil {
		return float64(0), nil
	}
	return amtFloat, nil
}

func parseFloat64ToStr(amt float64) string {
	return strconv.FormatFloat(amt, 'f', -1, 64)
}
