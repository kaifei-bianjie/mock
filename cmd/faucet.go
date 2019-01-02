package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/kaifei-bianjie/mock/conf"
	"github.com/kaifei-bianjie/mock/key"
	"github.com/kaifei-bianjie/mock/util/constants"
	"github.com/kaifei-bianjie/mock/util/helper"
	"github.com/kaifei-bianjie/mock/util/helper/account"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func FaucetInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "faucet-init",
		Short: "create sub faucet account use faucet account",
		Long: `
Note the account must has many token, so that this account can transfer token to other account.
`,
		Example: `
mock faucet-init --faucet-name {faucet-name} --seed="recycle light kid ..." \&
--sub-faucet-num {sub-faucet-num} --home {config-home} \&
--chain-id {chain-id} --node {node}
`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			// get flag and validate basic logic
			var (
				configContent    conf.ConfigContent
				configSubFaucets []conf.SubFaucet
			)
			seed := viper.GetString(FlagFaucetSeed)
			name := viper.GetString(FlagFaucetName)
			confHomeDir := viper.GetString(FlagConfDir)
			subFaucetNum := viper.GetInt(FlagSubFaucetAccNum)
			confFilePath := fmt.Sprintf("%v/%v", confHomeDir, constants.ConfigFileName)

			if subFaucetNum > 10 {
				return fmt.Errorf("num of sub faucet account shouldn't greater than 10")
			}

			exists, err := helper.CheckFileExist(confFilePath)
			if err != nil {
				panic(err)
			}
			if exists {
				return fmt.Errorf("config file alread exist in %v\n, "+
					"please remove it before exec this command", confHomeDir)
			}

			err = helper.CreateFolder(confHomeDir)
			if err != nil {
				panic(err)
			}

			// recover faucet by seed
			address, err := account.NewKey(name, constants.MockFaucetPassword, seed)
			if err != nil {
				return err
			}

			// create sub faucet account
			fmt.Printf("now create %v sub faucet account\n", subFaucetNum)
			subAccs, err := key.CreateFaucetSubAccount(name, constants.MockFaucetPassword, address, subFaucetNum)
			if err != nil {
				return err
			}

			// write config content to file
			for _, acc := range subAccs {
				subFaucet := conf.SubFaucet{
					FaucetName:     acc.LocalAccountName,
					FaucetPassword: acc.Password,
					FaucetAddr:     acc.Address,
				}

				configSubFaucets = append(configSubFaucets, subFaucet)
			}
			configContent.FaucetName = name
			configContent.FaucetAddr = address
			configContent.FaucetSeed = seed
			configContent.SubFaucets = configSubFaucets

			configBytes, err := json.MarshalIndent(configContent, "", "\t")
			if err != nil {
				return err
			}

			err = helper.WriteFile(confFilePath, configBytes)
			if err != nil {
				return err
			}

			fmt.Printf("success init faucet info in %v\n", confFilePath)
			return nil
		},
	}

	cmd.Flags().AddFlagSet(faucetFlagSet)
	cmd.MarkFlagRequired(FlagFaucetName)
	cmd.MarkFlagRequired(FlagFaucetSeed)
	cmd.MarkPersistentFlagRequired(FlagChainId)
	cmd.MarkPersistentFlagRequired(FlagNodeUrl)

	return cmd
}
