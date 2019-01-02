package constants

import "time"

const (
	// config file name
	ConfigFileName = "config.json"

	HeaderContentTypeJson = "application/json"

	// key password, prefix of key name
	KeyNamePrefix = "mock"
	KeyPassword   = "1234567890"

	// http uri
	UriKeyCreate     = "/keys"
	UriAccountInfo   = "/auth/accounts/%v"           // format is /auth/accounts/{address}
	UriTransfer      = "/bank/accounts/%s/transfers" // format is /bank/accounts/{address}/transfers
	UriTxSign        = "/tx/sign"
	UriTxBroadcastTx = "/txs/send"

	// http status code
	StatusCodeOk       = 200
	StatusCodeConflict = 409

	//
	MockFaucetName     = "mock-faucet"
	MockFaucetPassword = "1234567890"
	MockTransferAmount = "0.03iris"
	MockDefaultGas     = "200000"
	MockDefaultFee     = "0.01iris"
	Denom              = "iris"

	// http timeout
	HttpTimeout = 60 * time.Second
)
