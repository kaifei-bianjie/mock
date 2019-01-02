package key

import (
	"testing"

	"encoding/json"
	"github.com/kaifei-bianjie/mock/conf"
	"github.com/kaifei-bianjie/mock/util/constants"
	"github.com/kaifei-bianjie/mock/util/helper/account"
	"os"
	"strings"
)

var (
	faucetName = "mock-faucet-001"
)

func TestMain(m *testing.M) {
	// set config var
	conf.NodeUrl = "http://localhost:1317"
	conf.ChainId = "rainbow-dev"
	conf.BlockInterval = 5
	conf.FaucetSeed = "cube water sing thunder rib buyer assume rebuild cigar earn slight canoe apart grocery image satisfy genre woman mother can client science this tag"

	// create faucet account
	addr, err := account.NewKey(faucetName, constants.MockFaucetPassword, conf.FaucetSeed)
	if err != nil && !strings.Contains(err.Error(), "acount with name") {
		panic(err)
	}
	conf.FaucetAddress = addr

	existCode := m.Run()

	os.Exit(existCode)
}

func TestCreateFaucetSubAccount(t *testing.T) {
	type args struct {
		faucetName   string
		faucetPasswd string
		faucetAddr   string
		subAccNum    int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test create faucet sub account",
			args: args{
				faucetName:   faucetName,
				faucetPasswd: constants.KeyPassword,
				faucetAddr:   conf.FaucetAddress,
				subAccNum:    10,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := CreateFaucetSubAccount(tt.args.faucetName, tt.args.faucetPasswd, tt.args.faucetAddr, tt.args.subAccNum)
			if err != nil {
				t.Fatal(err)
			}
			resBytes, err := json.MarshalIndent(res, "", "")
			if err != nil {
				t.Fatal(err)
			}
			t.Log(string(resBytes))
		})
	}
}
