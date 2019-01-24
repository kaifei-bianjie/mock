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
	UriKeyCreate   = "/keys"
	UriAccountInfo = "/auth/accounts/%v"           // format is /auth/accounts/{address}
	UriTransfer    = "/bank/accounts/%s/transfers" // format is /bank/accounts/{address}/transfers
	UriTxSign      = "/tx/sign"

	// http status code
	StatusCodeOk       = 200
	StatusCodeConflict = 409

	//
	MockFaucetName     = "mock-faucet"
	MockFaucetPassword = "1234567890"
	MockTransferAmount = "1iris"
	MockDefaultGas     = "20000"
	MockDefaultFee     = "0.5iris"
	Denom              = "iris"

	// http timeout
	HttpTimeout = 60 * time.Second
)
