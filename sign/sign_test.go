package sign

import (
	"github.com/kaifei-bianjie/mock/conf"
	"github.com/kaifei-bianjie/mock/types"
	"github.com/kaifei-bianjie/mock/util/constants"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	conf.NodeUrl = "http://35.235.123.127:1317"
	conf.ChainId = "fuxi-8000"

	conf.BlockInterval = 5
	conf.DefaultReceiverAddr = "faa1r5q5wqwctgfpt3p56qsctptrcq4st6lssyzx65"

	code := m.Run()

	os.Exit(code)
}

func TestBroadcastSignedTx(t *testing.T) {

	type args struct {
		senderInfo types.AccountInfo
		receiver   string
		resChan    chan types.GenSignedTxDataRes
		chanNUm    int
	}

	resChannel := make(chan types.GenSignedTxDataRes)

	tests := []struct {
		name string
		args args
	}{
		{
			name: "test gen a signed tx",
			args: args{
				senderInfo: types.AccountInfo{
					LocalAccountName: constants.MockFaucetName,
					Password:         constants.MockFaucetPassword,
					AccountNumber:    "356",
					Sequence:         "1",
					Address:          "faa1gcveln4rdkw4gs9gyd8hkapcjz9q55q0khw3st",
				},
				receiver: conf.DefaultReceiverAddr,
				chanNUm:  1,
				resChan:  resChannel,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go GenSignedTxData(tt.args.senderInfo, tt.args.receiver, tt.args.resChan, tt.args.chanNUm)

			res := <-tt.args.resChan
			if res.ChanNum != 0 {
				t.Logf("%v build signed tx data over\n", res.ChanNum)
				t.Log(res.Res)
			}
			log.Println(res.ChanNum)
		})
	}
}
