package cmd

import (
	"fmt"
	"github.com/kaifei-bianjie/mock/conf"
	"github.com/kaifei-bianjie/mock/service"
	"github.com/kaifei-bianjie/mock/util/constants"
	"github.com/kaifei-bianjie/mock/util/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
	"time"
)

func GenSignedTxDataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gen-signed-tx",
		Short: "generate signed tx data",
		Long:  `batch generate signed tx data`,
		Example: `
	mock gen-signed-tx --num {num} --receiver {receiver-address} \&
--home {config-home} --chain-id {chain-id} --node {node-url}
`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			fmt.Println("Start")
			var (
				subFaucets []conf.SubFaucet
			)

			// create output dir
			outputDir := viper.GetString(FlagResOutput)
			err := helper.CreateFolder(outputDir)
			if err != nil {
				return err
			}

			// read config from config file
			// check config file if exists
			confHome := viper.GetString(FlagConfDir)
			confFilePath := fmt.Sprintf("%v/%v", confHome, constants.ConfigFileName)

			exists, err := helper.CheckFileExist(confFilePath)
			if err != nil {
				return err
			}
			if !exists {
				return fmt.Errorf("can't find config file in %v\n", confFilePath)
			}

			// read config from config file
			viper.SetConfigName("config")
			viper.AddConfigPath(confHome)
			err = viper.ReadInConfig()
			if err != nil {
				return err
			}
			err = viper.UnmarshalKey("sub_faucets", &subFaucets)
			if err != nil {
				return err
			}
			if len(subFaucets) <= 0 {
				return fmt.Errorf("can't read sub_faucets config")
			}

			// generate signed tx
			num := viper.GetInt(FlagNumSignedTx)
			if num < len(subFaucets) {
				return fmt.Errorf("%v(num) shouldn't less than %v(num of sub faucet)",
					num, len(subFaucets))
			}
			if num > 100000 {
				return fmt.Errorf("num should less than 100000\n")
			}

			signedTxData := service.BatchGenSignedTxData(num, subFaucets)

			if len(signedTxData) > 0 {
				// write result to file
				filename := fmt.Sprintf("res_signed_tx_%v", time.Now().Unix())
				filePath := fmt.Sprintf("%v/%v", outputDir, filename)
				err = helper.WriteFile(filePath, []byte(strings.Join(signedTxData, "\n")))
				if err != nil {
					return err
				}
			} else {
				fmt.Println("no signed tx data")
			}
			fmt.Println("End")
			return nil
		},
	}

	cmd.Flags().AddFlagSet(txFlagSet)

	cmd.MarkFlagRequired(FlagNumSignedTx)
	cmd.MarkFlagRequired(FlagReceiverAddr)
	cmd.MarkPersistentFlagRequired(FlagChainId)
	cmd.MarkPersistentFlagRequired(FlagNodeUrl)

	return cmd
}
