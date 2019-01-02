package account

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kaifei-bianjie/mock/types"
	"github.com/kaifei-bianjie/mock/util/constants"
	"github.com/kaifei-bianjie/mock/util/helper"
	"github.com/satori/go.uuid"
	"log"
)

func GenKeyName(namePrefix string, num int) string {
	uid := uuid.NewV4().String()
	return fmt.Sprintf("%s_%v_%v", namePrefix, uid, num)
}

// create key
func CreateAccount(name, password, seed string) (string, error) {
	req := types.KeyCreateReq{
		Name:     name,
		Password: password,
		Seed:     seed,
	}

	uri := constants.UriKeyCreate

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	reqBody := bytes.NewBuffer(reqBytes)

	statusCode, resBytes, err := helper.HttpClientPostJsonData(uri, reqBody)

	if err != nil {
		return "", err
	}

	if statusCode == constants.StatusCodeOk {
		res := types.KeyCreateRes{}
		if err := json.Unmarshal(resBytes, &res); err != nil {
			return "", err
		}
		return res.Address, nil
	} else if statusCode == constants.StatusCodeConflict {
		return "", fmt.Errorf("%v", string(resBytes))
	} else {
		log.Printf("%v, statusCode %v, res %v, err is %v\n",
			"createAccount", statusCode, string(resBytes), err)
		errRes := types.ErrorRes{}
		if err := json.Unmarshal(resBytes, &errRes); err != nil {
			return "", err
		}
		return "", fmt.Errorf("err code: %v, err msg: %v", errRes.Code, errRes.ErrorMessage)
	}
}

// get account info
func GetAccountInfo(address string) (types.AccountInfoRes, error) {
	var (
		accountInfo types.AccountInfoRes
	)
	uri := fmt.Sprintf(constants.UriAccountInfo, address)
	statusCode, resByte, err := helper.HttpClientGetData(uri)

	if err != nil {
		return accountInfo, err
	}

	if statusCode == constants.StatusCodeOk {
		if err := json.Unmarshal(resByte, &accountInfo); err != nil {
			return accountInfo, err
		}
		return accountInfo, nil
	} else {
		return accountInfo, fmt.Errorf("status code is not ok, code: %v", statusCode)
	}
}
