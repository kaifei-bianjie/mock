package account

import (
	"testing"

	"encoding/json"
	"github.com/kaifei-bianjie/mock/conf"
	"github.com/kaifei-bianjie/mock/util/constants"
	"math/rand"
	"os"
)

func TestMain(m *testing.M) {
	conf.NodeUrl = "http://localhost:1317"
	conf.ChainId = "rainbow-dev"

	conf.FaucetSeed = "cube water sing thunder rib buyer assume rebuild cigar earn slight canoe apart grocery image satisfy genre woman mother can client science this tag"

	conf.BlockInterval = 5
	conf.DefaultReceiverAddr = "faa1r5q5wqwctgfpt3p56qsctptrcq4st6lssyzx65"

	code := m.Run()
	os.Exit(code)
}

func TestCreateAccount(t *testing.T) {
	type args struct {
		name     string
		password string
		seed     string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test create account",
			args: args{
				name:     GenKeyName(constants.KeyNamePrefix, rand.Intn(10)),
				password: "1234567890",
				seed:     "",
			},
		},
		//{
		//	name: "test recover account",
		//	args: args{
		//		name:     constants.MockFaucetName,
		//		password: constants.MockFaucetPassword,
		//		seed:     conf.FaucetSeed,
		//	},
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := NewKey(tt.args.name, tt.args.password, tt.args.seed)
			if err != nil {
				t.Fatal(err)
			}
			t.Logf("address of new account is %v\n", res)
		})
	}
}

func TestGetAccountInfo(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test get account info",
			args: args{
				address: "faa1r5q5wqwctgfpt3p56qsctptrcq4st6lssyzx65",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := GetAccountInfo(tt.args.address)
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
