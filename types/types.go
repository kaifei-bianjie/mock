package types

import "github.com/irisnet/irishub/modules/auth"

type AccountInfo struct {
	LocalAccountName string `json:"name"`
	Password         string `json:"password"`
	AccountNumber    string `json:"account_number"`
	Sequence         string `json:"sequence"`
	Address          string `json:"address"`
}

type AccountInfoRes struct {
	AccountNumber string   `json:"account_number"`
	Sequence      string   `json:"sequence"`
	Address       string   `json:"address"`
	Coins         []string `json:"coins"`
}

type BaseTx struct {
	LocalAccountName string `json:"name"`
	Password         string `json:"password"`
	ChainID          string `json:"chain_id"`
	AccountNumber    string `json:"account_number"`
	Sequence         string `json:"sequence"`
	Gas              string `json:"gas"`
	Fees             string `json:"fee"`
	Memo             string `json:"memo"`
}

type TransferTxReq struct {
	Amount string `json:"amount"`
	Sender string `json:"sender"`
	BaseTx BaseTx `json:"base_tx"`
}

type ErrorRes struct {
	RestAPI      string `json:"rest api"`
	Code         int    `json:"code"`
	ErrorMessage string `json:"err message"`
}

type KeyCreateReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Seed     string `json:"seed"`
}

type KeyCreateRes struct {
	Address string `json:"address"`
}

type SignTxReq struct {
	Tx            auth.StdTx `json:"tx"`
	Name          string     `json:"name"`
	Password      string     `json:"password"`
	ChainID       string     `json:"chain_id"`
	AccountNumber int64      `json:"account_number"`
	Sequence      int64      `json:"sequence"`
	AppendSig     bool       `json:"append_sig"`
}

type BroadcastTxReq struct {
	Tx auth.StdTx `json:"tx"`
}

type PostTxReq struct {
	Tx auth.StdTx `json:"tx"`
}

type GenSignedTxDataRes struct {
	Res     string
	ChanNum int
}

type PostTx struct {
	Msgs       []string       `json:"msgs"`
	Fee        auth.StdFee    `json:"fee"`
	Signatures []StdSignature `json:"signatures"`
	Memo       string         `json:"memo"`
}

type StdSignature struct {
	PubKey        []byte `json:"pub_key"` // optional
	Signature     []byte `json:"signature"`
	AccountNumber int64  `json:"account_number"`
	Sequence      int64  `json:"sequence"`
}
