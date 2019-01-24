package sign

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/irisnet/irishub/codec"
	"github.com/irisnet/irishub/modules/auth"
	"github.com/irisnet/irishub/modules/bank"
	sdk "github.com/irisnet/irishub/types"
	"github.com/kaifei-bianjie/mock/conf"
	"github.com/kaifei-bianjie/mock/types"
	"github.com/kaifei-bianjie/mock/util/constants"
	"github.com/kaifei-bianjie/mock/util/helper"
	"github.com/kaifei-bianjie/mock/util/helper/tx"
	"log"
)

var (
	Cdc *codec.Codec
)

//custom tx codec
func init() {
	var cdc = codec.New()
	bank.RegisterCodec(cdc)
	auth.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	Cdc = cdc
}

// sign tx
func signTx(unsignedTx auth.StdTx, senderInfo types.AccountInfo) ([]byte, error) {
	// build request
	accountNumber, err := helper.ConvertStrToInt64(senderInfo.AccountNumber)
	if err != nil {
		return nil, err
	}
	sequence, err := helper.ConvertStrToInt64(senderInfo.Sequence)
	if err != nil {
		return nil, err
	}
	signTxReq := types.SignTxReq{
		Tx:            unsignedTx,
		Name:          senderInfo.LocalAccountName,
		Password:      senderInfo.Password,
		ChainID:       conf.ChainId,
		AccountNumber: accountNumber,
		Sequence:      sequence,
		AppendSig:     true,
	}

	// send sign tx request
	reqBytes, err := Cdc.MarshalJSON(signTxReq)
	if err != nil {
		return nil, err
	}
	reqBuffer := bytes.NewBuffer(reqBytes)
	statusCode, resBytes, err := helper.HttpClientPostJsonData(constants.UriTxSign, reqBuffer)

	// handle response
	if err != nil {
		return nil, err
	}

	if statusCode != constants.StatusCodeOk {
		return nil, fmt.Errorf("unexcepted status code: %v", statusCode)
	}

	return resBytes, nil
}

func buildBroadcastTxData(signedTx auth.StdTx) ([]byte, error) {
	req := types.BroadcastTxReq{
		Tx: signedTx,
	}

	return Cdc.MarshalJSON(req)
}

// generate signed tx
func GenSignedTxData(senderInfo types.AccountInfo, receiver string, resChan chan types.GenSignedTxDataRes, chanNum int) {
	var (
		unsignedTx, signedTx auth.StdTx
		method               = "GenSignedTxData"
	)
	log.Printf("%v: %v goroutine begin gen signed data\n", method, chanNum)

	signedTxDataRes := types.GenSignedTxDataRes{
		ChanNum: chanNum,
	}

	defer func() {
		if err := recover(); err != nil {
			log.Printf("%v: failed: %v\n", method, err)
		}

		resChan <- signedTxDataRes
	}()

	// build unsigned tx
	unsignedTxBytes, err := tx.SendTransferTx(senderInfo, receiver, "0.01iris", true)
	if err != nil {
		log.Printf("%v: build unsigned tx failed: %v\n", method, err)
		return
	}
	err = Cdc.UnmarshalJSON(unsignedTxBytes, &unsignedTx)
	if err != nil {
		log.Printf("%v: build unsigned tx failed: %v\n", method, err)
		return
	}

	// sign tx
	signedTxBytes, err := signTx(unsignedTx, senderInfo)
	if err != nil {
		log.Printf("%v: sign tx failed: %v\n", method, err)
		return
	}
	err = Cdc.UnmarshalJSON(signedTxBytes, &signedTx)
	if err != nil {
		log.Printf("%v: sign tx failed: %v\n", method, err)
		return
	}

	// build broadcast tx data
	broadcastTxBytes, err := buildBroadcastTxData(signedTx)
	if err != nil {
		log.Printf("%v: build post tx data failed: %v\n", method, err)
		return
	}

	signedTxDataRes.Res = base64.StdEncoding.EncodeToString(broadcastTxBytes)
}
