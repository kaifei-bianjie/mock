package tx

import (
	"testing"

	"github.com/kaifei-bianjie/mock/conf"
	"github.com/kaifei-bianjie/mock/types"
	"os"
)

func TestMain(m *testing.M) {
	conf.NodeUrl = "http://localhost:1317"
	conf.ChainId = "rainbow-qa"

	conf.FaucetSeed = "tent tube capable grit volume enforce wash snow tilt clip stable alert drip fence huge recycle excess focus jump antique creek area meadow alarm"

	conf.BlockInterval = 5
	conf.DefaultReceiverAddr = "faa1z72xhn7nq4u0jtvpkgehfwc8u6s5jjer2kvx28"

	code := m.Run()
	os.Exit(code)
}

func TestSendTransferTx(t *testing.T) {
	type args struct {
		senderInfo   types.AccountInfo
		receiver     string
		generateOnly bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test send transfer tx",
			args: args{
				senderInfo: types.AccountInfo{
					LocalAccountName: "kaifei",
					Password:         "1234567890",
					AccountNumber:    "51409",
					Sequence:         "0",
					Address:          "faa1z72xhn7nq4u0jtvpkgehfwc8u6s5jjer2kvx28",
				},
				receiver:     conf.DefaultReceiverAddr,
				generateOnly: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := SendTransferTx(tt.args.senderInfo, tt.args.receiver, "", tt.args.generateOnly)
			if err != nil {
				t.Fatal(err)
			}
			t.Log(string(res))
		})
	}
}
